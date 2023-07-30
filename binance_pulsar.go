// binance_pulsar.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var API_KEY = "jByRuDyDvM3bQl71hLgadWt932jodjvpJRqvXsQRIWHfpSZwxYBR7BWFBOXO7o6b"

var pairsList = [][]string{
	{"1inchbtc", "1inchbusd", "1inchusdt", "aavebnb", "aavebtc", "aavebusd"},
	{"alpineusdt", "ambbusd", "ambusdt", "ampbusd", "ampusdt", "ankrbtc"},
	{"atomusdt", "auctionbtc", "auctionbusd", "auctionusdt", "audiotry"},
	{"bnbeur", "bnbfdusd", "bnbgbp", "bnbtry", "bnbtusd", "bnbusdc"},
	{"celobtc", "celobusd", "celousdt", "celrbusd", "celrusdt"},
	{"cvxusdt", "darbusd", "dartry", "darusdt", "dashbtc"},
	{"egldusdt", "elfbtc", "elfbusd", "elfeth", "elfusdt"},
	{"firobtc", "firobusd", "firousdt", "fisbtc", "fisusdt"},
	{"gmttry", "gmtusdt", "gmxbusd", "gmxusdt", "gnousdt"},
	{"iotabtc", "iotausdt", "iotxusdt", "iqbusd", "irisusdt"},
	{"lokausdt", "loombtc", "loomusdt", "lptusdt", "lqtyusdt"},
	{"mblbusd", "mblusdt", "mboxbusd", "mboxtry", "mboxusdt"},
	{"omgusdt", "onebtc", "onebusd", "oneusdt", "ongbtc", "ongusdt"},
	{"prosusdt", "psgbusd", "psgusdt", "pundixusdt", "pyrusdt"},
	{"rvnbusd", "rvntry", "rvnusdt", "sandbtc", "sandbusd"},
	{"steemusdt", "stgbusd", "stgusdt", "stmxbusd", "stmxusdt"},
	{"tomousdt", "tornbusd", "trbbtc", "trbbusd", "trbusdt"},
	{"utkusdt", "vetbnb", "vetbtc", "vetbusd", "veteth"},
	{"xmrbtc", "xmrbusd", "xmreth", "xmrusdt", "xnousdt"},
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
