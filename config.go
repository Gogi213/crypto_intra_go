package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	APIKeys []string `json:"api_keys"`
	IPs     []string `json:"ips"`
	List1   []string `json:"list1"`
	List2   []string `json:"list2"`
	List3   []string `json:"list3"`
}

func loadConfig(configFile string) Config {
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal("Cannot open config file", err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Cannot get configuration from file", err)
	}

	return config
}
