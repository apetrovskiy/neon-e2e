package NeonE2e

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func readEnvVariables() {
	_ = godotenv.Load("/home/runner/work/neon-e2e/neon-e2e/.env")
	_ = godotenv.Load("/home/runner/work/neon-e2e/neon-e2e/variables.env")
}
