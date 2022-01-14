package go_eth

import (
	"fmt"
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	t.Parallel()
	allure.Test(t,
		allure.Epic(FeatureCommon),
		allure.Feature(StoryConnection),
		// allure.Story(StoryConnection),
		allure.Description("Connection to network"),
		allure.Action(func() {
			_, err := connect()
			assert.Nil(t, err, fmt.Sprintf(FailedToConnectTo, GetConfig().ProxyURL, err))
			if err != nil {
				t.Errorf(FailedToConnectTo, GetConfig().ProxyURL, err)
			}
		}))
}
