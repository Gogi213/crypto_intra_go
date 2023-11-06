// config.go
package main

import (
	"encoding/csv"
	"os"
)

type GoCryptoTraderConfig struct {
	APIKeys        []string `json:"apiKeys"`
	ProxyAddresses []string `json:"proxyAddresses"`
	EnabledPairs   []string `json:"enabledPairs"`
}

var exchangePairsList []string

func fillPairsFromCSV(startIndex int, maxPairs int) []string {
	var pairs []string
	for i := startIndex; i < startIndex+maxPairs; i++ {
		if i >= len(exchangePairsList) {
			break
		}
		pairs = append(pairs, exchangePairsList[i])
	}
	return pairs
}

func init() {
	// Загрузка пар из CSV в exchangePairsList
	file, err := os.Open("currency_pairs.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, record := range records {
		exchangePairsList = append(exchangePairsList, record[0])
	}

	// Разделение пар между инстансами
	maxPairs := len(exchangePairsList) / 6
	Instance1.EnabledPairs = fillPairsFromCSV(0, maxPairs)
	Instance2.EnabledPairs = fillPairsFromCSV(maxPairs, maxPairs)
	Instance3.EnabledPairs = fillPairsFromCSV(maxPairs*2, maxPairs)
	Instance4.EnabledPairs = fillPairsFromCSV(maxPairs*3, maxPairs)
	Instance5.EnabledPairs = fillPairsFromCSV(maxPairs*4, maxPairs)
	Instance6.EnabledPairs = fillPairsFromCSV(maxPairs*5, maxPairs)
}

var (
	Instance1 = GoCryptoTraderConfig{
		APIKeys:        []string{"jByRuDyDvM3bQl71hLgadWt932jodjvpJRqvXsQRIWHfpSZwxYBR7BWFBOXO7o6b", "ceox9ksHbEyMQiOkSkH0k2VnFWkiojJEucGIMRUZXUq3dLCuOqV81nQIjCh0PLK5"},
		ProxyAddresses: []string{"185.186.27.185:15421", "185.186.27.184:15421"},
	}

	Instance2 = GoCryptoTraderConfig{
		APIKeys:        []string{"jByRuDyDvM3bQl71hLgadWt932jodjvpJRqvXsQRIWHfpSZwxYBR7BWFBOXO7o6b", "ceox9ksHbEyMQiOkSkH0k2VnFWkiojJEucGIMRUZXUq3dLCuOqV81nQIjCh0PLK5"},
		ProxyAddresses: []string{"185.186.27.185:15421", "185.186.27.184:15421"},
	}

	Instance3 = GoCryptoTraderConfig{
		APIKeys:        []string{"jByRuDyDvM3bQl71hLgadWt932jodjvpJRqvXsQRIWHfpSZwxYBR7BWFBOXO7o6b", "ceox9ksHbEyMQiOkSkH0k2VnFWkiojJEucGIMRUZXUq3dLCuOqV81nQIjCh0PLK5"},
		ProxyAddresses: []string{"185.186.27.185:15421", "185.186.27.184:15421"},
	}

	Instance4 = GoCryptoTraderConfig{
		APIKeys:        []string{"cUe7mZpWMNi96Gunm57USrkstCaZOuOoR7qWzFHbAIjglh0wsAgCufwqQARHX59i", "AJaBKMzpGAyYeiP7hV4aN5IwuJL7W7RtExEdrxyBh6GUf875rIiS2Ek1iZIakx4C"},
		ProxyAddresses: []string{"185.186.27.107:15421", "45.135.248.104:15421"},
	}

	Instance5 = GoCryptoTraderConfig{
		APIKeys:        []string{"cUe7mZpWMNi96Gunm57USrkstCaZOuOoR7qWzFHbAIjglh0wsAgCufwqQARHX59i", "AJaBKMzpGAyYeiP7hV4aN5IwuJL7W7RtExEdrxyBh6GUf875rIiS2Ek1iZIakx4C"},
		ProxyAddresses: []string{"185.186.27.107:15421", "45.135.248.104:15421"},
	}

	Instance6 = GoCryptoTraderConfig{
		APIKeys:        []string{"cUe7mZpWMNi96Gunm57USrkstCaZOuOoR7qWzFHbAIjglh0wsAgCufwqQARHX59i", "AJaBKMzpGAyYeiP7hV4aN5IwuJL7W7RtExEdrxyBh6GUf875rIiS2Ek1iZIakx4C"},
		ProxyAddresses: []string{"185.186.27.107:15421", "45.135.248.104:15421"},
	}
)