package go_eth

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

var transferTestData = []struct {
	senderBalance         int
	recipientBalance      int
	transferAmountInEther int
	transferAmountInWei   string
}{
	{10, 10, 1, "1000000000000000000"},
	{5, 0, 2, "2000000000000000000"},
	{10, 3, 1, "1000000000000000000"},
	{5, 5, 1, "1000000000000000000"},
	{10, 10, 1, "1000000000000000000"},
	{10, 10, 1, "1000000000000000000"},
	{5, 5, 1, "1000000000000000000"},
	{5, 5, 1, "1000000000000000000"},
	{5, 5, 1, "1000000000000000000"},
	{5, 5, 1, "1000000000000000000"},
}

var bidirectionalTransferTestData = []struct {
	senderBalance                 int
	recipientBalance              int
	directTransferAmountInEther   int
	directTransferAmountInWei     string
	backwardTransferAmountInEther int
	backwardTransferAmountInWei   string
}{
	{10, 10, 1, "1000000000000000000", 2, "2000000000000000000"},
	{5, 0, 2, "2000000000000000000", 1, "1000000000000000000"},
	{10, 3, 1, "1000000000000000000", 2, "2000000000000000000"},
	{5, 5, 1, "1000000000000000000", 2, "2000000000000000000"},
	{10, 10, 1, "1000000000000000000", 2, "2000000000000000000"},
	{10, 10, 1, "1000000000000000000", 2, "2000000000000000000"},
	{5, 5, 1, "1000000000000000000", 2, "2000000000000000000"},
	{5, 5, 1, "1000000000000000000", 2, "2000000000000000000"},
	{5, 5, 1, "1000000000000000000", 2, "2000000000000000000"},
	{5, 5, 1, "1000000000000000000", 2, "2000000000000000000"},
}

func TestTransferEther(t *testing.T) {
	for _, td := range transferTestData {
		t.Run(fmt.Sprintf("Run, data = sender %d, recipient %d, transfer %d", td.senderBalance, td.recipientBalance, td.transferAmountInEther), func(t *testing.T) {
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

					transferEther(client, *senderAccount, *recipientAccount, td.transferAmountInWei)

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

func TestTwoTransfersOfEther(t *testing.T) {
	for _, td := range transferTestData {
		t.Run(fmt.Sprintf("Run, data = sender %d, recipient %d, transfers %d", td.senderBalance, td.recipientBalance, td.transferAmountInEther), func(t *testing.T) {
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

					transferEther(client, *senderAccount, *recipientAccount, td.transferAmountInWei)
					transferEther(client, *senderAccount, *recipientAccount, td.transferAmountInWei)

					senderBalance = getLastBlockBalance(client, senderAccount.Address.Hex())
					allure.Step(allure.Description("Checking sender's resulting balance"), allure.Action(func() {
						log.Println(senderBalance)
						expectedSenderBalance := ToWei(td.senderBalance-td.transferAmountInEther*2, 18)
						assert.Equal(t, expectedSenderBalance, senderBalance, SenderResultingBalanceIsWrong)
					}))

					recipientBalance = getLastBlockBalance(client, recipientAccount.Address.Hex())
					allure.Step(allure.Description("Checking recipient's resulting balance"), allure.Action(func() {
						log.Println(recipientBalance)
						expectedRecipientBalance := ToWei(td.recipientBalance+td.transferAmountInEther*2, 18)
						assert.Equal(t, expectedRecipientBalance, recipientBalance, RecipientResultingBalanceIsWrong)
					}))
				}))
		})
	}
}

func TestBidirectionalTransfersOfEther(t *testing.T) {
	for _, td := range bidirectionalTransferTestData {
		t.Run(fmt.Sprintf("Run, data = sender %d, recipient %d, transfer s->r %d, transfer r->s %d", td.senderBalance, td.recipientBalance, td.directTransferAmountInEther, td.backwardTransferAmountInEther), func(t *testing.T) {
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

					transferEther(client, *senderAccount, *recipientAccount, td.directTransferAmountInWei)
					transferEther(client, *recipientAccount, *senderAccount, td.backwardTransferAmountInWei)

					senderBalance = getLastBlockBalance(client, senderAccount.Address.Hex())
					allure.Step(allure.Description("Checking sender's resulting balance"), allure.Action(func() {
						log.Println(senderBalance)
						expectedSenderBalance := ToWei(td.senderBalance-td.directTransferAmountInEther+td.backwardTransferAmountInEther, 18)
						assert.Equal(t, expectedSenderBalance, senderBalance, SenderResultingBalanceIsWrong)
					}))

					recipientBalance = getLastBlockBalance(client, recipientAccount.Address.Hex())
					allure.Step(allure.Description("Checking recipient's resulting balance"), allure.Action(func() {
						log.Println(recipientBalance)
						expectedRecipientBalance := ToWei(td.recipientBalance+td.directTransferAmountInEther-td.backwardTransferAmountInEther, 18)
						assert.Equal(t, expectedRecipientBalance, recipientBalance, RecipientResultingBalanceIsWrong)
					}))
				}))
		})
	}
}
