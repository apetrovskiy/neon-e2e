package go_eth

import (
	"bytes"
	"encoding/json"
	"strconv"

	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/dailymotion/allure-go"
)

type FaucetRequest struct {
	Wallet string `json:"wallet"`
	Amount int    `json:"amount"`
}

func requestFaucet(wallet string, amount int) (statusCode int) {
	allure.Step(allure.Description("Requesting faucet "+GetConfig().FaucetUrl+" for "+strconv.Itoa(amount)+" to address "+wallet), allure.Action(func() {

		allure.Step(allure.Description("Checking useFaucet="+strconv.FormatBool(GetConfig().UseFaucet)+" and faucetUrl="+GetConfig().FaucetUrl), allure.Action(func() {
			if !GetConfig().UseFaucet || len(GetConfig().FaucetUrl) == 0 {
				statusCode = 404
				return
			}
		}))

		/*
		   reqBody, err := json.Marshal(map[string]string{
		     "amount": strconv.Itoa(amount),
		     "wallet": wallet,
		   })
		*/

		// faucetRequest := &FaucetRequest{Wallet: wallet, Amount: amount}
		/*
		   user := User{
		     Name: "Test User",
		     Job: "Go lang Developer",
		   }
		*/
		var faucetRequest FaucetRequest
		var reqBody []byte
		var err error
		var response *http.Response

		allure.Step(allure.Description("Preparing to POST request"), allure.Action(func() {
			faucetRequest = FaucetRequest{Wallet: wallet, Amount: amount}
			reqBody, err = json.Marshal(faucetRequest)
		}))
		allure.Step(allure.Description("Checking error"), allure.Action(func() {
			if err != nil {
				ReportErrorInAllure(err)
				return
			}
		}))

		allure.Step(allure.Description("Checking the request body"), allure.Action(func() {
			allure.Step(allure.Description("Checking the request body as string(reqBody)"), allure.Action(func() {
				log.Println("string(reqBody):")
				log.Println(string(reqBody))
			}))
			allure.Step(allure.Description("Checking the request body as bytes.NewBuffer(reqBody).String()"), allure.Action(func() {
				log.Println("bytes.NewBuffer(reqBody).String():")
				log.Println(bytes.NewBuffer(reqBody).String())
			}))
			allure.Step(allure.Description("Checking the request body as reqBody"), allure.Action(func() {
				log.Println("reqBody:")
				log.Println(reqBody)
			}))
		}))

		allure.Step(allure.Description("Sending POST request with body "+bytes.NewBuffer(reqBody).String()), allure.Action(func() {
			response, err = http.Post(GetConfig().FaucetUrl, "application/json", bytes.NewBuffer(reqBody))

			allure.Step(allure.Description("Checking error"), allure.Action(func() {
				ReportErrorInAllure(err)
			}))

			allure.Step(allure.Description("Checking status code "+strconv.Itoa(response.StatusCode)), allure.Action(func() {
				if response.StatusCode != 200 {
					allure.Step(allure.Description("Logging status code "+strconv.Itoa(response.StatusCode)), allure.Action(func() {
						log.Fatal(err)
					}))
				}
			}))
			statusCode = response.StatusCode
			return
		}))

	}))

	return
}
