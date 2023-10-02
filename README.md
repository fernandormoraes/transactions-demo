# Transactions Demo

A simple application who stores transactions and retrieve with converted amount using Department of the Treasury API.

## Run application

    go run .\cmd\main.go

## Endpoints examples

Some endpoints examples are presented below.

### POST a new transaction

    curl --request POST \
    --url http://localhost:28232/api/v1/transactions \
    --header 'Content-Type: application/json' \
    --data '{
        "date": "2023-08-10 21:41:00",
        "amount": 23.32,
        "description": "Test Description"
    }'

*Preview*

    {
        "CreatedAt": "2023-10-02T20:13:33.6935854-03:00",
        "UpdatedAt": "2023-10-02T20:13:33.6935854-03:00",
        "DeletedAt": null,
        "ID": "bffd759d-597a-4da2-9a51-b1f439e3eee0",
        "Description": "Test TEXTO ",
        "Date": "2023-05-02T21:41:00Z",
        "Amount": 23.32
    }

### Retrieve transaction

Receives a transaction date in YYYY-MM-dd format and currency description in given API format (https://fiscaldata.treasury.gov/api-documentation/)

    curl --request GET \
    --url 'http://localhost:28232/api/v1/transactions?transactionDate=2023-08-10&currencyDesc=Brazil-Real'

*Preview*    

    [
        {
            "Description": "Test TEXTO ",
            "Date": "2023-05-02T21:41:00Z",
            "Amount": 23.32,
            "ConvertedAmount": 123.27,
            "ExchangeRate": 5.286
        }
    ]

## TODO

- Swagger Docs
- Add test cases