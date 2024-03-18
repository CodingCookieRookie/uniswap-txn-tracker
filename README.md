# Uniswap Transaction Tracker

This project implements a Uniswap Transaction Tracker server designed to query transactions in Uniswap V3 USDC/ETH pool

## Note

Since the server operates on a pulling system, queries of data to server are strictly retrieved from Uniswap Transaction Tracker server database. This approach minimizes the risk of encountering rate limit issues due to excessive requests or potential bot attacks. Kindly exercise patience during server cold starts as data retrieval may take some time. You can use `/transactions` API along with start and end time to check which transaction data are available.
For further details, please refer to the design documentation in the link below.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.20 or later

## Configuration (Manadatory)
Using a `.env` file in `uniswap-txn-tracker` directory is mandatory to get data from third party services

### Creating a .env File
1. Create a file named `.env` in the root directory of the project.
2. Add configuration variables to the `.env` file. Below are the available variables you must configure:

```plaintext
BINANCE_API_KEY = "qqI4rRliLm8eUnu9qPj6pmvf9ZMU9r53lQpq..."
BINANCE_API_SECRET_KEY = "uTYnZHvLrhRqfDhgO4IJ..."
ETHERSCAN_API_KEY="Q2WIDXVJW4K3C2CQE6FGZZ..."
UNISWAP_V3_CONTRACT_ADDR="0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640"
INFURA_URL="wss://mainnet.infura.io/ws/v3/..."
```

### Docker run

First, clone the repository to your local machine:

Second, create a .env file with the following environments

```bash
git clone https://github.com/CodingCookieRookie/uniswap-txn-tracker.git
cd uniswap-txn-tracker
run `docker-compose up --build`
```

### Swagger APIs
1. You can find all the APIs using the swagger url once you have managed to run the application successfully
url: http://localhost:8080/swagger/index.html#/

### Design Documentation
1. You can find the technical design documentation for this app here
url: https://docs.google.com/document/d/1zH0h1reQdTv5dj2ChH6i_BcheP8FCSLQojCbK82b7bM/edit

### To run all tests
1. Run cmd `go test ./...`