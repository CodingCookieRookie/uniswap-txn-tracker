package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeiToEth(t *testing.T) {
	tests := []struct {
		name     string
		gasPrice uint64
		expected float64
	}{
		{"Convert 1 ETH in wei", 1000000000000000000, 1.0},
		{"Convert 0 wei", 0, 0.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := weiToEth(test.gasPrice)
			assert.Equal(t, test.expected, result)
		})
	}
}
