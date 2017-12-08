package tests_test

import (
	"tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TestDemo02", func() {
	Describe("Assert tests", func() {
		Context("Context: To and NotTo", func() {
			It("assert NotTo", func() {
				Expect(1).NotTo(Equal(2))
			})

			It("assert Zero", func() {
				Expect(0).To(BeZero())
			})
		})

		Context("Context: Should and ShouldNot", func() {
			It("assert Be", func() {
				Expect(true).Should(BeTrue())
			})
		})
	})

	FDescribe("Test benchmark", func() {
		Context("Test context", func() {
			Measure("it should do something hard efficiently", func(b Benchmarker) {
				runtime := b.Time("runtime", func() {
					ouput := tests.Fibonacci(20)
					Expect(ouput).To(Equal(17711))
				})

				Expect(runtime.Seconds()).Should(
					BeNumerically("<", 0.2),
					"SomethingHard() shouldn't take too long.")
			}, 10)
		})
	})
})
