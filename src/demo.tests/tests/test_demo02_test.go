package tests_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TestDemo02", func() {
	Describe("Assert tests", func() {
		Context("Context: To and NotTo", func() {
			It("[demo02] assert NotTo", func() {
				Expect(1).NotTo(Equal(2))
			})

			It("[demo02] assert Zero", func() {
				Expect(0).To(BeZero())
			})
		})

		Context("Context: Should and ShouldNot", func() {
			It("[demo02] assert Be", func() {
				Expect(true).Should(BeTrue())
			})
		})
	})

	// Describe("Test benchmark", func() {
	// 	Context("Test context", func() {
	// 		// cmd: ginkgo -v --focus="measure" src/demo.tests/tests/
	// 		Measure("[demo02] [measure] it should do something hard efficiently", func(b Benchmarker) {
	// 			runtime := b.Time("runtime", func() {
	// 				ouput := tests.Fibonacci(30)
	// 				Expect(ouput).To(Equal(2178309))
	// 			})

	// 			Expect(runtime.Seconds()).Should(BeNumerically("<", 0.5),
	// 				"SomethingHard() shouldn't take too long.")
	// 		}, 10)
	// 	})
	// })
})
