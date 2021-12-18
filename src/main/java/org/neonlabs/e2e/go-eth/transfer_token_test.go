package go_eth

import (
	"fmt"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

// TODO: finish it
func TestTransferToken(t *testing.T) {

	allure.SkipTest(t,
		// allure.Suite(Epic),
		allure.Epic(Epic),
		allure.Lead(FeatureExternallyOwnedAccounts),
		allure.Feature(FeatureExternallyOwnedAccounts),
		allure.Story(StoryTransfer),
		allure.Description("Transfer tokens"),
		allure.Action(func() {
			client, err := connect()
			assert.Nil(t, err, fmt.Sprintf(FaileToConnectTo, GetConfig().ProxyURL, err))
			if err != nil {
				t.Errorf(FaileToConnectTo, GetConfig().ProxyURL, err)
			}

			senderAccount := createWallet()
			assert.NotEqual(t, 0, len(senderAccount.Address.Hash()), FailedToCreateWallet)
			if len(senderAccount.Address) == 0 {
				t.Error(FailedToCreateWallet)
			}

			senderBalance := getLastBlockBalance(client, senderAccount.Address.Hex())
			log.Println(senderBalance)
			assert.Equal(t, GetConfig().InitialBalance, senderBalance, "Sender's initial balance is wrong")

			recipientAccount := createWallet()
			assert.NotEqual(t, 0, len(recipientAccount.Address.Hash()), FailedToCreateWallet)
			if len(recipientAccount.Address) == 0 {
				t.Error(FailedToCreateWallet)
			}

			recipientBalance := getLastBlockBalance(client, recipientAccount.Address.Hex())
			log.Println(recipientBalance)
			assert.Equal(t, GetConfig().InitialBalance, recipientBalance, "Recipient's initial balance is wrong")

			transferToken(client, *senderAccount, *recipientAccount, "1000000000000000000")

			// TODO: change to the right amounts
			senderBalance = getLastBlockBalance(client, senderAccount.Address.Hex())
			log.Println(senderBalance)
			assert.Equal(t, GetConfig().InitialBalance, senderBalance, "Sender's initial balance is wrong")

			recipientBalance = getLastBlockBalance(client, recipientAccount.Address.Hex())
			log.Println(recipientBalance)
			assert.Equal(t, GetConfig().InitialBalance, recipientBalance, "Recipient's initial balance is wrong")
		}))
}
