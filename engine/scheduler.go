package engine

import "time"

// Ensure live transaction runs once to get last block.
func RunInsertScheduler() {
	insertLiveTransactions()
	go insertHistoricalTransactions()
	go insertSwapPrices()
}

func RunInsertLiveTransactions() {
	for {
		insertLiveTransactions()
		time.Sleep(45 * time.Second)
	}
}
