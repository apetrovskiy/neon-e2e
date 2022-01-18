package go_eth

/*
import (
	// "context"
	"fmt"
	"math/big"

	log "github.com/sirupsen/logrus"

	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

var transferWithGasPriceTestData = []struct {
	senderBalance         int
	recipientBalance      int
	transferAmountInEther int
	transferAmountInWei   string
	gasLimit              uint64
	gasPrice              *big.Int
}{
	// {10, 10, 1, "1000000000000000000", 1000, big.NewInt(1000)},
	// {5, 0, 2, "2000000000000000000", 5000, big.NewInt(5000)},
	// {10, 3, 1, "1000000000000000000", 10000, big.NewInt(10000)},
	{5, 5, 1, "1000000000000000000", 20000, big.NewInt(20000)},
	{10, 10, 1, "1000000000000000000", 100000, big.NewInt(100000)},
}

func TestTransferEtherWithGasPrice(t *testing.T) {
  t.Parallel()
	for _, td := range transferWithGasPriceTestData {
		t.Run(fmt.Sprintf("Run, data = sender %d, recipient %d, transfer %d", td.senderBalance, td.recipientBalance, td.transferAmountInEther), func(t *testing.T) {
			allure.Test(t,
				allure.Epic(FeatureExternallyOwnedAccounts),
				allure.Lead(FeatureExternallyOwnedAccounts),
				// allure.Feature(StoryTransfer),
				allure.Story(StoryTransfer),
				allure.Description("Transfer Ether"),
				allure.Action(func() {

          client:=getClient()

					senderAccount := createWallet(td.senderBalance)
					allure.Step(allure.Description("Checking that sender's wallet is created"), allure.Action(func() {
						assert.NotEqual(t, 0, len(senderAccount.Address.Hash()), FailedToCreateWallet)
						if len(senderAccount.Address) == 0 {
							t.Error(FailedToCreateWallet)
						}
					}))

					senderBalance := getLastBlockBalance(client, senderAccount.Address.Hex())
					allure.Step(allure.Description("Checking sender's initial balance"), allure.Action(func() {
						log.Println(senderBalance)
						if td.senderBalance > 0 {
							assert.Equal(t, ToWei(td.senderBalance, 18), senderBalance, SenderInitialBalanceIsWrong)
						}
					}))

					recipientAccount := createWallet(td.recipientBalance)
					allure.Step(allure.Description("Checking that recipient's wallet is created"), allure.Action(func() {
						assert.NotEqual(t, 0, len(recipientAccount.Address.Hash()), FailedToCreateWallet)
						if len(recipientAccount.Address) == 0 {
							t.Error(FailedToCreateWallet)
						}
					}))

					recipientBalance := getLastBlockBalance(client, recipientAccount.Address.Hex())
					allure.Step(allure.Description("Checking recipient's initial balance"), allure.Action(func() {
						log.Println(recipientBalance)
						if td.recipientBalance > 0 {
							assert.Equal(t, ToWei(td.recipientBalance, 18), recipientBalance, RecipientInitialBalanceIsWrong)
						}
					}))

					// gasLimit := uint64(21000)
					// gasPrice, err := client.SuggestGasPrice(context.Background())
					// ReportErrorInAllure(err)

					transferEther(client, *senderAccount, *recipientAccount, td.transferAmountInWei) //, td.gasLimit, td.gasPrice)

					senderBalance = getLastBlockBalance(client, senderAccount.Address.Hex())
					allure.Step(allure.Description("Checking sender's resulting balance"), allure.Action(func() {
						log.Println(senderBalance)
						expectedSenderBalance := ToWei(td.senderBalance-td.transferAmountInEther, 18)
						assert.Equal(t, expectedSenderBalance, senderBalance, SenderResultingBalanceIsWrong)
					}))

					recipientBalance = getLastBlockBalance(client, recipientAccount.Address.Hex())
					allure.Step(allure.Description("Checking recipient's resulting balance"), allure.Action(func() {
						log.Println(recipientBalance)
						expectedRecipientBalance := ToWei(td.recipientBalance+td.transferAmountInEther, 18)
						assert.Equal(t, expectedRecipientBalance, recipientBalance, RecipientResultingBalanceIsWrong)
					}))
				}))
		})
	}
}
*/
