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

	Describe("Test string", func() {
		Context("Test context", func() {
			It("[demo01] text is not null", func() {
				GinkgoWriter.Write([]byte("TEST: run test01\n"))
				By("sub step description")
				Expect(myText != "").Should(BeTrue(), "Failed, not null")
			})
		})

		Context("Test context", func() {
			It("[demo01] text length should be 4", func() {
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

			It("[demo01] [failed] [defertest] Marking Specs as Failed", func() {
				fmt.Println("TEST: run test03")
				defer func() {
					By("Defer test")
				}()
				Fail("Mark failed")
				By("message after make failed") // skipped
			})

			It("[demo01] [parallel] [recover] run parallel", func() {
				By("parallel test: start")
				var wg sync.WaitGroup
				for i := 0; i < 10; i++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						defer GinkgoRecover()
						fmt.Println("myPrintRoutine start")
						time.Sleep(time.Duration(3) * time.Second)
						fmt.Println("myPrintRoutine end")
					}()
				}
				wg.Wait()
				By("parallel test: done")
			})

			It("[demo01] [parallel fn] [recover] run parallel", func() {
				By("parallel test: start")
				var wg sync.WaitGroup
				for i := 0; i < 10; i++ {
					wg.Add(1)
					go fnMyPrint(&wg)
				}
				wg.Wait()
				By("parallel test: done")
			})
		})
	})

	Describe("Test flag", func() {
		BeforeEach(func() {
			fmt.Println("exec BeforeEach in flag test")
		})

		AfterEach(func() {
			fmt.Println("exec AfterEach in flag test")
		})

		// cmd: ginkgo -v --focus="flagtest" src/demo.tests/bddtests/ -- -myFlag="flagtext"
		It("[demo01] [flagtest] get string flag text", func() {
			By("my flag value: " + myFlag)
			Expect(myFlag).To(MatchRegexp("flagtext|default"))
		})
	})
})

func fnMyPrint(wg *sync.WaitGroup) {
	defer wg.Done()
	defer GinkgoRecover()
	fmt.Println("myPrintRoutine start")
	time.Sleep(time.Duration(3) * time.Second)
	fmt.Println("myPrintRoutine end")
}
