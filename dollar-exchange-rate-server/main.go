package main

import (
	"context"
	"encoding/json"
	"github.com/jeferson-fs/dollar-exchange-rate-server/api"
	"github.com/jeferson-fs/dollar-exchange-rate-server/storage"
	"log"
	"net/http"
	"time"
)

const (
	maxApiTimeout = 200 * time.Millisecond
)

func main() {
	storage.InitDatabase()

	http.HandleFunc("/cotacao", getExchangeRateHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getExchangeRateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), maxApiTimeout)
	defer cancel()
	rate, err := api.FetchExchangeRate(ctx)
	if err != nil {
		http.Error(w, "Failed to fetch exchange rate", http.StatusInternalServerError)
		return
	}

	response := map[string]float64{"rate": rate}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}
