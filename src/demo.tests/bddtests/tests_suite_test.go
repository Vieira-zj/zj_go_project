package bddtests_test

import (
	"testing"

	"demo.tests/bddtests"

	"github.com/onsi/ginkgo/reporters"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTests(t *testing.T) {
	RegisterFailHandler(Fail)
	// RunSpecs(t, "Tests Suite")

	isJunitReport := false
	if isJunitReport {
		RunSpecsWithDefaultAndCustomReporters(t, "Tests Suite", []Reporter{reporters.NewJUnitReporter("xml-report-path")})
	} else {
		report := &bddtests.CustomReport{Tag: "*CUSTOM_REPORT =>"}
		RunSpecsWithDefaultAndCustomReporters(t, "Tests Suite", []Reporter{report})
	}
}
