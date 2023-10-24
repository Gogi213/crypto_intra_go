package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
)

type Ticker struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	AskPrice string `json:"askPrice"`
}

type VolumeInfo struct {
	Symbol string  `json:"symbol"`
	Volume float64 `json:"quoteVolume,string"`
}

type MergedInfo struct {
	Symbol   string
	BidPrice string
	AskPrice string
	Volume   float64
	Status   string
}

const volumeThreshold = 300000.0  // Пороговое значение объема торгов

func UpdateCurrencyPairsToCSV() {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/bookTicker")
	if err != nil {
		log.Fatal("Error fetching data from Binance API:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
		return
	}

	var tickers []Ticker
	if err := json.Unmarshal(body, &tickers); err != nil {
		log.Fatal("Error unmarshalling JSON:", err)
		return
	}

	resp, err = http.Get("https://api.binance.com/api/v3/ticker/24hr")
	if err != nil {
		log.Fatal("Error fetching volume data from Binance API:", err)
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading volume response body:", err)
		return
	}

	var volumes []VolumeInfo
	if err := json.Unmarshal(body, &volumes); err != nil {
		log.Fatal("Error unmarshalling volume JSON:", err)
		return
	}

	var merged []MergedInfo
	for _, t := range tickers {
		for _, v := range volumes {
			if t.Symbol == v.Symbol {
				status := "Trading"
				if v.Volume == 0 {
					status = "Not Trading"
				}
				merged = append(merged, MergedInfo{t.Symbol, t.BidPrice, t.AskPrice, v.Volume, status})
				break
			}
		}
	}

	sort.Slice(merged, func(i, j int) bool {
		return merged[i].Volume > merged[j].Volume
	})

	file, err := os.Create("currency_pairs.csv")
	if err != nil {
		log.Fatal("Could not create CSV file", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"symbol", "bid", "ask", "24h volume($)", "status"})

	for _, m := range merged {
		if m.Volume <= volumeThreshold && m.Status == "Trading" {
			bidPrice, err1 := strconv.ParseFloat(m.BidPrice, 64)
			askPrice, err2 := strconv.ParseFloat(m.AskPrice, 64)
			if err1 != nil || err2 != nil {
				log.Fatal("Error parsing bid or ask price to float:", err1, err2)
				return
			}
			if bidPrice != 0 && askPrice != 0 {
				writer.Write([]string{m.Symbol, m.BidPrice, m.AskPrice, strconv.FormatFloat(m.Volume, 'f', 2, 64), m.Status})
			}
		}
	}
}
