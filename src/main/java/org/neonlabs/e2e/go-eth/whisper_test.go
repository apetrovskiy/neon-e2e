package go_eth

import (
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestWhisperConnection(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature(FeatureWhisper),
		allure.Story("Connect Whisper client"),
		allure.Description("Connect Whisper client"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestGenerateWhisperKeyPair(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature(FeatureWhisper),
		allure.Story("Generate Whisper key pair"),
		allure.Description("Generate Whisper key pair"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestSendMessageOnWhisper(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature(FeatureWhisper),
		allure.Story("Send message on Whisper"),
		allure.Description("Send message on Whisper"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
func TestSubscribeToWhisperMessages(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature(FeatureWhisper),
		allure.Story("Subscribe to Whisper messages"),
		allure.Description("Subscribe to Whisper messages"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
