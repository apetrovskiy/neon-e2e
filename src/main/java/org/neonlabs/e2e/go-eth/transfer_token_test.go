package go_eth

import (
	"fmt"
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

// TODO: finish it
func TestTransferToken(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Lead(FeatureExternallyOwnedAccounts),
		allure.Feature(FeatureExternallyOwnedAccounts),
		allure.Story(StoryTransfer),
		allure.Description("Transfer tokens"),
		allure.Action(func() {
			client, err := connect()
			if err != nil {
				t.Errorf("Failed to connect to %s: %o", GetConfig().ProxyURL, err)
			}

			senderAccount := createWallet()
			if len(senderAccount.Address) == 0 {
				t.Error("Failed to create a new wallet")
			}

			senderBalance := getLastBlockBalance(client, senderAccount.Address.Hex())
			// TODO: logging
			fmt.Println(senderBalance)
			assert.Equal(t, GetConfig().InitialBalance, senderBalance, "Sender's initial balance is wrong")

			recipientAccount := createWallet()
			if len(recipientAccount.Address) == 0 {
				t.Error("Failed to create a new wallet")
			}

			recipientBalance := getLastBlockBalance(client, recipientAccount.Address.Hex())
			// TODO: logging
			fmt.Println(recipientBalance)
			assert.Equal(t, GetConfig().InitialBalance, recipientBalance, "Recipient's initial balance is wrong")

			transferToken(client, *senderAccount, *recipientAccount, "1000000000000000000")

			// TODO: change to the right amounts
			senderBalance = getLastBlockBalance(client, senderAccount.Address.Hex())
			// TODO: logging
			fmt.Println(senderBalance)
			assert.Equal(t, GetConfig().InitialBalance, senderBalance, "Sender's initial balance is wrong")

			recipientBalance = getLastBlockBalance(client, recipientAccount.Address.Hex())
			// TODO: logging
			fmt.Println(recipientBalance)
			assert.Equal(t, GetConfig().InitialBalance, recipientBalance, "Recipient's initial balance is wrong")
		}))
}
