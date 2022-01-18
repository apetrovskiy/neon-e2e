package go_eth

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

const (
	initialBalance      = 10
	attemptsToFindBlock = 1000
)

func TestQueryTransactionForCommonData(t *testing.T) {
	t.Parallel()
	allure.Test(t,
		allure.Epic(FeatureTransaction),
		allure.Feature("Query transaction"),
		// allure.Story("Query transaction"),
		allure.Description("Query transaction for common data"),
		allure.Action(func() {

			client := getClient()

			var block *types.Block
			for i := 0; i < attemptsToFindBlock; i++ {
				var result *types.Block
				allure.Step(allure.Description("Getting the highest block by number"), allure.Action(func() {
					defer func() *types.Block {
						if p := recover(); p != nil {
							block = GetBlockByNumber(client, big.NewInt(2))
						}
						if nil == block {
							return result
						}
						allure.Step(allure.Description("Checking block's transactions"), allure.Action(func() {
							for _, tx := range block.Transactions() {
								if nil == tx {
									continue
								}
								assert.NotNil(t, tx)
								allure.Step(allure.Description("Transaction is not nil"), allure.Action(func() {
									// assert.Greater(t, 0, len(tx.Hash().Hex()), "Transaction hash is of zero length")
									assert.Greater(t, 0, len(tx.Value().String()), "Transaction value is of zero length")
									assert.Greater(t, 0, tx.Gas(), "Transaction gas equals zero")
									assert.Greater(t, 0, tx.GasPrice().Uint64(), "Transaction gas price equals zero")
									// assert.Greater(t, 0, tx.Nonce(), "Transaction nonce equals zero")
									// assert.Greater(t, 0, len(tx.Data()), "Transaction data is of zero length")
									assert.Greater(t, 0, len(tx.To().Hex()), "Transaction To address is of zero length")
								}))
							}
						}))
						return result
					}()
				}))
				if nil == result {
					continue
				} else {
					break
				}
			}

		}))
}
func TestQueryTransactionForSenderAddress(t *testing.T) {
	t.Parallel()
	allure.Test(t,
		allure.Epic(FeatureTransaction),
		allure.Feature("Query transaction"),
		// allure.Story("Query transaction"),
		allure.Description("Query transaction for common data"),
		allure.Action(func() {

			client := getClient()

			chainId := getChainId(client)
			var block *types.Block
			for i := 0; i < attemptsToFindBlock; i++ {
				var result *types.Block
				allure.Step(allure.Description("Getting the highest block by number"), allure.Action(func() {
					defer func() *types.Block {
						if p := recover(); p != nil {
							block = GetBlockByNumber(client, big.NewInt(2))
						}
						if nil == block {
							return result
						}
						allure.Step(allure.Description("Checking block's transactions"), allure.Action(func() {
							for _, tx := range block.Transactions() {
								if msg, err := tx.AsMessage(types.NewEIP155Signer(chainId), big.NewInt(0)); err != nil {
									assert.Greater(t, 0, len(msg.From().Hex()), "Transaction From address is of zero lenght")
								}
							}
						}))
						return result
					}()
				}))
				if nil == result {
					continue
				} else {
					break
				}
			}
		}))
}
func TestCreateRawTransaction(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureTransaction),
		allure.Feature("Create raw transaction"),
		// allure.Story("Create raw transaction"),
		allure.Description("Create raw transaction"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestSendRawTransaction(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureTransaction),
		allure.Feature("Send raw transaction"),
		// allure.Story("Send raw transaction"),
		allure.Description("Send raw transaction"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
