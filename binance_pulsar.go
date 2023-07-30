package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type BinanceMessage struct {
	Symbol string `json:"s"`
}

var API_KEY = "jByRuDyDvM3bQl71hLgadWt932jodjvpJRqvXsQRIWHfpSZwxYBR7BWFBOXO7o6b"

var pairsList = [][]string{
	{"1inchbtc", "1inchbusd", "1inchusdt", "aavebnb", "aavebtc", "aavebusd", "aaveusdt", "acabtc", "acabusd", "acatry", "acausdt", "achbtc", "achbusd", "achtry", "achusdt", "acmbusd", "acmusdt", "adabnb", "adabrl", "adabtc", "adabusd", "adaeth", "adaeur", "adatry", "adausdt", "adadownusdt"},
	{"alpineusdt", "ambbusd", "ambusdt", "ampbusd", "ampusdt", "ankrbtc", "ankrbusd", "ankrtry", "ankrusdt", "antusdt", "apebtc", "apebusd", "apetry", "apeusdt", "api3usdt", "aptbtc", "aptbusd", "apttry", "aptusdt", "arbtc", "arusdt", "arbbtc", "arbeth", "arbtry", "arbtusd", "arbusdt"},
	{"atomusdt", "auctionbtc", "auctionbusd", "auctionusdt", "audiotry", "audiousdt", "avabtc", "avausdt", "avaxbnb", "avaxbtc", "avaxbusd", "avaxeth", "avaxeur", "avaxtry", "avaxusdt", "axsbtc", "axsbusd", "axsusdt", "badgerusdt", "bakebusd", "bakeusdt", "balusdt", "bandbusd", "btcusdt"},
}

func StartApp() {
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

				log.Println("Received message: ", string(message))

				var binanceMessage BinanceMessage
				err = json.Unmarshal(message, &binanceMessage)
				if err != nil {
					log.Println("Error unmarshalling message: ", err)
					continue
				}
			}
		}(pairs)
	}

	wg.Wait()
}
