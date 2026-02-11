package main

import (
	"fmt"
	"os"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/Hirogava/Go-Log-Linter/analyzer"
)

func main() {
	fmt.Fprintln(os.Stderr, "LogLinter: log message convention analyzer")
	fmt.Fprintln(os.Stderr, "Version: 1.0.0")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Usage: loglint <directory>")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Checks:")
	fmt.Fprintln(os.Stderr, "  - Log messages must start with lowercase")
	fmt.Fprintln(os.Stderr, "  - Log messages must be in English")
	fmt.Fprintln(os.Stderr, "  - Log messages must not contain forbidden special characters")
	fmt.Fprintln(os.Stderr, "  - Log messages must not contain sensitive keywords")
	fmt.Fprintln(os.Stderr, "")

	singlechecker.Main(analyzer.Analyzer)
}
