package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func readEnvVariables() {
	_ = godotenv.Load("/home/runner/work/neon-e2e/neon-e2e/.env")
	_ = godotenv.Load("/home/runner/work/neon-e2e/neon-e2e/variables.env")
	fmt.Println("=========000=========")
	fmt.Println(os.Getenv("ALLURE_RESULTS_PATH"))
}

func main() {
  readEnvVariables()
}
