package main

import "log"

func main() {
	log.Println("Starting application...")

	// Запускаем функции из других файлов
	go StartGin() // из файла app.go
	StartPulsar() // из файла binance_pulsar.go
}
