package go_eth

import (
	log "github.com/sirupsen/logrus"

	"github.com/dailymotion/allure-go"
)

func ReportErrorInAllure(err error) {
	if err != nil {
		allure.Step(allure.Description("Loging error "+err.Error()), allure.Action(func() {
			log.Fatal(err)
		}))
	}
}
