package go_eth

import (
	"fmt"
	"log"
	"math/big"

	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestTransferEther(t *testing.T) {

	allure.Test(t,
		allure.Epic(Epic),
		allure.Lead(FeatureExternallyOwnedAccounts),
		allure.Feature(FeatureExternallyOwnedAccounts),
		allure.Story(StoryTransfer),
		allure.Description("Transfer Ether"),
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

			transferEther(client, *senderAccount, *recipientAccount, "10000000000000000000")

			senderBalance = getLastBlockBalance(client, senderAccount.Address.Hex())
			log.Println(senderBalance)
			expectedSenderBalance, _ := new(big.Int).SetString("90000000000000000000", 0)
			assert.Equal(t, expectedSenderBalance, senderBalance, "Sender's initial balance is wrong")

			recipientBalance = getLastBlockBalance(client, recipientAccount.Address.Hex())
			log.Println(recipientBalance)
			expectedRecipientBalance, _ := new(big.Int).SetString("110000000000000000000", 0)
			assert.Equal(t, expectedRecipientBalance, recipientBalance, "Recipient's initial balance is wrong")
		}))
}
