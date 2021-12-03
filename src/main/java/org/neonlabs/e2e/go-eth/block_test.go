package go_eth

import (
	"fmt"
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestQueryBlockHeader(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature("Block"),
		allure.Story("Query block"),
		allure.Description("Block header"),
		allure.Action(func() {
			// TODO: logging
			fmt.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestQueryFullBlock(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature("Block"),
		allure.Story("Query block"),
		allure.Description("Full block"),
		allure.Action(func() {
			// TODO: logging
			fmt.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
