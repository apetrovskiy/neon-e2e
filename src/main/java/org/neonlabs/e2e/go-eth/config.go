package go_eth

import (
	"log"
	"math/big"
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
	SolanaExplorer, SolanaUrl                                   string
	UsersNumber                                                 int
	InitialBalance                                              *big.Int
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
	solanaExplorer := os.Getenv("SOLANA_EXPLORER")
	solanaUrl := os.Getenv("SOLANA_URL")
	usersNumber, _ := strconv.Atoi(os.Getenv("USERS_NUMBER"))
	bigIntHolder := new(big.Int)
	initialBalance, _ := bigIntHolder.SetString(os.Getenv("INITIAL_BALANCE"), 18)

	return &Config{
		NetworkName:    networkName,
		ProxyURL:       proxyUrl,
		NetworkId:      networkId,
		FaucetQuotient: faucetQuotient,
		FaucetUrl:      faucetUrl,
		AddressFrom:    addressFrom,
		AddressTo:      addressTo,
		PrivateKey:     privateKey,
		SolanaExplorer: solanaExplorer,
		SolanaUrl:      solanaUrl,
		UsersNumber:    usersNumber,
		InitialBalance: initialBalance,
	}
}
