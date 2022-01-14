package go_eth

import (
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestAddressIsValid(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureAddress),
		allure.Feature(StoryAddressValidation),
		// allure.Story(StoryAddressValidation),
		allure.Description("Address is valid"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestAddressIsAccount(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureAddress),
		allure.Feature(StoryAddressValidation),
		// allure.Story(StoryAddressValidation),
		allure.Description("Address is account"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestAddressIsSmartConstract(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureAddress),
		allure.Feature(StoryAddressValidation),
		// allure.Story(StoryAddressValidation),
		allure.Description("Address is smart contract"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
