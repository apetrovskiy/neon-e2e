package go_eth

import (
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestNewWallet(t *testing.T) {
	t.Parallel()
	allure.Test(t,
		allure.Epic(FeatureExternallyOwnedAccounts),
		allure.Feature(StoryWallet),
		// allure.Story(StoryWallet),
		allure.Description("Creating a new wallet"),
		allure.Action(func() {
			account := createWallet(GetConfig().RequestAmount)
			assert.NotEqual(t, 0, len(account.Address.Hash()), FailedToCreateWallet)
			if len(account.Address.Hash()) == 0 {
				t.Error(FailedToCreateWallet)
				assert.Fail(t, "Wallet hasn't been created")
			}
			t.Log("Wallet has been created")

		}))
}
