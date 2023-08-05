// binance_pulsar.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	goproxy "golang.org/x/net/proxy"
)

func StartPulsar(dataChannel chan []byte, config Config) {
	var wg sync.WaitGroup

	for i, pairs := range config.PairsList {
		wg.Add(1)
		go func(pairs []string, apiKey string, proxyAddress string) {
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

			for {
				_, message, err := conn.ReadMessage()
				if err != nil {
					log.Println("Error reading message: ", err)
					break
				}

				wg.Add(1)
				go func(message []byte) {
					defer wg.Done()
					dataChannel <- message
				}(message)
			}
		}(pairs, config.APIKeys[i%len(config.APIKeys)], config.ProxyAddresses[i%len(config.ProxyAddresses)])
	}

	wg.Wait()
}
