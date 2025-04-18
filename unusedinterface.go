package unusedinterface

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/gostaticanalysis/ident"
)

const doc = "unusedinterface find interface that are not used in their package"

// Analyzer find interface definition that are not used in their package.
var Analyzer = &analysis.Analyzer{
	Name: "unusedinterface",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer, ident.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	ifaces := make(map[string]bool)
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{(*ast.TypeSpec)(nil)}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.TypeSpec:
			if _, ok := n.Type.(*ast.InterfaceType); ok {
				if n.Name != nil {
					ifaces[n.Name.Name] = true
				}
			}
		}
	})

	if pass.ResultOf[ident.Analyzer] == nil {
		return nil, nil
	}
	m := pass.ResultOf[ident.Analyzer].(ident.Map)
	for o := range m {
		if _, ok := ifaces[o.Name()]; ok {
			if len(m[o]) == 1 {
				n := m[o][0]
				pass.Reportf(n.Pos(), "interface %s is defined but not used within the same package", n.Name)
			}
		}
	}
	return nil, nil
}
