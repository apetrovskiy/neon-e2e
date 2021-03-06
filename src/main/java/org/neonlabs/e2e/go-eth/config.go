package go_eth

import (
	"os"
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/dailymotion/allure-go"
	"github.com/joho/godotenv"
)

const projectDirName = "neon-e2e"

type Config struct {
	NetworkName, ProxyURL, NetworkId, CurrencySymbol, FaucetUrl string
	RequestAmount                                               int
	UseFaucet                                                   bool
	AddressFrom, AddressTo, PrivateKey                          string
	SolanaExplorer, SolanaUrl                                   string
	UsersNumber                                                 int
}

func loadEnv() {
	log.SetFormatter(&log.TextFormatter{})

	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	ReportErrorInAllure(err)
}

func GetConfig() *Config {
	var config Config

	allure.Step(allure.Description("Loading the .env file"), allure.Action(func() {
		loadEnv()

		networkName := os.Getenv("NETWORK_NAME")
		proxyUrl := os.Getenv("PROXY_URL")
		networkId := os.Getenv("NETWORK_ID")
		requestAmount, _ := strconv.Atoi(os.Getenv("REQUEST_AMOUNT"))
		faucetUrl := os.Getenv("FAUCET_URL")
		useFaucet, _ := strconv.ParseBool(os.Getenv("USE_FAUCET"))
		addressFrom := os.Getenv("ADDRESS_FROM")
		addressTo := os.Getenv("ADDRESS_TO")
		privateKey := os.Getenv("PRIVATE_KEY")
		solanaExplorer := os.Getenv("SOLANA_EXPLORER")
		solanaUrl := os.Getenv("SOLANA_URL")
		usersNumber, _ := strconv.Atoi(os.Getenv("USERS_NUMBER"))

		config = Config{
			NetworkName:    networkName,
			ProxyURL:       proxyUrl,
			NetworkId:      networkId,
			RequestAmount:  requestAmount,
			FaucetUrl:      faucetUrl,
			UseFaucet:      useFaucet,
			AddressFrom:    addressFrom,
			AddressTo:      addressTo,
			PrivateKey:     privateKey,
			SolanaExplorer: solanaExplorer,
			SolanaUrl:      solanaUrl,
			UsersNumber:    usersNumber,
		}
	}))
	return &config
}
