package neonE2e_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNeonE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NeonE2e Suite")
}
