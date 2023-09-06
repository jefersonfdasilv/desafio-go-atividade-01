package api

import (
	"context"
	"encoding/json"
	"github.com/jeferson-fs/dollar-exchange-rate-server/storage"
	"net/http"
	"strconv"
	"time"
)

const exchangeRateURL = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
const maxDatabaseTimeout = 20 * time.Millisecond

type ExchangeRateResponse struct {
	USDBRL struct {
		Bid       string `json:"bid"`
		Timestamp string `json:"timestamp"`
	} `json:"USDBRL"`
}

func FetchExchangeRate(ctx context.Context) (float64, error) {
	req, err := http.NewRequest("GET", exchangeRateURL, nil)
	if err != nil {
		return 0, err
	}

	// Associando o contexto à solicitação
	req = req.WithContext(ctx)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	// Verificando se o contexto foi cancelado devido ao prazo
	select {
	case <-ctx.Done():
		return 0, ctx.Err() // Retornando o erro de cancelamento
	default:
		{
			var responseData ExchangeRateResponse
			err = json.NewDecoder(response.Body).Decode(&responseData)
			if err != nil {
				return 0, err
			}

			rate, err := strconv.ParseFloat(responseData.USDBRL.Bid, 64)
			if err != nil {
				return 0, err
			}

			storage.InsertExchangeRate("USD", "BRL", rate)

			return rate, nil

		}
	}
}
