package neonE2e

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func readEnvVariables() {
	_ = godotenv.Load(".env")
	_ = godotenv.Load("variables.env")
	fmt.Println("=========000=========")
	fmt.Println(os.Getenv("ALLURE_RESULTS_PATH"))
}
