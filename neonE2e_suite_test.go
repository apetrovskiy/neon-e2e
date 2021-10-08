package neonE2e_test

import (
	// "fmt"
	// "os"
	"testing"

	// "github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNeonE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NeonE2e Suite")
	// _ = godotenv.Load("/home/runner/work/neon-e2e/neon-e2e/.env")
	// _ = godotenv.Load("/home/runner/work/neon-e2e/neon-e2e/variables.env")
	// fmt.Println("=========111=========")
	// fmt.Println(os.Getenv("ALLURE_RESULTS_PATH"))
}
