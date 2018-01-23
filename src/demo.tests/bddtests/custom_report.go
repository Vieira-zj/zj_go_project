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

// CustomReport : custome report demo
type CustomReport struct {
	Tag string
}

// SpecSuiteWillBegin : hook
func (r *CustomReport) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {
	fmt.Printf("%s random seend: %d\n", r.Tag, config.RandomSeed)
}

// BeforeSuiteDidRun : hook
func (r *CustomReport) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {
}

// SpecWillRun : hook
func (r *CustomReport) SpecWillRun(specSummary *types.SpecSummary) {
}

// SpecDidComplete : hook
func (r *CustomReport) SpecDidComplete(specSummary *types.SpecSummary) {
}

// AfterSuiteDidRun : hook
func (r *CustomReport) AfterSuiteDidRun(setupSummary *types.SetupSummary) {
}

// SpecSuiteDidEnd : hook
func (r *CustomReport) SpecSuiteDidEnd(summary *types.SuiteSummary) {
	fmt.Printf("%s, total: %v\n", r.Tag, summary.NumberOfTotalSpecs)
}
