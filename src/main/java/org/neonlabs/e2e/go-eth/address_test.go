package go_eth

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestAddressIsValid(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic("go-ethereum"),
		allure.Feature("Address"),
		allure.Story("Address validation"),
		allure.Description("Get the latest block balance"),
		allure.Action(func() {
			// TODO: logging
			fmt.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestAddressIsAccount(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic("go-ethereum"),
		allure.Feature("Address"),
		allure.Story("Address validation"),
		allure.Description("Get the latest block balance"),
		allure.Action(func() {
			// TODO: logging
			fmt.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestAddressIsSmartConstract(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic("go-ethereum"),
		allure.Feature("Address"),
		allure.Story("Address validation"),
		allure.Description("Get the latest block balance"),
		allure.Action(func() {
			// TODO: logging
			fmt.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
