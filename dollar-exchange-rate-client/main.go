package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type ExchangeRateResponse struct {
	Rate float64 `json:"rate"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequest("GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	req = req.WithContext(ctx)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing the body:", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Status Code %d\n", resp.StatusCode)
		return
	}
	var exchangeRate ExchangeRateResponse
	err = json.NewDecoder(resp.Body).Decode(&exchangeRate)
	if err != nil {
		fmt.Println("Error reading/decoding response body:", err)
		return
	}

	fmt.Println("Exchange rate value:", exchangeRate.Rate)
	_ = saveToFile("cotacao.txt", exchangeRate.Rate)
}

func saveToFile(filename string, value float64) error {
	content := fmt.Sprintf("Dólar: %.2f", value)

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(file)

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
