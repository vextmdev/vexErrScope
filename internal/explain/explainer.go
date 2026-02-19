package explain

import (
	"io"

	"github.com/vextmdev/vexErrScope/internal/analyze"
	"github.com/vextmdev/vexErrScope/internal/parse"
)

// Explanation represents the complete explanation of an error.
type Explanation struct {
	ErrorType  string
	Language   string
	RootCause  string
	SuggestFix string
	Confidence string
	Location   *Location
}

// Location represents where the error occurred.
type Location struct {
	Function string
	File     string
	Line     string
}

// Explain parses input and returns a structured explanation.
func Explain(r io.Reader) (*Explanation, error) {
	parsed, err := parse.Parse(r)
	if err != nil {
		return nil, err
	}

	if parsed == nil {
		return nil, nil
	}

	analysis := analyze.Analyze(parsed)
	if analysis == nil {
		return nil, nil
	}

	explanation := &Explanation{
		ErrorType:  analysis.ErrorType,
		Language:   analysis.Language,
		RootCause:  analysis.RootCause,
		SuggestFix: analysis.SuggestFix,
		Confidence: analysis.Confidence,
	}

	if analysis.StackFrame != nil {
		explanation.Location = &Location{
			Function: analysis.StackFrame.Function,
			File:     analysis.StackFrame.File,
			Line:     analysis.StackFrame.Line,
		}
	}

	return explanation, nil
}
