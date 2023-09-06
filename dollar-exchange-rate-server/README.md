# Dollar Exchange Rate Server
## Current exchange rate US DOLLAR (USD) to BRAZIL REAL (BRL) 

It's a simple Go Lang project to fetch and display the current exchange rate from US DOLLAR (USD) to BRAZIL REAL (BRL). The project will involve making an API request to a reliable exchange rate API and displaying the result in a user-friendly format.

### Project directories
```
./
|-- main.go
|-- go.mod
|-- go.sum
|-- api/
|   |-- exchange.go
|-- storage/
|   |-- database.go
|-- README.md

```

* main.go: This is the entry point of the application where you'll start the program and call necessary functions.

* go.mod: The Go module file that keeps track of your project's dependencies.

* go.sum: A checksum file for the dependencies.

* api/exchange.go: This package contains functions for interacting with the exchange rate API. It includes making HTTP requests, parsing JSON responses, and extracting the exchange rate value.

* storage/database.go: This package would contain functions for interacting with the SQLite database. It includes functions to create tables, insert data (exchange rates), and retrieve data from the database.

