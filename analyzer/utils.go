package analyzer

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func isLogCall(call *ast.CallExpr, pass *analysis.Pass) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	if selection := pass.TypesInfo.Selections[sel]; selection != nil {
		str := selection.Recv().String()

		if strings.Contains(str, "log/slog") || strings.Contains(str, "go.uber.org/zap") {
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

func getPackagePath(t types.Type) string {
	if ptr, ok := t.(*types.Pointer); ok {
		t = ptr.Elem()
	}

	if named, ok := t.(*types.Named); ok {
		if named.Obj().Pkg() != nil {
			return named.Obj().Pkg().Path()
		}
	}

	return ""
}
