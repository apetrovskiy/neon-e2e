package go_eth

import (
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestQueryTransaction(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureTransaction),
		allure.Feature("Query transaction"),
		// allure.Story("Query transaction"),
		allure.Description("Query transaction"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestCreateRawTransaction(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureTransaction),
		allure.Feature("Create raw transaction"),
		// allure.Story("Create raw transaction"),
		allure.Description("Create raw transaction"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestSendRawTransaction(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureTransaction),
		allure.Feature("Send raw transaction"),
		// allure.Story("Send raw transaction"),
		allure.Description("Send raw transaction"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
