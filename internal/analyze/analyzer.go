package analyze

import "github.com/vextmdev/vexErrScope/internal/parse"

// Analysis contains the result of error analysis.
type Analysis struct {
	ErrorType  string
	Language   string
	RootCause  string
	SuggestFix string
	Confidence string
	RawMessage string
	StackFrame *parse.StackFrame
}

// Pattern defines a known error pattern with explanation.
type Pattern struct {
	ErrorType  string
	RootCause  string
	SuggestFix string
	Confidence string
}

var goPatterns = map[string]Pattern{
	"nil_pointer_dereference": {
		ErrorType:  "Nil Pointer Dereference",
		RootCause:  "Attempted to access a field or method on a nil pointer. This typically happens when a pointer variable was never initialized or was explicitly set to nil.",
		SuggestFix: "Add nil checks before dereferencing pointers. Ensure the pointer is properly initialized before use. Consider using the comma-ok idiom for map lookups or type assertions.",
		Confidence: "high",
	},
	"invalid_memory_address": {
		ErrorType:  "Invalid Memory Address",
		RootCause:  "Program attempted to access memory at an invalid address, typically due to a nil pointer or corrupted pointer value.",
		SuggestFix: "Check for nil pointers before dereferencing. Verify that all pointers are properly initialized and not corrupted by concurrent access.",
		Confidence: "high",
	},
	"index_out_of_range": {
		ErrorType:  "Index Out of Range",
		RootCause:  "Attempted to access a slice or array element at an index that exceeds its bounds. The index is either negative or greater than or equal to the length.",
		SuggestFix: "Add bounds checking before accessing slice/array elements. Use len() to verify the slice has enough elements. Consider using range loops instead of index-based access.",
		Confidence: "high",
	},
	"runtime_error": {
		ErrorType:  "Runtime Error",
		RootCause:  "A general runtime error occurred during program execution.",
		SuggestFix: "Review the error message and stack trace for specific details. Common causes include nil pointers, out-of-bounds access, and division by zero.",
		Confidence: "medium",
	},
	"panic": {
		ErrorType:  "Panic",
		RootCause:  "The program encountered an unrecoverable error and panicked.",
		SuggestFix: "Review the panic message and stack trace to identify the root cause. Consider adding recover() in critical sections to handle panics gracefully.",
		Confidence: "low",
	},
}

// Analyze examines a parsed error and returns structured analysis.
func Analyze(parsed *parse.ParsedError) *Analysis {
	if parsed == nil {
		return nil
	}

	analysis := &Analysis{
		Language:   "Go",
		RawMessage: parsed.RawMessage,
		StackFrame: parsed.StackFrame,
	}

	if pattern, ok := goPatterns[parsed.ErrorType]; ok {
		analysis.ErrorType = pattern.ErrorType
		analysis.RootCause = pattern.RootCause
		analysis.SuggestFix = pattern.SuggestFix
		analysis.Confidence = pattern.Confidence
	} else {
		analysis.ErrorType = "Unknown Error"
		analysis.RootCause = parsed.RawMessage
		analysis.SuggestFix = "Review the error message and stack trace for more context."
		analysis.Confidence = "low"
	}

	return analysis
}
