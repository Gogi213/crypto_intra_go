// binance_pulsar
package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	goproxy "golang.org/x/net/proxy"
)

type Config struct {
	APIKeys        []string
	ProxyAddresses []string
	PairsList      [][]string
}

func StartPulsar(dataChannels []chan []byte, config Config, ports []string) {
	var wg sync.WaitGroup
	var logCount int

	for i, pairs := range config.PairsList {
		wg.Add(1)
		go func(pairs []string, apiKey string, proxyAddress string, port string) {
			defer wg.Done()

			for i, pair := range pairs {
				pairs[i] = pair + "@bookTicker"
			}

			dialer, err := goproxy.SOCKS5("tcp", proxyAddress, &goproxy.Auth{User: "user128676", Password: "atioln"}, goproxy.Direct)
			if err != nil {
				log.Fatal("Error creating dialer: ", err)
			}

			wsDialer := &websocket.Dialer{
				NetDial: dialer.Dial,
			}

			header := http.Header{}
			header.Add("X-MBX-APIKEY", apiKey)

			conn, _, err := wsDialer.Dial("wss://stream.binance.com:9443/stream", header)
			if err != nil {
				log.Fatal("Error connecting to WebSocket API: ", err)
			}
			defer conn.Close()

			params := map[string]interface{}{
				"method": "SUBSCRIBE",
				"params": pairs,
				"id":     1,
			}

			paramsJSON, _ := json.Marshal(params)
			if err := conn.WriteMessage(websocket.TextMessage, paramsJSON); err != nil {
				log.Fatal(err)
			}

			tcpConn, err := net.Dial("tcp", "localhost:"+port)
			if err != nil {
				log.Fatal(err)
			}
			defer tcpConn.Close()

			for {
				_, message, err := conn.ReadMessage()
				if err != nil {
					log.Println("Error reading message: ", err)
					break
				}

				wg.Add(1)
				go func(message []byte) {
					defer wg.Done()

					_, err = tcpConn.Write(append(message, '\n'))
					if err != nil {
						log.Println(err)
					}

					var data map[string]interface{}
					json.Unmarshal(message, &data)
					if _, ok := data["u"]; ok {
						logCount++
						log.Printf("Log calls: %d", logCount)
					}

					log.Printf("Sending message to dataChannel: %s", string(message))

					// Write data to CSV
					if data, ok := data["data"].(map[string]interface{}); ok {
						writeDataToCSV([]string{data["s"].(string), data["b"].(string), data["B"].(string), data["a"].(string), data["A"].(string)})
					}
				}(message)
			}
		}(pairs, config.APIKeys[i%len(config.APIKeys)], config.ProxyAddresses[i%len(config.ProxyAddresses)], ports[i%len(ports)])
	}

	wg.Wait()
}
