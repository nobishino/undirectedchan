package undirectedchan

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "undirectedchan is a vet tool that reports usage of undirected channel type as function arguments or results."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "undirectedchan",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		checkFuncType(pass, n)
	})

	return nil, nil
}

func checkFuncType(pass *analysis.Pass, n ast.Node) {
	var typ *ast.FuncType
	switch n := n.(type) {
	case *ast.FuncDecl:
		typ = n.Type
	case *ast.FuncLit:
		typ = n.Type
	default:
		return
	}
	if typ.Params != nil {
		for _, arg := range typ.Params.List {
			chType, ok := arg.Type.(*ast.ChanType)
			if !ok {
				continue
			}
			if chType.Dir == ast.RECV|ast.SEND {
				pass.Reportf(n.Pos(), "channel argument should be directed")
			}
		}
	}
	if typ.Results != nil {
		for _, res := range typ.Results.List {
			chType, ok := res.Type.(*ast.ChanType)
			if !ok {
				continue
			}
			if chType.Dir == ast.RECV|ast.SEND {
				pass.Reportf(n.Pos(), "channel result should be directed")
			}
		}
	}
}
