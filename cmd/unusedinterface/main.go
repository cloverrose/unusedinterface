package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/cloverrose/unusedinterface"
)

func main() { unitchecker.Main(unusedinterface.Analyzer) }
