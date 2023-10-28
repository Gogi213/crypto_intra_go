// config.go
package main

type GoCryptoTraderConfig struct {
    APIKeys        []string `json:"apiKeys"`
    ProxyAddresses []string `json:"proxyAddresses"`
    EnabledPairs   []string `json:"enabledPairs"`
    MaxPairs       int      `json:"maxPairs"`
}

// Предположим, что exchangePairsList обновляется при каждом запуске сервера
var exchangePairsList = []string{
    // ваш список валютных пар
}

func fillPairsFromDataFrame(instance *GoCryptoTraderConfig) {
    if len(exchangePairsList) > instance.MaxPairs {
        instance.EnabledPairs = exchangePairsList[:instance.MaxPairs]
    } else {
        instance.EnabledPairs = exchangePairsList
    }
}

var (
    Instance1 = GoCryptoTraderConfig{
        APIKeys:        []string{"jByRuDyDvM3bQl71hLgadWt932jodjvpJRqvXsQRIWHfpSZwxYBR7BWFBOXO7o6b", "ceox9ksHbEyMQiOkSkH0k2VnFWkiojJEucGIMRUZXUq3dLCuOqV81nQIjCh0PLK5"},
        ProxyAddresses: []string{"185.186.27.185:15421", "185.186.27.184:15421"},
        MaxPairs:       100,
    }
    Instance2 = GoCryptoTraderConfig{
        APIKeys:        []string{"jByRuDyDvM3bQl71hLgadWt932jodjvpJRqvXsQRIWHfpSZwxYBR7BWFBOXO7o6b", "ceox9ksHbEyMQiOkSkH0k2VnFWkiojJEucGIMRUZXUq3dLCuOqV81nQIjCh0PLK5"},
        ProxyAddresses: []string{"185.186.27.185:15421", "185.186.27.184:15421"},
        MaxPairs:       100,
    }
    Instance3 = GoCryptoTraderConfig{
        APIKeys:        []string{"jByRuDyDvM3bQl71hLgadWt932jodjvpJRqvXsQRIWHfpSZwxYBR7BWFBOXO7o6b", "ceox9ksHbEyMQiOkSkH0k2VnFWkiojJEucGIMRUZXUq3dLCuOqV81nQIjCh0PLK5"},
        ProxyAddresses: []string{"185.186.27.185:15421", "185.186.27.184:15421"},
        MaxPairs:       100,
    }
    Instance4 = GoCryptoTraderConfig{
        APIKeys:        []string{"cUe7mZpWMNi96Gunm57USrkstCaZOuOoR7qWzFHbAIjglh0wsAgCufwqQARHX59i", "AJaBKMzpGAyYeiP7hV4aN5IwuJL7W7RtExEdrxyBh6GUf875rIiS2Ek1iZIakx4C"},
        ProxyAddresses: []string{"185.186.27.107:15421", "45.135.248.104:15421"},
        MaxPairs:       100,
    }
    Instance5 = GoCryptoTraderConfig{
        APIKeys:        []string{"cUe7mZpWMNi96Gunm57USrkstCaZOuOoR7qWzFHbAIjglh0wsAgCufwqQARHX59i", "AJaBKMzpGAyYeiP7hV4aN5IwuJL7W7RtExEdrxyBh6GUf875rIiS2Ek1iZIakx4C"},
        ProxyAddresses: []string{"185.186.27.107:15421", "45.135.248.104:15421"},
        MaxPairs:       100,
    }
    Instance6 = GoCryptoTraderConfig{
        APIKeys:        []string{"cUe7mZpWMNi96Gunm57USrkstCaZOuOoR7qWzFHbAIjglh0wsAgCufwqQARHX59i", "AJaBKMzpGAyYeiP7hV4aN5IwuJL7W7RtExEdrxyBh6GUf875rIiS2Ek1iZIakx4C"},
        ProxyAddresses: []string{"185.186.27.107:15421", "45.135.248.104:15421"},
        MaxPairs:       100,
    }
)

func init() {
    fillPairsFromDataFrame(&Instance1)
    fillPairsFromDataFrame(&Instance2)
    fillPairsFromDataFrame(&Instance3)
    fillPairsFromDataFrame(&Instance4)
    fillPairsFromDataFrame(&Instance5)
    fillPairsFromDataFrame(&Instance6)
}
