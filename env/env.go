package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	LOG_FILE                 string
	BINANCE_API_KEY          string
	BINANCE_API_SECRET_KEY   string
	MYSQL_URI                string
	ETHERSCAN_API_KEY        string
	UNISWAP_V3_CONTRACT_ADDR string
)

func init() {
	godotenv.Load()
	LOG_FILE = os.Getenv("LOG_FILE")
	BINANCE_API_KEY = os.Getenv("BINANCE_API_KEY")
	BINANCE_API_SECRET_KEY = os.Getenv("BINANCE_API_SECRET_KEY")
	MYSQL_URI = os.Getenv("MYSQL_URI")
	ETHERSCAN_API_KEY = os.Getenv("ETHERSCAN_API_KEY")
	UNISWAP_V3_CONTRACT_ADDR = os.Getenv("UNISWAP_V3_CONTRACT_ADDR")
}
