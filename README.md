# Uniswap Transaction Tracker

This project implements a Uniswap Transaction Tracker server designed to query transactions in Uniswap V3 USDC/ETH pool

## Note

Since the server operates on a pulling system, queries of data to server are strictly retrieved from Uniswap Transaction Tracker server database. This approach minimizes the risk of encountering rate limit issues due to excessive requests or potential bot attacks. Kindly exercise patience during server cold starts as data retrieval may take some time. You can use `/transactions` API along with start and end time to check which transaction data are available.
For further details, please refer to the design documentation in the link below.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
- Docker
- Go 1.21 (For running go tests locally)

### Configuration (MANDATORY)
Using a `.env` file in `uniswap-txn-tracker` directory is mandatory to get data from third party services

#### Steps to Create and Configure `.env` File:
   - Navigate to the root directory of your project
   - Create a new file and name it `.env`
   - Open the `.env` file you just created
   - Copy and paste the following configuration variables into the `.env` file. Ensure to replace the placeholder values with your actual API keys and URLs

```plaintext
UNISWAP_V3_CONTRACT_ADDR="0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640" # Uniswap V3: USDC 3 contract address
BINANCE_API_KEY = "qqI4rRliLm8eUnu9qPj6pmvf9ZMU9r53lQpq..." # Your Binance API Key Here
BINANCE_API_SECRET_KEY = "uTYnZHvLrhRqfDhgO4IJ..." # Your Binance Secret Key Here"
ETHERSCAN_API_KEY="Q2WIDXVJW4K3C2CQE6FGZZ..." # "Your Etherscan API Key Here"
INFURA_URL="wss://mainnet.infura.io/ws/v3/eed3d2afcd..." "Infura Endpoint URL with your Infura API Key Here"
```

#### Obtaining API Keys:

- **BINANCE_API_KEY and BINANCE_API_SECRET_KEY**: [Binance API Key Creation Guide](https://www.binance.com/en/support/faq/how-to-create-api-keys-on-binance-360002502072)
- **ETHERSCAN_API_KEY**: [Etherscan API Documentation](https://docs.etherscan.io/getting-started/viewing-api-usage-statistics)
- **INFURA API KEY**: [Infura API Documentation](https://docs.infura.io/api/getting-started)

#### Need Help?

If you encounter any issues or require further assistance with the setup, please do not hesitate to contact the developer at `alvinchee98@gmail.com`.

### Docker run

First, clone the repository to your local machine

Second, go into uniswap-txn-tracker directory

Third, create and configure the .env file with the above mandatory environment variables

Fourth, run 'docker-compose up --build'

```bash
git clone https://github.com/CodingCookieRookie/uniswap-txn-tracker.git
cd uniswap-txn-tracker
touch .env
docker-compose up --build
```

### Swagger APIs
1. You can find all the APIs using the swagger url once you have managed to run the application successfully
url: http://localhost:8080/swagger/index.html#/

### Design Documentation
1. You can find the technical design documentation for this app here
url: https://docs.google.com/document/d/1zH0h1reQdTv5dj2ChH6i_BcheP8FCSLQojCbK82b7bM/edit

### To run all tests
1. Run cmd `go test ./...`