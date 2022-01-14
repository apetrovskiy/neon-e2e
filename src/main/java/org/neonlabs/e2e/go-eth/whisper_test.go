package go_eth

import (
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestWhisperConnection(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureWhisper),
		allure.Feature("Connect Whisper client"),
		// allure.Story("Connect Whisper client"),
		allure.Description("Connect Whisper client"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestGenerateWhisperKeyPair(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureWhisper),
		allure.Feature("Generate Whisper key pair"),
		// allure.Story("Generate Whisper key pair"),
		allure.Description("Generate Whisper key pair"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestSendMessageOnWhisper(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureWhisper),
		allure.Feature("Send message on Whisper"),
		// allure.Story("Send message on Whisper"),
		allure.Description("Send message on Whisper"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestSubscribeToWhisperMessages(t *testing.T) {
	t.Parallel()
	allure.SkipTest(t,
		allure.Epic(FeatureWhisper),
		allure.Feature("Subscribe to Whisper messages"),
		// allure.Story("Subscribe to Whisper messages"),
		allure.Description("Subscribe to Whisper messages"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
