package team

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ZupIT/ritchie-cli/functional"
)

func TestRitScaffold(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rit Suite")
}

var _ = Describe("RitScaffold", func() {
	BeforeSuite(func() {
		functional.FuncValidateLoginRequired()
	})
})
