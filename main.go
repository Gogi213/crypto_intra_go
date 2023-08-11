// main.go
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Config struct {
	APIKeys        []string
	ProxyAddresses []string
	PairsList      [][]string
}

func main() {
	dataChannel1 := make(chan []byte, 300000)
	dataChannel2 := make(chan []byte, 300000)
	defer close(dataChannel1)
	defer close(dataChannel2)

	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "server":
			go StartServer([]chan []byte{dataChannel1, dataChannel2}, "12345", "12346")
		case "instance1":
			go StartPulsar(dataChannel1, Instance1, "12345")
		case "instance2":
			go StartPulsar(dataChannel2, Instance2, "12346")
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