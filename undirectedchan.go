package undirectedchan

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "undirectedchan is ..."

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
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			for _, arg := range n.Type.Params.List {
				chType, ok := arg.Type.(*ast.ChanType)
				if !ok {
					continue
				}
				if chType.Dir == ast.RECV|ast.SEND {
					pass.Reportf(n.Pos(), "channel argument should be directed")
				}
				fmt.Printf("%#v\n", chType)
			}
		}

	})

	return nil, nil
}
