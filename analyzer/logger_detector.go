package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func isLogCall(call *ast.CallExpr, pass *analysis.Pass) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	if selection := pass.TypesInfo.Selections[sel]; selection != nil {
		recv := selection.Recv()

		if getPackagePath(recv) == "log/slog" || getPackagePath(recv) == "go.uber.org/zap" {
			return true
		}
	}

	if obj := pass.TypesInfo.Uses[sel.Sel]; obj != nil {
		if pkg := obj.Pkg(); pkg != nil {
			path := pkg.Path()

			if path == "log/slog" || path == "go.uber.org/zap" {
				return true
			}
		}
	}

	return false
}
