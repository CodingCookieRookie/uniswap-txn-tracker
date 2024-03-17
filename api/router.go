package api

import (
	"fmt"
	"net/http"

	"github.com/CodingCookieRookie/uniswap-txn-tracker/errors"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("transaction/fee", ginResponseWithError(GetTransactionFee))
		v1.GET("transactions", ginResponseWithError(GetHistoricalTxns))
	}
	r.Run()
}

func generateStatusMsg(err error) string {
	if err == nil {
		return fmt.Sprintf("%v OK", http.StatusOK)
	} else if _, ok := err.(*errors.UserError); ok {
		return fmt.Sprintf("%v Bad Request Error: %v", http.StatusBadRequest, err.Error())
	} else if _, ok := err.(*errors.ServerError); ok {
		return fmt.Sprintf("%v Interval Server Error", http.StatusInternalServerError)
	} else {
		return fmt.Sprintf("520 Unknown Error: %v", err)
	}
}

func ginResponseWithError(f func(ctx *gin.Context) (any, error)) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		resp, err := f(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"statusMsg": generateStatusMsg(err),
				"body":      resp,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"statusMsg": generateStatusMsg(err),
				"body":      resp,
			})
		}
	}
}
