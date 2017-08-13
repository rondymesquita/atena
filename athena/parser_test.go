package athena

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

	Context("Should verify if value has match with patterns", func() {
		var parser *Parser
		BeforeEach(func() {
			parser = NewParser()
		})

		Context("When string contains number", func() {
			It("And appears only once", func() {
				matched, err := parser.HasMatch("ATN_sample_1.doc", "ATN_sample_<number>.doc", )
				Expect(matched).Should(BeTrue())
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("And appears more than once", func() {
				matched, err := parser.HasMatch("ATN_sample_1_2_3.doc", "ATN_sample_<number>_<number>_<number>.doc", )
				Expect(matched).Should(BeTrue())
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("And they are repeated", func() {
				matched, err := parser.HasMatch("ATN_sample_123456.doc", "ATN_sample_<number>.doc", )
				Expect(matched).Should(BeTrue())
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("And the string does not match with pattern", func() {
				matched, err := parser.HasMatch("ATN_sample_1_2_a.doc", "ATN_sample_<number>_<number>_<number>.doc", )
				Expect(matched).Should(BeFalse())
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		It("When string contains date", func() {
			matched, err := parser.HasMatch("ATN_sample_05_22_2017.doc", "ATN_sample_<mm>_<dd>_<aaaa>.doc")
			Expect(matched).Should(BeTrue())
			Expect(err).ShouldNot(HaveOccurred())
		})

		Context("When string contains any word", func() {
			It("As a date", func() {
				matched, err := parser.HasMatch("ATN_sample_05_22_2017.doc", "ATN_sample_<*>.doc")
				Expect(matched).Should(BeTrue())
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("As any text", func() {
				matched, err := parser.HasMatch("ATN_sample_document_02.doc", "ATN_<*>.doc")
				Expect(matched).Should(BeTrue())
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("As on file extension", func() {
				matched, err := parser.HasMatch("ATN_sample_document_02.doc", "ATN_<*>.doc<*>")
				Expect(matched).Should(BeTrue())
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("As on file extension appended", func() {
				matched, err := parser.HasMatch("ATN_sample_document_02.docx", "ATN_<*>.doc<*>")
				Expect(matched).Should(BeTrue())
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("And the string does not match with pattern", func() {
				matched, _ := parser.HasMatch("ATN_sample_document_02.doc", "ATN_<*>.doc<*>")
				Expect(matched).Should(BeFalse())
				//Expect(err).ShouldNot(HaveOccurred())
			})
		})

	})

})
