// binance_pulsar.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/mailru/easyjson"
)

var API_KEY = "jByRuDyDvM3bQl71hLgadWt932jodjvpJRqvXsQRIWHfpSZwxYBR7BWFBOXO7o6b"

var pairsList = [][]string{
	{"1inchbtc", "1inchbusd", "1inchusdt"},
	{"alpineusdt", "ambbusd", "ambusdt"},
	{"atomusdt", "auctionbtc", "btcusdt"},
}

func StartPulsar(dataChannel chan []byte) {
	var wg sync.WaitGroup

	for _, pairs := range pairsList {
		wg.Add(1)
		go func(pairs []string) {
			defer wg.Done()

			for i, pair := range pairs {
				pairs[i] = pair + "@bookTicker"
			}

			header := http.Header{}
			header.Add("X-MBX-APIKEY", API_KEY)

			conn, _, err := websocket.DefaultDialer.Dial("wss://stream.binance.com:9443/stream", header)
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

				dataChannel <- message
			}
		}(pairs)
	}

	wg.Wait()
}
