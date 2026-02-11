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

func isZapSugar(call *ast.CallExpr, pass *analysis.Pass) bool {
    sel, ok := call.Fun.(*ast.SelectorExpr)
    if !ok {
        return false
    }

    if innerCall, ok := sel.X.(*ast.CallExpr); ok {
        if innerSel, ok := innerCall.Fun.(*ast.SelectorExpr); ok {
			if innerSel.Sel.Name == "Sugar" {
				if selection := pass.TypesInfo.Selections[innerSel]; selection != nil {
					recv := selection.Recv()
					if getPackagePath(recv) == "go.uber.org/zap" {
						return true
					}
				}

				if obj := pass.TypesInfo.Uses[innerSel.Sel]; obj != nil {
					if pkg := obj.Pkg(); pkg != nil {
						path := pkg.Path()
						if path == "go.uber.org/zap" {
							return true
						}
					}
				}
			}
        }
    }

	return false
}
