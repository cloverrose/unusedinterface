package unusedinterface_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/cloverrose/unusedinterface"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, unusedinterface.Analyzer, "a")
}
