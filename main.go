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
	dataChannel := make(chan []byte, 20000)

	// Start profiling server
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "instance1":
			go StartGin(dataChannel) // Запуск Gin только для instance1
			StartPulsar(dataChannel, Instance1)
		case "instance2":
			StartPulsar(dataChannel, Instance2)
		default:
			log.Fatal("Unknown instance. Please provide either 'instance1' or 'instance2'.")
		}
	} else {
		log.Fatal("Please provide an instance argument (instance1/instance2)")
	}
}
