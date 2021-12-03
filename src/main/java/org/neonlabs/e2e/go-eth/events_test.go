package go_eth

import (
	"fmt"
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestNewBlocksSubscription(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature(FeatureEvents),
		allure.Story("Subscription to new blocks"),
		allure.Description("Subscription to new blocks"),
		allure.Action(func() {
			// TODO: logging
			fmt.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestEventLogSubscription(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature(FeatureEvents),
		allure.Story("Subscription to event logs"),
		allure.Description("Subscription to event logs"),
		allure.Action(func() {
			// TODO: logging
			fmt.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestReadEventLogs(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature(FeatureEvents),
		allure.Story("Read event logs"),
		allure.Description("Read event logs"),
		allure.Action(func() {
			// TODO: logging
			fmt.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestReadERC20TokenEventLogs(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature(FeatureEvents),
		allure.Story("Read ERC20 token event logs"),
		allure.Description("Read ERC20 token event logs"),
		allure.Action(func() {
			// TODO: logging
			fmt.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestRead0xProtocolEventLogs(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature(FeatureEvents),
		allure.Story("Read 0x protocol event logs"),
		allure.Description("Read 0x protocol event logs"),
		allure.Action(func() {
			// TODO: logging
			fmt.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
