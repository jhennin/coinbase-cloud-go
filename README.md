# Coinbase Cloud Exchange API with GoLang

This project is designed to get you started with [Coinbase Cloud Exchange API](https://docs.cloud.coinbase.com/exchange/docs) with GoLang. It provides example code for integrating with endpoints in the [Exchange API](https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getcurrency).

## How to Run the Application
Clone the repository: 

`git clone https://github.com/jhennin/coinbase-cloud-go.git`

Change to the project root directory to **run the application**:

`go run .`

## Run Tests
All tests currently live in the `coinbaseProAPI_test.go` file at the project root. 

Change directory to your project root.

`cd [PROJECT ROOT DIRECTORY]`


Run single unit test:

`go test -run Test_getACurrency`


Run all tests in package:

`go test`

