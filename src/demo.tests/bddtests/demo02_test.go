package bddtests_test

import (
	"demo.tests/bddtests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func init() {
	By("init from demo02_test.go")
}

var _ = Describe("TestDemo02", func() {
	Describe("Asserter tests", func() {
		Context("Context", func() {
			It("[demo02.asserter] NotTo", func() {
				Expect(1).NotTo(Equal(2))
			})

			It("[demo02.asserter] BeZero", func() {
				Expect(0).To(BeZero())
			})

			It("[demo02.asserter] Or", func() {
				Expect(2).To(BeNumerically(">", 1), BeNumerically("<", 3))
			})
		})

		Context("Context", func() {
			It("[demo02.asserter] BeTrue", func() {
				Expect(true).Should(BeTrue())
			})
		})
	})

	Describe("Test external", func() {
		Context("Test context", func() {
			// cmd: ginkgo -v --focus="measure" src/demo.tests/bddtests/
			Measure("[demo02.measure] it should do something hard efficiently", func(b Benchmarker) {
				runtime := b.Time("runtime", func() {
					ouput := bddtests.Fibonacci(30)
					Expect(ouput).To(Equal(2178309))
				})

				Expect(runtime.Seconds()).Should(BeNumerically("<", 0.5),
					"SomethingHard() shouldn't take too long.")
			}, 10)

			DescribeTable("[demo02.DescribeTable] the > inequality",
				func(x int, y int, expected bool) {
					Expect(x > y).To(Equal(expected))
				},
				Entry("x > y", 1, 0, true),
				Entry("x = y", 0, 0, false),
				Entry("x < y", 0, 1, false),
			)

			DescribeTable("[demo02.DescribeTable.fn] the add function", fnAddTest,
				Entry("1 + 1 = 2", 1, 1, 2),
				Entry("1 + -1 = 0", 1, -1, 0),
				Entry("-1 + -1 = -2", -1, -1, -2),
			)
		})
	})
})

func fnAddTest(x, y, expected int) {
	Expect(x + y).To(Equal(expected))
}
