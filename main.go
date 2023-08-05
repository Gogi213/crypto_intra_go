// main.go
package main

import (
	"log"
	"net/http"
	"os"
)

type Config struct {
	APIKeys        []string
	ProxyAddresses []string
	PairsList      [][]string
}

func main() {
	dataChannel1 := make(chan []byte, 20000)
	dataChannel2 := make(chan []byte, 20000)

	// Start profiling server
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "instance1":
			go StartGin(dataChannel1, dataChannel2) // Запуск Gin только для instance1
			StartPulsar(dataChannel1, Instance1)
		case "instance2":
			StartPulsar(dataChannel2, Instance2)
		default:
			log.Fatal("Unknown instance. Please provide either 'instance1' or 'instance2'.")
		}
	} else {
		log.Fatal("Please provide an instance argument (instance1/instance2)")
	}
}
