package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
	"strconv"

	"github.com/Hirogava/Go-Log-Linter/analyzer/rules"
	"golang.org/x/tools/go/analysis"
)

func checkLogCall(pass *analysis.Pass, call *ast.CallExpr, cfg *Config) {
	if len(call.Args) == 0 {
		return
	}

	var msg string
	var err error

	msgLit, ok := call.Args[0].(*ast.BasicLit)
	if ok && msgLit.Kind == token.STRING {
		msg, err = parseStringLiteral(msgLit.Value)
		if err != nil {
			return
		}
	}

	msgBin, ok := call.Args[0].(*ast.BinaryExpr)
	if ok {
		left, ok := msgBin.X.(*ast.BasicLit)
		if ok && left.Kind == token.STRING {
			msg, err = parseStringLiteral(left.Value)
			if err != nil {
				return
			}
		}

		right, ok := msgBin.Y.(*ast.BasicLit)
		if ok && right.Kind == token.STRING {
			rightMsg, err := parseStringLiteral(right.Value)
			if err == nil {
				msg += rightMsg
			}
		}
	}

	if msg == "" {
		return
	}

	ruleChecker := rules.NewChecker(cfg)
	for _, rule := range ruleChecker.Rules() {
		if ruleErr := rule.Check(msg); ruleErr != nil {
			pass.Reportf(call.Pos(), "%s", ruleErr.Error())
		}
	}
}

func parseStringLiteral(s string) (string, error) {
	if len(s) < 2 {
		return "", nil
	}

	return strconv.Unquote(s)
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
