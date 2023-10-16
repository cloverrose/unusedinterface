package main

import (
	"github.com/cloverrose/unusedinterface"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(unusedinterface.Analyzer) }
