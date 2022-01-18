package go_eth

import (
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

// TODO: finish it
func TestTransferToken(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		// allure.Suite(Epic),
		allure.Epic(FeatureExternallyOwnedAccounts),
		allure.Lead(FeatureExternallyOwnedAccounts),
		allure.Feature(StoryTransfer),
		// allure.Story(StoryTransfer),
		allure.Description("Transfer tokens"),
		allure.Action(func() {

			client := getClient()

			initialAmount := 10
			transferAmount := "10000000000000000000"
			senderAccount := createWallet(initialAmount)
			assert.NotEqual(t, 0, len(senderAccount.Address.Hash()), FailedToCreateWallet)
			if len(senderAccount.Address) == 0 {
				t.Error(FailedToCreateWallet)
			}

			senderBalance := getLastBlockBalance(client, senderAccount.Address.Hex())
			log.Println(senderBalance)
			assert.Equal(t, GetConfig().RequestAmount, senderBalance, SenderInitialBalanceIsWrong)

			recipientAccount := createWallet(initialAmount)
			assert.NotEqual(t, 0, len(recipientAccount.Address.Hash()), FailedToCreateWallet)
			if len(recipientAccount.Address) == 0 {
				t.Error(FailedToCreateWallet)
			}

			recipientBalance := getLastBlockBalance(client, recipientAccount.Address.Hex())
			log.Println(recipientBalance)
			assert.Equal(t, GetConfig().RequestAmount, recipientBalance, RecipientInitialBalanceIsWrong)

			transferToken(client, *senderAccount, *recipientAccount, transferAmount)

			// TODO: change to the right amounts
			senderBalance = getLastBlockBalance(client, senderAccount.Address.Hex())
			log.Println(senderBalance)
			assert.Equal(t, GetConfig().RequestAmount, senderBalance, SenderResultingBalanceIsWrong)

			recipientBalance = getLastBlockBalance(client, recipientAccount.Address.Hex())
			log.Println(recipientBalance)
			assert.Equal(t, GetConfig().RequestAmount, recipientBalance, RecipientResultingBalanceIsWrong)
		}))
}
