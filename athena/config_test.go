package athena

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("Config", func() {

	Describe("d", func() {
		Context("d", func() {
			//It("d", func() {
			//	config := Config{}
			//	Expect(config).Should(BeEquivalentTo(Config{}))
			//	config.Load()
			//	Expect(config).ShouldNot(BeEquivalentTo(Config{}))
			//	Expect(config).ShouldNot(BeNil())
			//
			//})
		})

	})

})