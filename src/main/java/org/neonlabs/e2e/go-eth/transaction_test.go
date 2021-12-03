package go_eth

import (
	"log"
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestQueryTransaction(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature("Transaction"),
		allure.Story("Query transaction"),
		allure.Description("Query transaction"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestCreateRawTransaction(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature("Transaction"),
		allure.Story("Create raw transaction"),
		allure.Description("Create raw transaction"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestSendRawTransaction(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature("Transaction"),
		allure.Story("Send raw transaction"),
		allure.Description("Send raw transaction"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
