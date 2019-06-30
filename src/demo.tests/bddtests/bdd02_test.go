package bddtests_test

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"demo.tests/bddtests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("[test.demo02] Demo", func() {
	Describe("[test.asserter.suite11] Desc", func() {
		Context("[suite11.context01] Ctx", func() {
			It("[suite11.case01] ShouldNot asserter", func() {
				const tag = "To:"
				Expect(1).ShouldNot(Equal(2), "%s: assert equal", tag)
			})

			It("[suite11.case02] BeXXX asserter", func() {
				const tag = "Should:"
				Expect(1 == 1).Should(BeTrue(), "%s: assert true", tag)
				Expect(0).Should(BeZero(), "%s: assert zero", tag)
				Expect(2).Should(BeNumerically("<", 10), "%s: assert #number < 10", tag)
				Expect(func() error {
					return nil
				}()).Should(BeNil(), "%s: assert func ret nil", tag)
			})

			It("[suite11.case03] string asserter", func() {
				Expect("hello").Should(ContainSubstring("ll"))
				Expect("zheng").Should(HavePrefix("zh"))
				Expect("027-1234").Should(MatchRegexp(`^\d{3}`))
			})

			It("[suite11.case04] slice asserter", func() {
				s := []int{}
				Expect(s).Should(BeEmpty())
				s = make([]int, 2, 4)
				Expect(s).Should(HaveLen(2))
				s[0] = 0
				s[1] = 1
				Expect(s).Should(ContainElement(1))
				s = append(s, 2)
				s = append(s, 3)
				Expect(s).Should(ConsistOf([]int{0, 1, 2, 3}))

				Expect([]string{"Foo", "FooBar"}).Should(ConsistOf(ContainSubstring("Bar"), "Foo"))
			})

			It("[suite11.case05] map asserter", func() {
				m := map[int]string{
					1: "one",
					2: "two",
				}
				Expect(m).Should(HaveKey(1))
				Expect(m).Should(HaveKeyWithValue(2, "two"))

				Expect(map[string]string{"Foo": "Bar", "BazFoo": "Duck"}).Should(
					HaveKey(MatchRegexp(`.+Foo$`)))
				Expect(map[string]int{"Foo": 2, "BazFoo": 4}).Should(
					HaveKeyWithValue(MatchRegexp(`.+Foo$`), BeNumerically(">", 3)))
			})

			It("[suite11.case06] And, Or in asserter", func() {
				Expect(2).To(And(BeNumerically(">", 1), BeNumerically("<", 3)), "assert number by and")
				Expect(2).To(Or(BeNumerically(">", 10), BeNumerically("<", 3)), "assert number by or")
			})
		})

		Context("[suite11.context02] Ctx", func() {
			It("[suite11.case07] error asserter", func() {
				fnMockErr := func(isOccur bool) error {
					if isOccur {
						return errors.New("mock error")
					}
					return nil
				}

				err := fnMockErr(false)
				Expect(err).Should(Succeed(), "failed: error occurred: %v\n", err)
				Expect(fnMockErr(true)).Should(HaveOccurred())
			})

			It("[suite11.case08] panic asserter", func() {
				fnMockPanic := func() {
					panic("mock panic")
				}
				Expect(fnMockPanic).Should(Panic())
			})

			It("[suite11.case09] file asserter", func() {
				Expect("./fibonacci.go").Should(BeAnExistingFile())
				Expect("./fibonacci.go").Should(BeARegularFile()) // file type asserter
			})

			It("[suite11.case10] async asserter", func() {
				fnRandomInt := func() int {
					randInt := rand.Intn(5)
					fmt.Println("get int number:", randInt)
					return randInt
				}

				timeout := time.Duration(3) * time.Second
				interval := time.Duration(500) * time.Millisecond
				Eventually(func() int {
					return fnRandomInt()
				}, timeout, interval).Should(Equal(3))
			})

			It("[suite11.case11] channel asserter", func() {
				ch := make(chan int)
				go func(ch chan<- int) {
					for i := 0; i < 6; i++ {
						time.Sleep(time.Duration(300) * time.Millisecond)
						ch <- i
					}
					close(ch)
				}(ch)

				go func(ch <-chan int) {
					for i := range ch {
						fmt.Println("channel ret value:", i)
					}
				}(ch)

				timeout := time.Duration(3) * time.Second
				interval := time.Duration(500) * time.Millisecond
				Eventually(ch, timeout, interval).Should(BeClosed())
			})
		})
	})

	Describe("[test.share.suite12] Share Asserter", func() {
		var data interface{}

		AssertTureBehavior := func() {
			It("[suite12.test01] data should not be nil", func() {
				Expect(data).ShouldNot(BeNil())
			})

			It("[suite12.test02] data should not be empty", func() {
				Expect(data).ShouldNot(BeEmpty())
			})
		}

		Context("[suite12.context01] String Verification", func() {
			BeforeEach(func() {
				data = "hello world"
			})
			AssertTureBehavior()
		})

		Context("[suite12.context02] Bytes Verification", func() {
			BeforeEach(func() {
				data = []byte("Golang")
			})
			AssertTureBehavior()
		})
	})

	Describe("[test.table.suite13] Cases Table", func() {
		Context("[suite13.context01] Ctx", func() {
			DescribeTable("[suite13.case01] > inequality",
				func(x int, y int, expected bool) {
					Expect(x > y).Should(Equal(expected))
				},
				Entry("x > y", 1, 0, true),
				Entry("x = y", 0, 0, false),
				Entry("x < y", 0, 1, false),
			)

			DescribeTable("[suite13.case02] add test func", funcTestAdd,
				Entry("1 + 1 = 2", 1, 1, 2),
				Entry("1 + -1 = 0", 1, -1, 0),
				Entry("-1 + -1 = -2", -1, -1, -2),
			)
		})
	})

	Describe("[test.benchmark.suite14] Benchmark", func() {
		Measure("[suite14.measure01] it should do something hard efficiently", func(b Benchmarker) {
			runtime := b.Time("runtime", func() {
				for i := 0; i < 10; i++ {
					bddtests.Fibonacci(30)
				}
			})

			Expect(runtime.Seconds()).Should(BeNumerically("<", 0.5),
				"SomethingHard() shouldn't take too long.")
		}, 3)
	})
})

func funcTestAdd(x, y, expected int) {
	Expect(x + y).Should(Equal(expected))
}
