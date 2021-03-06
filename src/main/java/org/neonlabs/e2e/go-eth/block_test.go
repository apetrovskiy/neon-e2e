package go_eth

import (
	"context"
	// "fmt"
	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
	// "math/big"
	"testing"
)

func TestQueryBlockHeader(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureBlocks),
		allure.Feature(StoryQueryBlock),
		// allure.Story(StoryQueryBlock),
		allure.Description("Block header"),
		allure.Action(func() {
			client := getClient()
			header, err := client.HeaderByNumber(context.Background(), nil)
			assert.Nil(t, err, "Failed to get block header")
			ReportErrorInAllure(err)
			assert.Greater(t, 0, header.Number, "Block header number greater than 0")
		}))
}
func TestQueryFullBlock(t *testing.T) {
	t.Parallel()
	allure.Test(t,
		allure.Epic(FeatureBlocks),
		allure.Feature(StoryQueryBlock),
		// allure.Story(StoryQueryBlock),
		allure.Description("Full block"),
		allure.Action(func() {
			client := getClient()
			// blockNumber := big.NewInt(5671744)
			blockNumber := GetHighestBlockNumber(client)
			// block, err := client.BlockByNumber(context.Background(), blockNumber)
			// assert.Nil(t, err, "Failed to get full block")
			// ReportErrorInAllure(err)
			block := GetBlockByNumber(client, blockNumber)
			assert.Greater(t, 0, block.Number().Uint64())
			// assert.Greater(t, 0, block.Time().Uint64())
			assert.Greater(t, 0, block.Difficulty().Uint64())
			assert.NotNil(t, block.Hash().Hex())
			assert.Greater(t, 0, len(block.Transactions()))
		}))
}
