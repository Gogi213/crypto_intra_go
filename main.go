// main.go
package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {
	dataChannel := make(chan []byte, 20000)

	// Start profiling server
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	// Запускаем функции из других файлов
	go StartGin(dataChannel) // из файла app.go
	StartPulsar(dataChannel) // из файла binance_pulsar.go
}
