package analyzer

import (
	"flag"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// config хранит настройки линтера
var config = DefaultConfig()

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "checks log messages for conventions like lowercase, english, no special chars, and no sensitive keywords",
	Run:  run,
	Flags: func() flag.FlagSet {
		fs := flag.NewFlagSet("loglint", flag.ContinueOnError)
		config.RegisterFlags(fs)
		return *fs
	}(),
}

func New() *analysis.Analyzer {
	return Analyzer
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if isZapSugar(call, pass) {
				return true
			}

			if !isLogCall(call, pass) {
				return true
			}

			checkLogCall(pass, call, config)
			return true
		})
	}

	return nil, nil
}
