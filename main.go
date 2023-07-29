// main.go
package main

import "log"

func main() {
	log.Println("Starting application...")

	// Запускаем функции из других файлов
	go StartGin() // из файла app.go
	StartApp()    // из файла binance_pulsar.go
}
