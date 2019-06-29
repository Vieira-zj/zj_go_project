package bddtests_test

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var myFlag string

func init() {
	fmt.Println("init in demo01_test.go")
	fmt.Printf("$GOROOT: %s\n", os.Getenv("GOROOT"))
	fmt.Printf("$GOPATH: %v\n", os.Getenv("GOPATH"))

	flag.StringVar(&myFlag, "myFlag", "default", "myFlag is used to control my behavior")
}

// cmd: ginkgo -v --focus="demo01" src/demo.tests/bddtests/
var _ = Describe("[test.demo01]", func() {
	var myText string

	BeforeSuite(func() {
		fmt.Println("HOOK: BeforeSuite run")
	})

	AfterSuite(func() {
		log.Println("\nHOOK: AfterSuite run")
	})

	BeforeEach(func() {
		GinkgoWriter.Write([]byte("HOOK: BeforeEach run\n"))
		myText = "test"
	})

	AfterEach(func() {
		GinkgoWriter.Write([]byte("HOOK: AfterEach run\n"))
	})

	JustBeforeEach(func() {
		GinkgoWriter.Write([]byte("HOOK: JustBeforeEach run\n"))
	})

	Describe("[test.hooks.suite01] Desc", func() {
		// panic: You may only call BeforeSuite once!
		// BeforeSuite(func() {
		// 	fmt.Println("TEST: try exec BeforeSuite in sub")
		// })

		BeforeEach(func() {
			fmt.Println("HOOK: BeforeEach run in sub")
		})

		JustBeforeEach(func() {
			fmt.Println("HOOK: JustBeforeEach run in sub")
		})

		It("[suite01.case01] test return in It()", func() {
			fmt.Println("case01 test start")
			isRet := true
			if isRet {
				fmt.Println("case01 return")
				return
			} else {
				fmt.Println("case01 asserter")
				Expect(true).Should(BeTrue())
			}
		})
	})

	Describe("[test.asserter.suite02] Desc", func() {
		It("[suite02.case01] text is not null", func() {
			GinkgoWriter.Write([]byte("TEST: asserter test01\n"))
			By("case01 step description")
			Expect(myText).ShouldNot(BeEmpty(), "Failed: text is empty")
		})

		It("[suite02.case02] text length should be 4", func() {
			GinkgoWriter.Write([]byte("TEST: asserter test02\n"))
			Expect(len(myText)).To(Equal(4), "Failed: text length != 4")
		})
	})

	Describe("[test.routine.suite03] Desc", func() {
		Context("[suite03.context01] Ctx", func() {
			It("[suite03.case01] run routine, and wait done", func(done Done) {
				go func() {
					time.Sleep(time.Duration(2) * time.Second)
					Expect(true).To(BeTrue())
					fmt.Println("sub routine done")
					close(done)
				}()
				fmt.Println("routine test: done")
			}, 3) // timeout = 3s

			It("[suite03.case02] run routines with lock", func() {
				const count = 3
				for i := 0; i < count; i++ {
					go routineMyPrint()
				}
				time.Sleep(time.Duration(2) * time.Second)
				fmt.Println("routine test: done")
			})
		})

		Context("[suite03.context02] Ctx", func() {
			BeforeEach(func() {
				fmt.Println("HOOK: BeforeEach in context02")
			})

			AfterEach(func() {
				fmt.Println("HOOK: AfterEach in context02")
			})

			It("[suite03.case03] run failed routines, and recover", func() {
				fmt.Println("suite03.case03 test start")
				var wg sync.WaitGroup
				for i := 0; i < 3; i++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						defer GinkgoRecover()
						fmt.Println("sub routine start")
						time.Sleep(time.Second)
						Fail("marked failed in routine")
						fmt.Println("sub routine end")
					}()
				}
				wg.Wait()
				fmt.Println("suite03.case03 test end")
			})

			It("[suite03.case04] run failed routines, and recover", func() {
				fmt.Println("suite03.case04 test start")
				var wg sync.WaitGroup
				for i := 0; i < 3; i++ {
					wg.Add(1)
					go funcMyPrint(&wg)
				}
				wg.Wait()
				Expect("test").Should(BeEmpty(), "skip")
				fmt.Println("suite03.case04 test end")
			})
		})
	})

	Describe("[test.flag.suite04] Desc", func() {
		BeforeEach(func() {
			fmt.Println("HOOK: BeforeEach in flagtest")
		})

		AfterEach(func() {
			fmt.Println("HOOK: AfterEach in flagtest")
		})

		It("[suite04.case01] Marking Specs as Failed", func() {
			fmt.Println("run fake test start")
			defer func() {
				fmt.Println("defer test")
			}()
			Fail("mark failed in test")
			fmt.Println("run fake test end") // skipped
		})

		// cmd: ginkgo -v --focus="suite04.case02" src/demo.tests/bddtests/ -- -myFlag="test"
		It("[suite04.case02] get text from input flag", func() {
			fmt.Println("flag text:", myFlag)
			Expect(myFlag).To(MatchRegexp("test|default"))
		})
	})
})

var i int
var m *sync.Mutex

func routineMyPrint() {
	defer GinkgoRecover()

	m = new(sync.Mutex)
	m.Lock()
	i++
	fmt.Printf("routine run at: %d\n", i)
	m.Unlock()
	time.Sleep(time.Duration(500) * time.Millisecond)
	fmt.Println("[routineMyPrint] end")
}

func funcMyPrint(wg *sync.WaitGroup) {
	defer wg.Done()
	defer GinkgoRecover()
	fmt.Println("funcMyPrint routine start")
	time.Sleep(time.Second)
	Fail("mark failed in funcMyPrint routine")
	fmt.Println("funcMyPrint routine end")
}
