package engine

import (
	"context"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/CodingCookieRookie/uniswap-txn-tracker/env"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/log"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/model"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/mysql"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	uniswapContractABIByte = `[ { "anonymous": false, "inputs": [ {"indexed": true, "internalType": "address", "name": "sender", "type": "address"}, {"indexed": true, "internalType": "address", "name": "recipient", "type": "address"}, {"indexed": false, "internalType": "int256", "name": "amount0", "type": "int256"}, {"indexed": false, "internalType": "int256", "name": "amount1", "type": "int256"}, {"indexed": false, "internalType": "uint160", "name": "sqrtPriceX96", "type": "uint160"}, {"indexed": false, "internalType": "uint128", "name": "liquidity", "type": "uint128"}, {"indexed": false, "internalType": "int24", "name": "tick", "type": "int24"} ], "name": "Swap", "type": "event" }]`
	eventSignature         = "Swap(address,address,int256,int256,uint160,uint128,int24)"
	swapBlockRangeSize     = 5000
)

var (
	client             *ethclient.Client
	uniswapContractABI abi.ABI
	signatureHash      common.Hash
)

func init() {
	var err error
	uniswapContractABI, err = abi.JSON(strings.NewReader(uniswapContractABIByte))
	if err != nil {
		log.Errorf("error deserialising contract: %v", err)
		return
	}
	signatureHash = crypto.Keccak256Hash([]byte(eventSignature))
	initEthClient()
}

func initEthClient() {
	var err error
	client, err = ethclient.Dial(env.INFURA_URL)
	if err != nil {
		log.Errorf("failed to connect to the Ethereum client: %v", err)
		return
	}
}

// has to get lastLiveBlock first
func insertHistoricalSwapPrices() {
	var once sync.Once
	for i := lastLiveBlock; i >= TxnEarliestBlock; i -= swapBlockRangeSize {
		log.Debugf("i: %d", i)
		query := ethereum.FilterQuery{
			Addresses: []common.Address{common.HexToAddress(env.UNISWAP_V3_CONTRACT_ADDR)},
			Topics:    [][]common.Hash{{signatureHash}},
			FromBlock: big.NewInt(int64(i - swapBlockRangeSize)),
			ToBlock:   big.NewInt(int64(i)),
		}
		logs, err := client.FilterLogs(context.Background(), query)

		if err != nil {
			log.Errorf("error filtering logs: %v", err)
			once.Do(initEthClient)  // retry initEthClient once for each storeSwapPrices
			i += swapBlockRangeSize // reinsert failed swap block range
		}

		swaps := make([]*model.SwapEvent, 0, len(logs))
		for _, vLog := range logs {
			if len(vLog.Topics) < 3 {
				continue
			}
			from := common.HexToAddress(vLog.Topics[1].Hex()).String()
			to := common.HexToAddress(vLog.Topics[2].Hex()).String()

			eventData, err := uniswapContractABI.Unpack("Swap", vLog.Data)
			if err != nil {
				log.Errorf("error unpacking event data: %v", err)
				continue
			}

			vLogTxnHash := vLog.TxHash

			if len(eventData) != 5 { // error in event data
				log.Error("swap event data does not have 5 fields")
				continue
			}
			sqrtPriceX96 := eventData[2].(*big.Int).String()
			swaps = append(swaps, &model.SwapEvent{
				TxnHash:      vLogTxnHash.String(),
				SqrtPriceX96: sqrtPriceX96,
				From:         from,
				To:           to,
			})
		}
		err = mysql.ReplaceSwapBulkByBatch(swaps)
		if err != nil {
			log.Errorf("error replacing swap bulk: %v", err)
			time.Sleep(generateRandomJitter(5, 8) * time.Second)
			continue
		}
		log.Debugf("successfully inserted swap bulk")
		time.Sleep(generateRandomJitter(5, 8) * time.Second)
	}
}

func insertLiveSwapPrices() {
	prevLastBlock := lastLiveBlock
	for {
		query := ethereum.FilterQuery{
			Addresses: []common.Address{common.HexToAddress(env.UNISWAP_V3_CONTRACT_ADDR)},
			Topics:    [][]common.Hash{{signatureHash}},
			FromBlock: big.NewInt(int64(prevLastBlock)),
			ToBlock:   big.NewInt(int64(prevLastBlock + swapBlockRangeSize)),
		}
		logs, err := client.FilterLogs(context.Background(), query)

		if err != nil {
			time.Sleep(generateRandomJitter(5, 7) * time.Second)
			continue
		}
		swaps := make([]*model.SwapEvent, 0, len(logs))
		currentMaxLastBlock := prevLastBlock
		for _, vLog := range logs {
			if len(vLog.Topics) < 3 {
				continue
			}
			from := common.HexToAddress(vLog.Topics[1].Hex()).String()
			to := common.HexToAddress(vLog.Topics[2].Hex()).String()

			eventData, err := uniswapContractABI.Unpack("Swap", vLog.Data)
			if err != nil {
				log.Errorf("error unpacking event data: %v", err)
				continue
			}

			vLogTxnHash := vLog.TxHash

			if len(eventData) != 5 { // error in event data
				log.Error("swap event data does not have 5 fields")
				continue
			}
			sqrtPriceX96 := eventData[2].(*big.Int).String()
			swaps = append(swaps, &model.SwapEvent{
				TxnHash:      vLogTxnHash.String(),
				SqrtPriceX96: sqrtPriceX96,
				From:         from,
				To:           to,
			})
			currentMaxLastBlock = max(currentMaxLastBlock, vLog.BlockNumber)
		}
		err = mysql.ReplaceSwapBulkByBatch(swaps)
		if err != nil {
			log.Errorf("error replacing swap bulk: %v", err)
			time.Sleep(generateRandomJitter(5, 7) * time.Second)
			continue
		}
		log.Debugf("successfully inserted swap bulk")
		prevLastBlock = currentMaxLastBlock
		time.Sleep(generateRandomJitter(5, 7) * time.Second)
	}
}
