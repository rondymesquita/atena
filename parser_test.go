package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestParser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parser Suite")
}

var _ = Describe("Parser behaviors", func() {

	Context("Should get the correct regexp from string template", func() {
		var parser *Parser
		BeforeEach(func(){
			parser = NewParser()
		})

		It("When string contains number", func() {
			regexp := parser.RegexpFrom("ATN_sample_<number>.doc")
			Expect(regexp).Should(Equal("ATN_sample_\\d*.doc"))
		})

		It("When string contains date", func() {
			regexp := parser.RegexpFrom("ATN_sample_<mm>_<dd>_<aaaa>.doc")
			Expect(regexp).Should(Equal("ATN_sample_\\d{2}_\\d{2}_\\d{4}.doc"))
		})

		It("When string contains number", func() {
			regexp := parser.RegexpFrom("ATN_sample_<*>.doc")
			Expect(regexp).Should(Equal("ATN_sample_.*.doc"))
		})
	})

})
