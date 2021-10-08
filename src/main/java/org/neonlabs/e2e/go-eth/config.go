package go_eth

const projectDirName = "neon-e2e"

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
	currencySymbol := os.Getenv("CURRENCY_SYMBOL")
	faucetQuotient := os.Getenv("FAUCET_QUOTIENT")
	faucetUrl := os.Getenv("FAUCET_URL")

	return &Config{
		DB: &DBConfig{
			NetworkName:    networkName,
			ProxyURL:       proxyUrl,
			NetworkId:      networkId,
			CurrencySymbol: currencySymbol,
			FaucetQuotient: faucetQuotient,
			FaucetUrl:      faucetUrl,
		},
	}
}
