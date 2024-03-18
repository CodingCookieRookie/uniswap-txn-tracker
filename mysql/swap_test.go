package mysql

import (
	"testing"

	"github.com/CodingCookieRookie/uniswap-txn-tracker/model"
	"github.com/stretchr/testify/assert"
)

func TestGenerateUniqueSwapId(t *testing.T) {
	t.Run("nil input", func(t *testing.T) {
		result := generateUniqueSwapId(nil, 10)
		assert.Empty(t, result)
	})
	t.Run("valid input with specific length", func(t *testing.T) {
		swap := &model.SwapEvent{
			From:    "0xFromAddress",
			To:      "0xToAddress",
			TxnHash: "0xTxnHash",
		}
		length := 8 // Arbitrary length for testing
		result := generateUniqueSwapId(swap, length)
		assert.Len(t, result, length)
	})

	t.Run("length exceeds encoded size", func(t *testing.T) {
		swap := &model.SwapEvent{
			From:    "0xFromAddress",
			To:      "0xToAddress",
			TxnHash: "0xTxnHash",
		}
		excessiveLength := 1000 // Arbitrary large length
		result := generateUniqueSwapId(swap, excessiveLength)
		assert.NotEmpty(t, result)
		fullLengthResult := generateUniqueSwapId(swap, 0)
		assert.Equal(t, fullLengthResult, result)
	})

	t.Run("negative length value", func(t *testing.T) {
		swap := &model.SwapEvent{
			From:    "0xFromAddress",
			To:      "0xToAddress",
			TxnHash: "0xTxnHash",
		}
		negativeLength := -1
		result := generateUniqueSwapId(swap, negativeLength)
		assert.NotEmpty(t, result)
		fullLengthResult := generateUniqueSwapId(swap, 0)
		assert.Equal(t, fullLengthResult, result)
	})
}
