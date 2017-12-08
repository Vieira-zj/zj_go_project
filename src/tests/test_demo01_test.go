package tests_test

import (
	"flag"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var myFlag string

func init() {
	// cmd: ginkgo -v src/tests/ -- -myFlag="flag text"
	flag.StringVar(&myFlag, "myFlag", "default", "myFlag is used to control my behavior")
}

var _ = Describe("TestDemo01", func() {
	var myText string

	BeforeSuite(func() {
		GinkgoWriter.Write([]byte("TEST: exec BeforeSuite\n"))
	})

	AfterSuite(func() {
		GinkgoWriter.Write([]byte("TEST: exec AfterSuite\n"))
		By("my flag value: " + myFlag)
	})

	BeforeEach(func() {
		GinkgoWriter.Write([]byte("TEST: exec BeforeEach\n"))
		myText = "test"
	})

	AfterEach(func() {
		GinkgoWriter.Write([]byte("TEST: exec AfterEach\n"))
	})

	JustBeforeEach(func() {
		GinkgoWriter.Write([]byte("TEST: exec JustBeforeEach\n"))
	})

	Describe("Test string", func() {
		Context("Test context", func() {
			It("string is not null", func() {
				GinkgoWriter.Write([]byte("TEST: run test01\n"))
				By("sub step description")
				Expect(myText != "").Should(BeTrue())
			})
		})

		Context("Test context", func() {
			It("string length should be 4", func() {
				GinkgoWriter.Write([]byte("TEST: run test02\n"))
				Expect(len(myText)).To(Equal(4))
			})
		})

		Context("Test context", func() {
			It("Marking Specs as Failed", func() {
				GinkgoWriter.Write([]byte("TEST: run test03\n"))
				Fail("Mark failed")
			})
		})
	})
})
