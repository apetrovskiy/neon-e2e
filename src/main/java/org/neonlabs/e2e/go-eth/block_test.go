package go_eth

import (
	"context"
	"fmt"
	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestQueryBlockHeader(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature(FeatureBlocks),
		allure.Story(StoryQueryBlock),
		allure.Description("Block header"),
		allure.Action(func() {
			// TODO: move to the before block
			client, err := connect()
			assert.Nil(t, err, fmt.Sprintf(FailedToConnectTo, GetConfig().ProxyURL, err))
			header, err := client.HeaderByNumber(context.Background(), nil)
			assert.Nil(t, err, "Failed to get block header")
			ReportErrorInAllure(err)
			assert.Greater(t, 0, header.Number, "Block header number greater than 0")
		}))
}
func TestQueryFullBlock(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature(FeatureBlocks),
		allure.Story(StoryQueryBlock),
		allure.Description("Full block"),
		allure.Action(func() {
			// TODO: move to the before block
			client, err := connect()
			assert.Nil(t, err, fmt.Sprintf(FailedToConnectTo, GetConfig().ProxyURL, err))
			blockNumber := big.NewInt(5671744)
			block, err := client.BlockByNumber(context.Background(), blockNumber)
			assert.Nil(t, err, "Failed to get full block")
			ReportErrorInAllure(err)
			assert.Greater(t, 0, block.Number().Uint64())
			// assert.Greater(t, 0, block.Time().Uint64())
			assert.Greater(t, 0, block.Difficulty().Uint64())
			assert.NotNil(t, block.Hash().Hex())
			assert.Greater(t, 0, len(block.Transactions()))
		}))
}
