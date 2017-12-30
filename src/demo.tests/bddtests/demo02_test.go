package bddtests_test

import (
	"errors"
	"fmt"
	"time"

	"demo.tests/bddtests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("TestDemo02", func() {
	Describe("Asserter tests", func() {
		Context("Context, part1", func() {
			It("[demo02.asserter.part1] Not", func() {
				Expect(1).ShouldNot(Equal(2), "assert equal")
			})

			It("[demo02.asserter.part1] BeXXX", func() {
				Expect(1 == 1).To(BeTrue(), "assert bool")
				Expect(0 == 1).To(BeFalse(), "assert bool")
				Expect(0).To(BeZero(), "assert zero")
				Expect(func() error {
					return nil
				}()).To(BeNil(), "assert nil")
			})

			It("[demo02.asserter.part1] string", func() {
				Expect("hello").Should(ContainSubstring("ll"), "assert sub string")
				Expect("zheng").Should(HavePrefix("zh"), "assert string prefix")
				Expect("027-1234").Should(MatchRegexp("^\\d{3}"), "regexp assert")
			})

			It("[demo02.asserter.part1] slice", func() {
				s := []int{}
				Expect(s).Should(BeEmpty(), "assert slice empty")
				s = make([]int, 2, 4)
				Expect(s).Should(HaveLen(2), "assert slice length")
				s[0] = 0
				s[1] = 1
				Expect(s).Should(ContainElement(1), "assert slice element")
				s = append(s, 2)
				s = append(s, 3)
				Expect(s).Should(ConsistOf([]int{0, 1, 2, 3}), "assert slice elements")

				Expect([]string{"Foo", "FooBar"}).Should(
					ConsistOf(ContainSubstring("Bar"), "Foo"), "assert slice elements")
			})

			It("[demo02.asserter.part1] map", func() {
				m := map[int]string{
					1: "one",
					2: "two",
				}
				Expect(m).Should(HaveKey(1), "assert map key")
				Expect(m).Should(HaveKeyWithValue(2, "two"), "assert map key-value")

				Expect(map[string]string{"Foo": "Bar", "BazFoo": "Duck"}).Should(
					HaveKey(MatchRegexp(`.+Foo$`)))
				Expect(map[string]int{"Foo": 2, "BazFoo": 4}).Should(
					HaveKeyWithValue(MatchRegexp(`.+Foo$`), BeNumerically(">", 3)))
			})

			It("[demo02.asserter.part1] And, Or", func() {
				Expect(2).To(And(BeNumerically(">", 1), BeNumerically("<", 3)), "assert by and")
				Expect(2).To(Or(BeNumerically(">", 1), BeNumerically("<", 3)), "assert by or")
			})
		})

		Context("Context, part2", func() {
			It("[demo02.asserter.part2] error handle", func() {
				fnMockErr := func(isOccur bool) error {
					if isOccur {
						return errors.New("mock error")
					}
					return nil
				}
				Expect(fnMockErr(false)).To(Succeed())
				err := fnMockErr(true)
				Expect(err).To(HaveOccurred(), "FAILED: error occurred: %v\n", err)
			})

			It("[demo02.asserter.part2], panic handle", func() {
				fnMockPanic := func() {
					panic("mock panic")
				}
				Expect(fnMockPanic).Should(Panic(), "assert panic")
			})

			It("[demo02.asserter.part2] file handle", func() {
				Expect("./fibonacci.go").Should(BeAnExistingFile(), "assert file exist")
				Expect("./fibonacci.go").Should(BeARegularFile(), "assert file type")
			})

			It("[demo02.asserter.part2] async handle", func() {
				Eventually(func() []string {
					time.Sleep(time.Duration(500) * time.Millisecond)
					s := make([]string, 2)
					s[0] = "val1"
					s[1] = "val2"
					return s
				}, time.Duration(2)*time.Second).Should(HaveLen(2), "assert async fn")
			})

			It("[demo02.asserter.part2] channel handle", func() {
				ch := make(chan int)
				go func(ch chan<- int) {
					time.Sleep(time.Duration(500) * time.Millisecond)
					for i := 0; i < 5; i++ {
						ch <- i
					}
					close(ch)
				}(ch)
				Eventually(ch, time.Duration(800)*time.Millisecond).Should(BeClosed())
				for i := range ch {
					fmt.Printf("ret code %d\n", i)
				}
			})
		})
	})

	Describe("Test external", func() {
		// cmd: ginkgo -v --focus="measure" src/demo.tests/bddtests/
		Measure("[demo02.measure] it should do something hard efficiently", func(b Benchmarker) {
			runtime := b.Time("runtime", func() {
				ouput := bddtests.Fibonacci(30)
				Expect(ouput).To(Equal(2178309))
			})

			Expect(runtime.Seconds()).Should(BeNumerically("<", 0.5),
				"SomethingHard() shouldn't take too long.")
		}, 10)

		Context("Test context", func() {
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
