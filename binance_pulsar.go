// binance_pulsar.go
package main

import (
	"log"
	"net"
	"net/http"
	"sync"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/valyala/fastjson"
	goproxy "golang.org/x/net/proxy"
)

type Config = GoCryptoTraderConfig

func StartPulsar(dataChannels []chan []byte, config GoCryptoTraderConfig, ports []string) {
	var wg sync.WaitGroup
	var logCount int

	log.Println("StartPulsar called")  // Log 1: Check if StartPulsar is called

	// Transform ExchangePairsList into a 2D slice of strings
	var pairsList [][]string
	chunkSize := len(ExchangePairsList) / 6  // Size of each sub-slice
	for i := 0; i < len(ExchangePairsList); i += chunkSize {
		end := i + chunkSize
		if end > len(ExchangePairsList) {
			end = len(ExchangePairsList)
		}
		pairsList = append(pairsList, ExchangePairsList[i:end])
	}

	log.Printf("Pairs List: %v", pairsList)  // Log 2: Check the content of pairsList

	for i, pairs := range pairsList {
		wg.Add(1)
		go func(pairs []string, apiKey string, proxyAddress string, port string) {
			defer wg.Done()

			log.Printf("Goroutine for pairs: %v started", pairs)  // Log 3: Check if the goroutine starts and what pairs are passed

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

					var p fastjson.Parser
					v, err := p.Parse(string(message))
					if err != nil {
						log.Fatalf("Failed to parse JSON: %s", err)
					}

					if v.Exists("u") {
						logCount++
						log.Printf("Log calls: %d", logCount)
					}
				}(message)
			}
		}(pairs, config.APIKeys[i%len(config.APIKeys)], config.ProxyAddresses[i%len(config.ProxyAddresses)], ports[i%len(ports)])
	}

	wg.Wait()
	log.Println("All goroutines finished")  // Log 5: All goroutines are done
}
