package go_eth

import (
	"testing"

	"github.com/dailymotion/allure-go"
)

func TestConnection(t *testing.T) {
	allure.Test(t,
		allure.Description("Connection to network"),
		allure.Action(func() {
			_, err := connect()
			if err != nil {
				t.Errorf("Failed to connect to %s: %o", GetConfig().ProxyURL, err)
			}
		}))
}
