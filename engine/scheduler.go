package engine

import (
	"math/rand"
	"time"
)

// Ensure live transaction runs once to get last block.
func RunInsertScheduler() {
	insertLiveTransactions()
	go insertHistoricalTransactions()
	go insertHistoricalSwapPrices()
	go insertLiveSwapPrices()
}

func RunInsertLiveTransactions() {
	for {
		insertLiveTransactions()
		time.Sleep(generateRandomJitter(45, 60) * time.Second)
	}
}

// Generates random jitter time between min and max seconds for goroutine to sleep to reduce server load at one instant.
func generateRandomJitter(minSeconds, maxSeconds int) time.Duration {
	rand.NewSource(time.Now().UnixNano())
	jitterSeconds := minSeconds + rand.Intn(maxSeconds-minSeconds+1)
	return time.Duration(jitterSeconds) * time.Second
}
