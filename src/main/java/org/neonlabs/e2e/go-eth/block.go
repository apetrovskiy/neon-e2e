package go_eth

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetHighestBlockNumber(client *ethclient.Client) *big.Int {
	header, err := client.HeaderByNumber(context.Background(), nil)
	ReportErrorInAllure(err)
	return header.Number
}

func GetBlockByNumber(client *ethclient.Client, blockNumber *big.Int) *types.Block {
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	ReportErrorInAllure(err)
	return block
}
