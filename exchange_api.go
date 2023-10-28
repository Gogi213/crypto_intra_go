// exchange_api.go
package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/cheynewallace/tabby"
	"github.com/thrasher-corp/gocryptotrader/exchanges/binance"
)

// ExchangePairsList - глобальная переменная для хранения списка пар
var ExchangePairsList []string

type MergedInfo struct {
	Symbol   string
	BidPrice float64
	AskPrice float64
	Volume   float64
	Status   string
}

func UpdateCurrencyPairsToCSV() {
	b := binance.Binance{}
	b.SetDefaults()

	tickers, err := b.GetTickers(context.Background())
	if err != nil {
		log.Fatal("Error fetching data from Binance API:", err)
		return
	}

	var merged []MergedInfo
	for _, t := range tickers {
		status := "Trading"
		if t.LastPrice == 0 {
			status = "Not Trading"
		}
		merged = append(merged, MergedInfo{t.Symbol, t.BidPrice, t.AskPrice, t.Volume, status})
	}

	file, err := os.Create("currency_pairs.csv")
	if err != nil {
		log.Fatal("Could not create CSV file", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"symbol", "bid", "ask", "24h volume($)", "status"})

	t := tabby.New()
	t.AddHeader("SYMBOL", "BID", "ASK", "24H VOLUME($)", "STATUS")

	for _, m := range merged {
		if m.Status == "Trading" {
			if m.BidPrice != 0 && m.AskPrice != 0 {
				row := []string{m.Symbol, strconv.FormatFloat(m.BidPrice, 'f', 2, 64), strconv.FormatFloat(m.AskPrice, 'f', 2, 64), strconv.FormatFloat(m.Volume, 'f', 2, 64), m.Status}
				writer.Write(row)

				// Используем глобальную переменную ExchangePairsList
				ExchangePairsList = append(ExchangePairsList, m.Symbol)

				t.AddLine(m.Symbol, strconv.FormatFloat(m.BidPrice, 'f', 2, 64), strconv.FormatFloat(m.AskPrice, 'f', 2, 64), strconv.FormatFloat(m.Volume, 'f', 2, 64), m.Status)
			}
		}
	}

	// Сортируем ExchangePairsList
	sort.Sort(sort.Reverse(sort.StringSlice(ExchangePairsList)))

	// Выводим таблицу в консоль
	t.Print()
}
