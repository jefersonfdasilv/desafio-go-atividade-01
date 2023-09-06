package storage

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

var db *sql.DB

func InitDatabase() {
	var err error
	db, err = sql.Open("sqlite3", "exchange_rate.db")
	if err != nil {
		log.Fatal(err)
	}
	createTable()
}

func createTable() {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS exchange_rates (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            currency_from TEXT NOT NULL,
            currency_to TEXT NOT NULL,
            rate REAL NOT NULL,
            timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
        );
    `)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertExchangeRate(currencyFrom, currencyTo string, rate float64) {
	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Millisecond)
	defer cancel()
	_, err := db.ExecContext(ctx, "INSERT INTO exchange_rates (currency_from, currency_to, rate) VALUES (?, ?, ?)", currencyFrom, currencyTo, rate)
	if err != nil {
		log.Fatal(err)
	}
	select {
	case <-ctx.Done():
		log.Println("It seams our database is really slow or very busy", ctx.Err())
		return
	default:
		log.Println("Insert has ben completed on time")
	}
}

func GetLatestExchangeRate(currencyFrom, currencyTo string) (float64, error) {
	var rate float64
	err := db.QueryRow("SELECT rate FROM exchange_rates WHERE currency_from = ? AND currency_to = ? ORDER BY timestamp DESC LIMIT 1", currencyFrom, currencyTo).Scan(&rate)
	if err != nil {
		return 0, err
	}
	return rate, nil
}
