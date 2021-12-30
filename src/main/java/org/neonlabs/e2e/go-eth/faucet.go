package go_eth

import (
	"bytes"
	"encoding/json"
	"strconv"

	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/dailymotion/allure-go"
)

func requestFaucet(amount int, wallet string) {
	allure.Step(allure.Description("Requesting faucet "+GetConfig().FaucetUrl+" for address "+wallet), allure.Action(func() {
		reqBody, err := json.Marshal(map[string]string{
			"amount": strconv.Itoa(amount),
			"wallet": wallet,
		})
		if err != nil {
			log.Println(err)
		}

		http.Post(GetConfig().FaucetUrl, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			log.Println(err)
		}
	}))
}
