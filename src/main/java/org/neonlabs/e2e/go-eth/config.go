package go_eth

import (
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
)

const projectDirName = "neon-e2e"

type Config struct {
	NetworkName, ProxyURL, NetworkId, CurrencySymbol, FaucetUrl string
	FaucetQuotient                                              int
	AddressFrom, AddressTo, PrivateKey                          string
}

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetConfig() *Config {
	loadEnv()

	networkName := os.Getenv("NETWORK_NAME")
	proxyUrl := os.Getenv("PROXY_URL")
	networkId := os.Getenv("NETWORK_ID")
	faucetQuotient, _ := strconv.Atoi(os.Getenv("FAUCET_QUOTIENT"))
	faucetUrl := os.Getenv("FAUCET_URL")
	addressFrom := os.Getenv("ADDRESS_FROM")
	addressTo := os.Getenv("ADDRESS_TO")
	privateKey := os.Getenv("PRIVATE_KEY")

	return &Config{
		NetworkName:    networkName,
		ProxyURL:       proxyUrl,
		NetworkId:      networkId,
		FaucetQuotient: faucetQuotient,
		FaucetUrl:      faucetUrl,
		AddressFrom:    addressFrom,
		AddressTo:      addressTo,
		PrivateKey:     privateKey,
	}
}
