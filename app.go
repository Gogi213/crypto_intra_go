package main

import (
	"github.com/mailru/easyjson"
	"log"
)

type BinanceMessage struct {
	Stream string  `json:"stream"`
	Data   Message `json:"data"`
}

type Message struct {
	S  string `json:"s"`
	B  string `json:"b"`
	BB string `json:"B"`
	A  string `json:"a"`
	AA string `json:"A"`
}

func StartPulsar() {
	// Here you should replace with your own data source
	// For example, you can use a channel to receive data
	dataChannel := make(chan []byte)

	for dataBytes := range dataChannel {
		var msg BinanceMessage
		if err := easyjson.Unmarshal(dataBytes, &msg); err != nil {
			log.Fatal("Failed to unmarshal data: ", err)
		}

		// Convert the Message to JSON bytes
		jsonData, err := easyjson.Marshal(msg.Data)
		if err != nil {
			log.Fatal("Failed to marshal data: ", err)
		}

		// Send the JSON bytes to the global dataChannel
		dataChannel <- jsonData
	}
}
