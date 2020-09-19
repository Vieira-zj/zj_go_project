package diskusage_test

import (
	"os"
	"path/filepath"
	"testing"

	mysvc "src/tools.app/services/diskusage"
)

func TestPrintFilesTree(t *testing.T) {
	t.Logf("Case01: print files tree for a directory by limit levels.")
	path := filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files")

	diskUage := mysvc.NewDiskUsage()
	t.Logf("Step01: check print files tree for 2 levels.")
	if err := diskUage.PrintFilesTree(path, 2); err != nil {
		t.Fatal(err)
	}

	t.Skip("skip TestPrintFilesTree for max levels.")
	t.Logf("Step02: check print files tree for max levels.")
	if err := diskUage.PrintFilesTree(path, -1); err != nil {
		t.Fatal(err)
	}
}
