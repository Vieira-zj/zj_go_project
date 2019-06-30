package bddtests

import (
	"fmt"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

// type Reporter interface {
//     SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary)
//     BeforeSuiteDidRun(setupSummary *types.SetupSummary)
//     SpecWillRun(specSummary *types.SpecSummary)
//     SpecDidComplete(specSummary *types.SpecSummary)
//     AfterSuiteDidRun(setupSummary *types.SetupSummary)
//     SpecSuiteDidEnd(summary *types.SuiteSummary)
// }

// CustomReport : custom testing report
type CustomReport struct {
	Tag string
}

// SpecSuiteWillBegin : hook
func (r *CustomReport) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {
	config.RandomSeed = 666
	fmt.Println("[SpecSuiteWillBegin] hook start")
	fmt.Printf("%s random seed: %d\n", r.Tag, config.RandomSeed)
}

// BeforeSuiteDidRun : hook
func (r *CustomReport) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {
	fmt.Println("[BeforeSuiteDidRun] hook start")
}

// SpecWillRun : hook
func (r *CustomReport) SpecWillRun(specSummary *types.SpecSummary) {
	// run before each It() test case
	// fmt.Println("[SpecWillRun] hook start")
}

// SpecDidComplete : hook
func (r *CustomReport) SpecDidComplete(specSummary *types.SpecSummary) {
	// run after each It() test case
	// fmt.Println("[SpecDidComplete] hook start")
}

// AfterSuiteDidRun : hook
func (r *CustomReport) AfterSuiteDidRun(setupSummary *types.SetupSummary) {
	fmt.Println("[AfterSuiteDidRun] hook start")
}

// SpecSuiteDidEnd : hook
func (r *CustomReport) SpecSuiteDidEnd(summary *types.SuiteSummary) {
	fmt.Println("[SpecSuiteDidEnd] hook start")
	fmt.Printf("%s test cases count: %v\n", r.Tag, summary.NumberOfTotalSpecs)
	fmt.Printf("%s test cases skipped: %v\n", r.Tag, summary.NumberOfSkippedSpecs)
}

// NewCustomReport returns custom testing report instance.
func NewCustomReport(tag string) *CustomReport {
	return &CustomReport{Tag: tag}
}
