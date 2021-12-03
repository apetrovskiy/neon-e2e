package go_eth

import (
	"fmt"
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {

	allure.Test(t,
		allure.Epic(Epic),
		allure.Feature(FeatureCommon),
		allure.Story(StoryConnection),
		allure.Description("Connection to network"),
		allure.Action(func() {
			_, err := connect()
			assert.Nil(t, err, fmt.Sprintf(FaileToConnectTo, GetConfig().ProxyURL, err))
			if err != nil {
				t.Errorf(FaileToConnectTo, GetConfig().ProxyURL, err)
			}
		}))
}
