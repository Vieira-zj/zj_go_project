package bddtests_test

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var myFlag string

func init() {
	fmt.Println("init from demo01_test.go")
	fmt.Printf("$GOROOT: %s\n", os.Getenv("GOROOT"))
	fmt.Printf("$GOPATH: %v\n", strings.Split(os.Getenv("GOPATH"), ":")[0:3])

	flag.StringVar(&myFlag, "myFlag", "default", "myFlag is used to control my behavior")
}

// cmd: ginkgo -v --focus="demo01" src/demo.tests/bddtests/
var _ = Describe("TestDemo01", func() {
	var myText string

	BeforeSuite(func() {
		fmt.Println("TEST: exec BeforeSuite")
	})

	AfterSuite(func() {
		log.Println("\nTEST: exec AfterSuite")
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

	Describe("Desc", func() {
		Context("Test context", func() {
			It("[demo01.asserter] text is not null", func() {
				GinkgoWriter.Write([]byte("TEST: run test01\n"))
				By("sub step description")
				Expect(myText != "").Should(BeTrue(), "Failed, not null")
			})

			It("[demo01.asserter] text length should be 4", func() {
				GinkgoWriter.Write([]byte("TEST: run test02\n"))
				Expect(len(myText)).To(Equal(4), "Failed, text length = 4")
			})
		})

		Context("Test context", func() {
			BeforeEach(func() {
				fmt.Println("exec BeforeEach in defer test")
			})

			AfterEach(func() {
				fmt.Println("exec AfterEach in defer test")
			})

			It("[demo01.snyc] run parallel", func() {
				By("parallel test: start")
				var wg sync.WaitGroup
				for i := 0; i < 3; i++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						defer GinkgoRecover()
						fmt.Println("myPrintRoutine start")
						time.Sleep(time.Duration(1) * time.Second)
						Fail("make failed test")
						fmt.Println("myPrintRoutine end")
					}()
				}
				wg.Wait()
				fmt.Println("parallel test: done")
			})

			It("[demo01.sync.recover] run parallel", func() {
				By("parallel test: start")
				var wg sync.WaitGroup
				for i := 0; i < 3; i++ {
					wg.Add(1)
					go fnMyPrint(&wg)
				}
				wg.Wait()
				fmt.Println("parallel test: done")
			})
		})

		Context("Test context", func() {
			It("[demo01.routine.done] run parallel", func(done Done) {
				go func() {
					time.Sleep(time.Duration(500) * time.Millisecond)
					Expect(true).To(BeTrue())
					fmt.Println("routine done.")
					close(done)
				}()
				fmt.Println("parallel test: done")
			})

			It("[demo01.routine.wait] run parallel", func() {
				const count = 3
				for i := 0; i < count; i++ {
					go routineMyPrint()
				}
				time.Sleep(time.Duration(2) * time.Second)
				Fail("make failed")
				fmt.Println("parallel test: done")
			})
		})
	})

	Describe("desc", func() {
		BeforeEach(func() {
			fmt.Println("exec BeforeEach in flag test")
		})

		AfterEach(func() {
			fmt.Println("exec AfterEach in flag test")
		})

		It("[demo01.defer] Marking Specs as Failed", func() {
			fmt.Println("TEST: run test03")
			defer func() {
				fmt.Println("defer test")
			}()
			Fail("mark failed in test")
			fmt.Println("message after make failed") // skipped
		})

		// cmd: ginkgo -v --focus="flag" src/demo.tests/bddtests/ -- -myFlag="flagtext"
		It("[demo01.flag] get string flag text", func() {
			By("my flag value: " + myFlag)
			Expect(myFlag).To(MatchRegexp("flagtext|default"))
		})
	})
})

func fnMyPrint(wg *sync.WaitGroup) {
	defer wg.Done()
	defer GinkgoRecover()
	fmt.Println("myPrintRoutine start")
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("myPrintRoutine end")
}

var i int
var m *sync.Mutex

func routineMyPrint() {
	defer GinkgoRecover()

	m = new(sync.Mutex)
	m.Lock()
	i++
	fmt.Printf("run test routine at: %d\n", i)
	m.Unlock()

	time.Sleep(time.Duration(500) * time.Millisecond)
	Fail("make failed in routine")
	fmt.Println("routineMyPrint end")
}
