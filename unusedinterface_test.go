package unusedinterface_test

import (
	"testing"

	"github.com/cloverrose/unusedinterface"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, unusedinterface.Analyzer, "a")
}
