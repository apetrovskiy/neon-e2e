package go_eth

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestNewBlocksSubscription(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic("go-ethereum"),
		allure.Feature("Subscription"),
		allure.Story("Subscription to new blocks"),
		allure.Description("Subscription to new blocks"),
		allure.Action(func() {
			// TODO: logging
			fmt.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
