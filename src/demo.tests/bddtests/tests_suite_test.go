package bddtests_test

import (
	"fmt"
	"testing"

	"src/demo.tests/bddtests"

	"github.com/onsi/ginkgo/reporters"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTests(t *testing.T) {
	RegisterFailHandler(Fail)
	// RunSpecs(t, "Tests Suite")

	isJunitReport := false
	if isJunitReport {
		fmt.Println("Run tests with junit report")
		RunSpecsWithDefaultAndCustomReporters(t, "Tests Suite", []Reporter{reporters.NewJUnitReporter("xml-report-path")})
	} else {
		fmt.Println("Run tests with custom testing report")
		report := bddtests.NewCustomReport("***CUSTOM_REPORT***")
		RunSpecsWithDefaultAndCustomReporters(t, "Tests Suite", []Reporter{report})
	}
}
