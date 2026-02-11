package analyzer

import (
	"go/ast"
	"go/token"
	"strconv"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "checks log messages",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
    for _, file := range pass.Files {
        ast.Inspect(file, func(n ast.Node) bool {
            call, ok := n.(*ast.CallExpr)
            if !ok {
                return true
            }

			if !isLogCall(call, pass) {
				return true
			}

			checkLogCall(pass, call)
            return true
        })
    }

    return nil, nil
}

func checkLogCall(pass *analysis.Pass, call *ast.CallExpr) {
	if len(call.Args) == 0 {
		return
	}

	msgLit, ok := call.Args[0].(*ast.BasicLit)
	if !ok || msgLit.Kind != token.STRING {
		return
	}

	msg, err := parseStringLiteral(msgLit.Value)
	if err != nil {
		return
	}
	
	for _, rule := range rules {
		if err := rule.Check(msg); err != nil {
			pass.Reportf(call.Pos(), "%s", err.Error())
		}
	}
}

func parseStringLiteral(s string) (string, error) {
	if len(s) < 2 {
		return "", nil
	}

	return strconv.Unquote(s)
}
