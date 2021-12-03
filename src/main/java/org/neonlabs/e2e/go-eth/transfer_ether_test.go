package go_eth

import (
	"fmt"
	"math/big"

	// TODO: clean it up
	// "math/big"
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestTransferEther(t *testing.T) {

	allure.Test(t,
		allure.Epic(Epic),
		allure.Lead(FeatureExternallyOwnedAccounts),
		allure.Feature(Epic),
		allure.Story("Transfer"),
		allure.Description("Transfer Ether"),
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

			transferEther(client, *senderAccount, *recipientAccount, "10000000000000000000")

			// TODO: change to the right amounts
			senderBalance = getLastBlockBalance(client, senderAccount.Address.Hex())
			// TODO: logging
			fmt.Println(senderBalance)
			expectedSenderBalance, _ := new(big.Int).SetString("90000000000000000000", 0)
			assert.Equal(t, expectedSenderBalance, senderBalance, "Sender's initial balance is wrong")

			recipientBalance = getLastBlockBalance(client, recipientAccount.Address.Hex())
			// TODO: logging
			fmt.Println(recipientBalance)
			expectedRecipientBalance, _ := new(big.Int).SetString("110000000000000000000", 0)
			assert.Equal(t, expectedRecipientBalance, recipientBalance, "Recipient's initial balance is wrong")
		}))
}
