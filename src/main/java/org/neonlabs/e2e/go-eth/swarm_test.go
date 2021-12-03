package go_eth

import (
	"fmt"
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestUploadToSwarm(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature(FeatureSwarm),
		allure.Story("Upload files to swarm"),
		allure.Description("Upload files to swarm"),
		allure.Action(func() {
			// TODO: logging
			fmt.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestDownloadFromSwarm(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature(FeatureSwarm),
		allure.Story("Download files from swarm"),
		allure.Description("Download files from swarm"),
		allure.Action(func() {
			// TODO: logging
			fmt.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
