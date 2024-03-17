package main

import (
	"github.com/CodingCookieRookie/uniswap-txn-tracker/api"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/log"
)

//	@title			Uniswap Transaction Tracker
//	@version		1.0
//	@description	Server for uniswap transactions operations.

//	@contact.name	Alvin Chee
//	@contact.url	https://github.com/CodingCookieRookie
//	@contact.email	alvinchee98@gmail.com

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	log.Info("Starting Uniswap Transaction Tracker")
	api.InitRouter()
}
