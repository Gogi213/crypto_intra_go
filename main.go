// main.go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create data channels for 6 instances
	dataChannels := make([][]chan []byte, 6)
	for i := 0; i < 6; i++ {
		dataChannels[i] = []chan []byte{
			make(chan []byte, 300000),
			make(chan []byte, 300000),
			make(chan []byte, 300000),
		}
	}

	// Close all channels when main function exits
	for _, chs := range dataChannels {
		for _, ch := range chs {
			defer close(ch)
		}
	}

	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "server":
			go StartServer(dataChannels, []string{"5000", "5001", "5002", "6000", "6001", "6002", "7000", "7001", "7002", "8000", "8001", "8002", "9000", "9001", "9002", "10000", "10001", "10002"})
			UpdateCurrencyPairsToCSV()
		case "instance1":
			go StartPulsar(dataChannels[0], Instance1, []string{"5000", "5001", "5002"})
		case "instance2":
			go StartPulsar(dataChannels[1], Instance2, []string{"6000", "6001", "6002"})
		case "instance3":
			go StartPulsar(dataChannels[2], Instance3, []string{"7000", "7001", "7002"})
		case "instance4":
			go StartPulsar(dataChannels[3], Instance4, []string{"8000", "8001", "8002"})
		case "instance5":
			go StartPulsar(dataChannels[4], Instance5, []string{"9000", "9001", "9002"})
		case "instance6":
			go StartPulsar(dataChannels[5], Instance6, []string{"10000", "10001", "10002"})
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
