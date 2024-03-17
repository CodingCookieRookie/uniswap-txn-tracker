package main

import (
	"github.com/CodingCookieRookie/uniswap-txn-tracker/api"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/log"
)

func main() {
	log.Info("Starting Uniswap Transaction Tracker")
	api.InitRouter()
}
