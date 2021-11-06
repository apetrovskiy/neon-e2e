package go_eth

import (
	"fmt"
	"testing"

	"github.com/dailymotion/allure-go"
)

var testData = []struct {
	x   int
	y   int
	d   int
	out int
}{
	{10, 85, 30, 3},
	{10, 234_342_345, 8, 29_292_792},
	{10, 234_342_345, 1, 234_342_335},
	{10, 2_234_342_345, 1, 2_234_342_335},
	{10, 2_234_342_345, 1_000_000_000, 3},
	{100, 100, 2, 0},
}

func TestFrogJmp(t *testing.T) {
	allure.Epic("go temporary tests")
	allure.Story("go temporary tests")
	for _, td := range testData {
		t.Run(fmt.Sprintf("Run, data = %o", td), func(t *testing.T) {
			allure.Test(t, allure.Action(func() {
				if td.out != SolutionFrogJmp(td.x, td.y, td.d) {

					fmt.Println(GetConfig().NetworkId)
					fmt.Println(GetConfig().ProxyURL)
					fmt.Println(GetConfig().FaucetUrl)

					t.Errorf("FAIL!")
				}
			}))
		})
	}
}

func TestPassedExample(t *testing.T) {

	allure.Epic("go temporary tests")
	allure.Story("go temporary tests")
	allure.Test(t,
		allure.Description("This is a test to show allure implementation with a passing test"),
		allure.Action(func() {
			s := "Hello world"
			if len(s) == 0 {
				t.Errorf("Expected 'hello world' string, but got %s ", s)
			}
		}))
}
