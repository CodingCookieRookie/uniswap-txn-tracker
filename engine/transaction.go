package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/CodingCookieRookie/uniswap-txn-tracker/env"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/log"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/model"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/mysql"
	binance_connector "github.com/binance/binance-connector-go"
)

const (
	etherscanAPI = "https://api.etherscan.io/api?module=account&action=tokentx&address=%v&page=%v&offset=%v&startblock=%v&endblock=%v&sort=%v&apikey=%v"

	GWEIsToEth = 1000000000
	ETHUSDT    = "ETHUSDT"

	txnPage                  = 1
	txnMaxBlockRange         = 1000
	txnHistoricalSort        = "desc"
	txnLiveSort              = "desc"
	TxnEarliestBlock  uint64 = 12376729

	klineMax      = 1000
	klineInterval = "1s"
)

var (
	lastLiveBlock uint64
)

// Returns a map of timestamp in seconds to eth price.
func generateTimeStampToEthPriceMapForTxns(txnResp model.TxnsResp) *map[uint64]string {
	startTimeInSec := time.Now().Unix()
	endTimeInSec := int64(0)
	for _, txn := range txnResp.Result {
		timeStampInUnixSec, err := strconv.ParseInt(txn.TimeStamp, 10, 64)
		if err != nil {
			log.Errorf("error parsing time stamp to uint64: %v", err)
			continue
		}
		if timeStampInUnixSec < startTimeInSec {
			startTimeInSec = timeStampInUnixSec
		}

		if timeStampInUnixSec > endTimeInSec {
			endTimeInSec = timeStampInUnixSec
		}
	}

	binanceConnectorClient := binance_connector.NewClient(env.BINANCE_API_KEY, env.BINANCE_API_SECRET_KEY)
	klineService := binanceConnectorClient.NewKlinesService()
	m := make(map[uint64]string)
	for i := startTimeInSec; i < endTimeInSec; i += klineMax {
		log.Debugf("called kline service for %v 1-second price data", klineMax)
		var startTimeInMS uint64 = uint64(i) * 1000
		var endTimeInMS uint64 = uint64(startTimeInMS+klineMax-1) * 1000
		klines, err := klineService.Symbol(ETHUSDT).StartTime(startTimeInMS).EndTime(endTimeInMS).Interval("1s").Limit(klineMax).Do(context.Background()) // limit at 1000 seconds
		if err != nil {
			log.Errorf("error getting klines from binance connector: %v", err)
			return nil
		}

		for _, kline := range klines {
			openTimeInSec := kline.OpenTime / 1000
			m[openTimeInSec] = kline.Open
		}
	}

	return &m
}

// Periodically inserts historical transactions.
func insertHistoricalTransactions() {
	log.Infof("lastLiveBlockNum: %v", lastLiveBlock)
	for startBlock := lastLiveBlock; startBlock >= TxnEarliestBlock; startBlock -= txnMaxBlockRange {
		log.Info("Insert Historical Transactions")
		endBlock := startBlock + txnMaxBlockRange
		etherscanAPIURL := fmt.Sprintf(etherscanAPI, env.UNISWAP_V3_CONTRACT_ADDR, txnPage, txnMaxBlockRange, startBlock, endBlock, txnHistoricalSort, env.ETHERSCAN_API_KEY)

		resp, err := http.Get(etherscanAPIURL)
		if err != nil {
			log.Errorf("error getting transactions from api, err: %v", err)
			continue
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Errorf("error reading body from api, err: %v", err)
			continue
		}
		var txnResp model.TxnsResp
		err = json.Unmarshal(body, &txnResp)
		if err != nil {
			log.Errorf("error unmarshalling body from api, err: %v", err)
			continue
		}

		timestampToEthPriceMap := generateTimeStampToEthPriceMapForTxns(txnResp)
		err = mysql.ReplaceTransactionBulkByBatch(txnResp.Result, timestampToEthPriceMap)
		if err != nil {
			log.Errorf("error inserting historical transactions into db, err: %v", err)
		} else {
			log.Info("successfully inserted historical transactions into db")
		}
		time.Sleep(time.Second * 3)
	}
}

// Periodically inserts live transactions.
func insertLiveTransactions() {
	startBlock := 0
	endBlock := math.MaxInt64
	etherscanAPIURL := fmt.Sprintf(etherscanAPI, env.UNISWAP_V3_CONTRACT_ADDR, txnPage, txnMaxBlockRange, startBlock, endBlock, txnLiveSort, env.ETHERSCAN_API_KEY)

	resp, err := http.Get(etherscanAPIURL)
	if err != nil {
		log.Errorf("error getting transactions from api, err: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("error reading body from api, err: %v", err)
		return
	}
	var txnResp model.TxnsResp
	err = json.Unmarshal(body, &txnResp)
	if err != nil {
		log.Errorf("error unmarshalling body from api, err: %v", err)
		return
	}

	if len(txnResp.Result) == 0 { // reached a block
		log.Infof("finished populating all historical transactions")
		return
	}

	if lastLiveBlock == 0 {
		lastLiveBlock, err = strconv.ParseUint(txnResp.Result[0].BlockNumber, 10, 64)
		if err != nil {
			log.Errorf("error parsing last live block number: %v", err)
		}
	}

	timestampToEthPriceMap := generateTimeStampToEthPriceMapForTxns(txnResp)
	err = mysql.ReplaceTransactionBulkByBatch(txnResp.Result, timestampToEthPriceMap)
	if err != nil {
		log.Errorf("error inserting live transactions into db, err: %v", err)
	} else {
		log.Info("successfully inserted live transactions into db")
	}
}
