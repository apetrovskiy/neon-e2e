package go_eth

import (
	"fmt"

	log "github.com/sirupsen/logrus"

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
			assert.Nil(t, err, fmt.Sprintf(FailedToConnectTo, GetConfig().ProxyURL, err))
			if err != nil {
				t.Errorf(FailedToConnectTo, GetConfig().ProxyURL, err)
			}

			initialAmount := 10
			transferAmount := "1000000000000000000"
			senderAccount := createWallet(initialAmount)
			allure.Step(allure.Description("Checking that sender's wallet is created"), allure.Action(func() {
				assert.NotEqual(t, 0, len(senderAccount.Address.Hash()), FailedToCreateWallet)
				if len(senderAccount.Address) == 0 {
					t.Error(FailedToCreateWallet)
				}
			}))

			senderBalance := getLastBlockBalance(client, senderAccount.Address.Hex())
			allure.Step(allure.Description("Checking sender's initial balance"), allure.Action(func() {
				log.Println(senderBalance)
				assert.Equal(t, ToWei(initialAmount, 18), senderBalance, SenderInitialBalanceIsWrong)
			}))

			recipientAccount := createWallet(initialAmount)
			allure.Step(allure.Description("Checking that recipient's wallet is created"), allure.Action(func() {
				assert.NotEqual(t, 0, len(recipientAccount.Address.Hash()), FailedToCreateWallet)
				if len(recipientAccount.Address) == 0 {
					t.Error(FailedToCreateWallet)
				}
			}))

			recipientBalance := getLastBlockBalance(client, recipientAccount.Address.Hex())
			allure.Step(allure.Description("Checking recipient's initial balance"), allure.Action(func() {
				log.Println(recipientBalance)
				assert.Equal(t, ToWei(initialAmount, 18), recipientBalance, RecipientInitialBalanceIsWrong)
			}))

			transferEther(client, *senderAccount, *recipientAccount, transferAmount)

			senderBalance = getLastBlockBalance(client, senderAccount.Address.Hex())
			allure.Step(allure.Description("Checking sender's resulting balance"), allure.Action(func() {
				log.Println(senderBalance)
				expectedSenderBalance := ToWei(initialAmount-1, 18)
				assert.Equal(t, expectedSenderBalance, senderBalance, SenderResultingBalanceIsWrong)
			}))

			recipientBalance = getLastBlockBalance(client, recipientAccount.Address.Hex())
			allure.Step(allure.Description("Checking recipient's resulting balance"), allure.Action(func() {
				log.Println(recipientBalance)
				expectedRecipientBalance := ToWei(initialAmount+1, 18)
				assert.Equal(t, expectedRecipientBalance, recipientBalance, RecipientResultingBalanceIsWrong)
			}))
		}))
}
