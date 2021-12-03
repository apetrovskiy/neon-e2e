package go_eth

import (
	"log"
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestWhisperConnection(t *testing.T) {

	allure.SkipTest(t,
		allure.Epic(Epic),
		allure.Feature(FeatureWhister),
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
		allure.Feature(FeatureWhister),
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
		allure.Feature(FeatureWhister),
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
		allure.Feature(FeatureWhister),
		allure.Story("Subscribe to Whisper messages"),
		allure.Description("Subscribe to Whisper messages"),
		allure.Action(func() {
			log.Println("111")
			assert.Equal(t, 1, 1, "111")
		}))
}
