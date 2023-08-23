// main.go
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	dataChannels1 := []chan []byte{
		make(chan []byte, 300000),
		make(chan []byte, 300000),
		make(chan []byte, 300000),
	}
	dataChannels2 := []chan []byte{
		make(chan []byte, 300000),
		make(chan []byte, 300000),
		make(chan []byte, 300000),
	}

	for _, ch := range append(dataChannels1, dataChannels2...) {
		defer close(ch)
	}

	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "server":
			go StartServer([][]chan []byte{dataChannels1, dataChannels2}, []string{"50000", "50001", "50002"}, []string{"60000", "60001", "60002"})
		case "instance1":
			go StartPulsar(dataChannels1, Instance1, []string{"50000", "50001", "50002"})
		case "instance2":
			go StartPulsar(dataChannels2, Instance2, []string{"60000", "60001", "60002"})
		default:
			log.Fatalf("Unknown command. Please provide either 'server', 'instance1', or 'instance2'. Got: %s", args[1])
		}
	} else {
		log.Fatal("Please provide a command argument (server/instance1/instance2)")
	}

	// Wait for user input to stop the program
	var input string
	fmt.Scanln(&input)
}

func writeDataToCSV(data []string) {
	file, err := os.OpenFile("data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Unable to open or create file: ", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(data)
	if err != nil {
		log.Fatal("Unable to write data to file: ", err)
	}
}
