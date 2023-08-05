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

var API_KEYS = []string{"jByRuDyDvM3bQl71hLgadWt932jodjvpJRqvXsQRIWHfpSZwxYBR7BWFBOXO7o6b", "ceox9ksHbEyMQiOkSkH0k2VnFWkiojJEucGIMRUZXUq3dLCuOqV81nQIjCh0PLK5", "cUe7mZpWMNi96Gunm57USrkstCaZOuOoR7qWzFHbAIjglh0wsAgCufwqQARHX59i", "AJaBKMzpGAyYeiP7hV4aN5IwuJL7W7RtExEdrxyBh6GUf875rIiS2Ek1iZIakx4C"}

var SOCKS5_PROXY_ADDRESSES = []string{"212.115.48.122:13225", "195.225.96.39:13225", "193.160.211.65:13225", "45.86.3.128:13225"}

var pairsList = [][]string{
	{"1inchbtc", "1inchbusd", "1inchusdt", "aavebnb", "aavebtc", "aavebusd", "aaveusdt", "acabtc", "acabusd", "acatry", "acausdt", "achbtc", "achbusd", "achtry", "achusdt", "acmbusd", "acmusdt", "adabnb", "adabrl", "adabtc", "adabusd", "adaeth", "adaeur", "adatry", "adausdt", "adadownusdt", "adaupusdt", "adxusdt", "aergousdt", "agixbtc", "agixbusd", "agixtry", "agixusdt", "agldbtc", "agldbusd", "agldusdt", "akrousdt", "alcxusdt", "algobtc", "algobusd", "algotry", "algousdt", "alicebusd", "aliceusdt", "alpacabtc", "alpacabusd", "alpacausdt", "alphausdt", "alpinebusd", "alpinetry", "alpineusdt", "ambbusd", "ambusdt", "ampbusd", "ampusdt", "ankrbtc", "ankrbusd", "ankrtry", "ankrusdt", "antusdt", "apebtc", "apebusd", "apetry", "apeusdt", "api3usdt", "aptbtc", "aptbusd", "apttry", "aptusdt", "arbtc", "arusdt", "arbbtc", "arbeth", "arbtry", "arbtusd", "arbusdt", "ardrusdt", "arkbusd", "arkmbnb", "arkmbtc", "arkmtry", "arkmtusd", "arkmusdt", "arpabusd", "arpatry", "arpausdt", "asrusdt", "astbtc", "astusdt", "astrbtc", "astrbusd", "astrusdt", "atausdt", "atmbusd", "atmusdt"},
	{"atomusdt", "auctionbtc", "auctionbusd", "auctionusdt", "audiotry", "audiousdt", "avabtc", "avausdt", "avaxbnb", "avaxbtc", "avaxbusd", "avaxeth", "avaxeur", "avaxtry", "avaxusdt", "axsbtc", "axsbusd", "axsusdt", "badgerusdt", "bakebusd", "bakeusdt", "balusdt", "bandbusd", "bandusdt", "barbusd", "barusdt", "batbusd", "batusdt", "bchbnb", "bchbtc", "bchbusd", "bcheur", "bchtry", "bchusdt", "bdotdot", "belbusd", "beltry", "belusdt", "betausdt", "betheth", "bethusdt", "bicousdt", "bifiusdt", "blzbtc", "blzusdt", "bnbbidr", "bnbbrl", "bnbbtc", "bnbbusd", "bnbeth", "bnbeur", "bnbfdusd", "bnbgbp", "bnbtry", "bnbtusd", "bnbusdc", "bnbusdt", "bnbdownusdt", "bnbupusdt", "bntusdt", "bnxusdt", "bondbusd", "bondusdt", "bswbusd", "bswtry", "bswusdt", "btcars", "btcbidr", "btcbrl", "btcbusd", "btcdai", "btceur", "btcgbp", "btcngn", "btcpln", "btcrub", "btctry", "btctusd", "btcusdc", "btcusdt", "btczar", "btcdownusdt", "btcupusdt", "btsusdt", "bttctry", "bttcusdt", "burgerbusd", "burgerusdt", "busdbidr", "busdbrl", "busddai", "busdpln", "busdrub", "busdtry", "busdusdt"},
	{"celobtc", "celobusd", "celousdt", "celrbusd", "celrusdt", "cfxbtc", "cfxbusd", "cfxtry", "cfxusdt", "chessbtc", "chessbusd", "chessusdt", "chrbtc", "chrbusd", "chrusdt", "chzbtc", "chzbusd", "chztry", "chzusdt", "citybusd", "citytry", "cityusdt", "ckbusdt", "clvbtc", "clvbusd", "clvusdt", "combotry", "combousdt", "compbtc", "compbusd", "comptry", "compusdt", "cosbtc", "costry", "cosusdt", "cotibusd", "cotiusdt", "creambusd", "crvbtc", "crvbusd"},
}

func StartPulsar(dataChannel chan []byte) {
	var wg sync.WaitGroup

	for i, pairs := range pairsList {
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
		}(pairs, API_KEYS[i%len(API_KEYS)], SOCKS5_PROXY_ADDRESSES[i%len(SOCKS5_PROXY_ADDRESSES)])
	}

	wg.Wait()
}
