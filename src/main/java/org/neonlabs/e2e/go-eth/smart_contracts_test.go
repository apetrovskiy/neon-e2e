package go_eth

import (
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestDeploySmartContract(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureSmartContracts),
		allure.Feature("Deploy smart contracts"),
		// allure.Story("Deploy smart contracts"),
		allure.Description("Deploy smart contracts"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestLoadSmartContract(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureSmartContracts),
		allure.Feature("Load smart contracts"),
		// allure.Story("Load smart contracts"),
		allure.Description("Load smart contracts"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestQuerySmartConstract(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureSmartContracts),
		allure.Feature("Query smart contracts"),
		// allure.Story("Query smart contracts"),
		allure.Description("Query smart contracts"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestWriteToSmartConstract(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureSmartContracts),
		allure.Feature("Write to smart contracts"),
		// allure.Story("Write to smart contracts"),
		allure.Description("Write to smart contracts"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestReadSmartConstractBytecode(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureSmartContracts),
		allure.Feature("Read smart contract bytecode"),
		// allure.Story("Read smart contract bytecode"),
		allure.Description("Read smart contract bytecode"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestQueryERC20TokenSmartConstract(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureSmartContracts),
		allure.Feature("Query smart contracts"),
		// allure.Story("Query smart contracts"),
		allure.Description("Query ERC20 token smart contracts"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
