package main

import (
	"fmt"
	"io"
	"os"

	"github.com/vextmdev/vexErrScope/internal/explain"
	"github.com/vextmdev/vexErrScope/internal/render"
)

const usage = `vexErrScope - Scan logs and stack traces for error patterns

Usage:
  ./vexErrScope explain <file>    Analyze error in file
  ./vexErrScope explain           Read from stdin
  ./vexErrScope help              Show this help

Examples:
  ./vexErrScope explain crash.log
  cat error.txt | vexErrScope explain
`

func main() {
	if err := run(os.Args[1:], os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	if len(args) == 0 || args[0] == "help" || args[0] == "-h" || args[0] == "--help" {
		fmt.Fprint(stdout, usage)
		return nil
	}

	if args[0] != "explain" {
		return fmt.Errorf("unknown command: %s\nRun 'vexErrScope help' for usage", args[0])
	}

	var reader io.Reader

	if len(args) > 1 {
		file, err := os.Open(args[1])
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()
		reader = file
	} else {
		reader = stdin
	}

	explanation, err := explain.Explain(reader)
	if err != nil {
		return fmt.Errorf("failed to analyze: %w", err)
	}

	return render.Render(stdout, explanation)
}
