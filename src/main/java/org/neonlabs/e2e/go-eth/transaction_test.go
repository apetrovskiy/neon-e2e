package go_eth

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestQueryTransaction(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic("go-ethereum"),
		allure.Feature("Transaction"),
		allure.Story("Query transaction"),
		allure.Description("Query transaction"),
		allure.Action(func() {
			// TODO: logging
			fmt.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestCreateRawTransaction(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic("go-ethereum"),
		allure.Feature("Transaction"),
		allure.Story("Create raw transaction"),
		allure.Description("Create raw transaction"),
		allure.Action(func() {
			// TODO: logging
			fmt.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestSendRawTransaction(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic("go-ethereum"),
		allure.Feature("Transaction"),
		allure.Story("Send raw transaction"),
		allure.Description("Send raw transaction"),
		allure.Action(func() {
			// TODO: logging
			fmt.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
