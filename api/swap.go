package api

import (
	"github.com/CodingCookieRookie/uniswap-txn-tracker/service"
	"github.com/gin-gonic/gin"
)

// GetUniswapSwapPrice godoc
// @Summary		returns uniswap swap price with corresponding transaction hash
// @Description	returns uniswap swap price with corresponding transaction hash
// @Tags			accounts
// @Accept			json
// @Produce		json
// @Param			txn_hash		query		string	true	"Transaction hash"
// @Success		200			{object}	model.SwapResponse
// @Failure		500			{object}	errors.ServerError	"Server Error"
// @Router			/swap [get]
func GetUniswapSwapPrice(c *gin.Context) (any, error) {
	txnHash := c.Query("txn_hash")
	return service.GetUniswapSwapPrice(txnHash)
}
