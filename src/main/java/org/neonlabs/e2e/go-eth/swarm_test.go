package go_eth

import (
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestUploadToSwarm(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureSwarm),
		allure.Feature("Upload files to swarm"),
		// allure.Story("Upload files to swarm"),
		allure.Description("Upload files to swarm"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestDownloadFromSwarm(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureSwarm),
		allure.Feature("Download files from swarm"),
		// allure.Story("Download files from swarm"),
		allure.Description("Download files from swarm"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
