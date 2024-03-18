package engine

import "time"

// Ensure live transaction runs once to get last block.
func RuncInsertHistoricalTransactions() {
	insertLiveTransactions()
	go insertHistoricalTransactions()
	go insertSwapPrices()
}

func RunInsertLiveTransactions() {
	for {
		insertLiveTransactions()
		time.Sleep(30 * time.Second)
	}
}
