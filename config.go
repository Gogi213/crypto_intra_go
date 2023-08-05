// config.go
package main

var Instance1 = Config{
	APIKeys:        []string{"jByRuDyDvM3bQl71hLgadWt932jodjvpJRqvXsQRIWHfpSZwxYBR7BWFBOXO7o6b", "ceox9ksHbEyMQiOkSkH0k2VnFWkiojJEucGIMRUZXUq3dLCuOqV81nQIjCh0PLK5"},
	ProxyAddresses: []string{"212.115.48.122:13225", "195.225.96.39:13225"},
	PairsList: [][]string{
		{"1inchbtc", "1inchbusd", "1inchusdt"},
		{"atomusdt", "auctionbtc", "auctionbusd"},
	},
}

var Instance2 = Config{
	APIKeys:        []string{"cUe7mZpWMNi96Gunm57USrkstCaZOuOoR7qWzFHbAIjglh0wsAgCufwqQARHX59i", "AJaBKMzpGAyYeiP7hV4aN5IwuJL7W7RtExEdrxyBh6GUf875rIiS2Ek1iZIakx4C"},
	ProxyAddresses: []string{"193.160.211.65:13225", "45.86.3.128:13225"},
	PairsList: [][]string{
		{"btcusdt", "ethusdt", "btsusdt"},
		{"cakeusdt", "compusdt", "cotiusdt"},
	},
}
