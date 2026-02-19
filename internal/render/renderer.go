package render

import (
	"fmt"
	"io"
	"strings"

	"github.com/vextmdev/vexErrScope/internal/explain"
)

// Render outputs a formatted explanation to the writer.
func Render(w io.Writer, exp *explain.Explanation) error {
	if exp == nil {
		fmt.Fprintln(w, "No error pattern detected.")
		return nil
	}

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Error Type:        %s\n", exp.ErrorType))
	sb.WriteString(fmt.Sprintf("Detected Language: %s\n", exp.Language))
	sb.WriteString(fmt.Sprintf("Root Cause:        %s\n", exp.RootCause))
	sb.WriteString(fmt.Sprintf("Suggested Fix:     %s\n", exp.SuggestFix))
	sb.WriteString(fmt.Sprintf("Confidence:        %s\n", exp.Confidence))

	if exp.Location != nil {
		sb.WriteString("\nLocation:\n")
		if exp.Location.Function != "" {
			sb.WriteString(fmt.Sprintf("  Function: %s\n", exp.Location.Function))
		}
		if exp.Location.File != "" {
			sb.WriteString(fmt.Sprintf("  File:     %s\n", exp.Location.File))
		}
		if exp.Location.Line != "" {
			sb.WriteString(fmt.Sprintf("  Line:     %s\n", exp.Location.Line))
		}
	}

	_, err := io.WriteString(w, sb.String())
	return err
}
