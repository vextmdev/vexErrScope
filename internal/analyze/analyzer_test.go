package analyze

import (
	"testing"

	"github.com/vextmdev/vexErrScope/internal/parse"
)

func TestAnalyze_NilPointerDereference(t *testing.T) {
	parsed := &parse.ParsedError{
		RawMessage: "runtime error: invalid memory address or nil pointer dereference",
		ErrorType:  "nil_pointer_dereference",
		StackFrame: &parse.StackFrame{
			Function: "main.processData",
			File:     "/home/user/project/main.go",
			Line:     "42",
		},
	}

	analysis := Analyze(parsed)
	if analysis == nil {
		t.Fatal("expected analysis result, got nil")
	}

	if analysis.ErrorType != "Nil Pointer Dereference" {
		t.Errorf("expected error type 'Nil Pointer Dereference', got '%s'", analysis.ErrorType)
	}

	if analysis.Language != "Go" {
		t.Errorf("expected language 'Go', got '%s'", analysis.Language)
	}

	if analysis.Confidence != "high" {
		t.Errorf("expected confidence 'high', got '%s'", analysis.Confidence)
	}

	if analysis.RootCause == "" {
		t.Error("expected non-empty root cause")
	}

	if analysis.SuggestFix == "" {
		t.Error("expected non-empty suggested fix")
	}

	if analysis.StackFrame == nil {
		t.Error("expected stack frame to be preserved")
	}
}

func TestAnalyze_IndexOutOfRange(t *testing.T) {
	parsed := &parse.ParsedError{
		RawMessage: "runtime error: index out of range [5] with length 3",
		ErrorType:  "index_out_of_range",
	}

	analysis := Analyze(parsed)
	if analysis == nil {
		t.Fatal("expected analysis result, got nil")
	}

	if analysis.ErrorType != "Index Out of Range" {
		t.Errorf("expected error type 'Index Out of Range', got '%s'", analysis.ErrorType)
	}

	if analysis.Confidence != "high" {
		t.Errorf("expected confidence 'high', got '%s'", analysis.Confidence)
	}
}

func TestAnalyze_RuntimeError(t *testing.T) {
	parsed := &parse.ParsedError{
		RawMessage: "runtime error: integer divide by zero",
		ErrorType:  "runtime_error",
	}

	analysis := Analyze(parsed)
	if analysis == nil {
		t.Fatal("expected analysis result, got nil")
	}

	if analysis.ErrorType != "Runtime Error" {
		t.Errorf("expected error type 'Runtime Error', got '%s'", analysis.ErrorType)
	}

	if analysis.Confidence != "medium" {
		t.Errorf("expected confidence 'medium', got '%s'", analysis.Confidence)
	}
}

func TestAnalyze_UnknownError(t *testing.T) {
	parsed := &parse.ParsedError{
		RawMessage: "something completely unexpected",
		ErrorType:  "unknown_type",
	}

	analysis := Analyze(parsed)
	if analysis == nil {
		t.Fatal("expected analysis result, got nil")
	}

	if analysis.ErrorType != "Unknown Error" {
		t.Errorf("expected error type 'Unknown Error', got '%s'", analysis.ErrorType)
	}

	if analysis.Confidence != "low" {
		t.Errorf("expected confidence 'low', got '%s'", analysis.Confidence)
	}
}

func TestAnalyze_NilInput(t *testing.T) {
	analysis := Analyze(nil)
	if analysis != nil {
		t.Errorf("expected nil for nil input, got %+v", analysis)
	}
}

func TestAnalyze_AllKnownPatterns(t *testing.T) {
	tests := []struct {
		errorType    string
		expectedType string
		expectedConf string
	}{
		{"nil_pointer_dereference", "Nil Pointer Dereference", "high"},
		{"invalid_memory_address", "Invalid Memory Address", "high"},
		{"index_out_of_range", "Index Out of Range", "high"},
		{"runtime_error", "Runtime Error", "medium"},
		{"panic", "Panic", "low"},
	}

	for _, tc := range tests {
		parsed := &parse.ParsedError{
			RawMessage: "test message",
			ErrorType:  tc.errorType,
		}

		analysis := Analyze(parsed)
		if analysis == nil {
			t.Fatalf("expected analysis for %s, got nil", tc.errorType)
		}

		if analysis.ErrorType != tc.expectedType {
			t.Errorf("for %s: expected error type '%s', got '%s'",
				tc.errorType, tc.expectedType, analysis.ErrorType)
		}

		if analysis.Confidence != tc.expectedConf {
			t.Errorf("for %s: expected confidence '%s', got '%s'",
				tc.errorType, tc.expectedConf, analysis.Confidence)
		}
	}
}
