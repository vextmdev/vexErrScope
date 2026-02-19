package parse

import (
	"strings"
	"testing"
)

func TestParse_NilPointerDereference(t *testing.T) {
	input := `panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x1234567]

goroutine 1 [running]:
main.processData(0x0)
	/home/user/project/main.go:42 +0x26
main.main()
	/home/user/project/main.go:15 +0x1a
`

	parsed, err := Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if parsed == nil {
		t.Fatal("expected parsed result, got nil")
	}

	if parsed.ErrorType != "nil_pointer_dereference" {
		t.Errorf("expected error type 'nil_pointer_dereference', got '%s'", parsed.ErrorType)
	}

	if !strings.Contains(parsed.RawMessage, "nil pointer dereference") {
		t.Errorf("expected raw message to contain 'nil pointer dereference', got '%s'", parsed.RawMessage)
	}

	if parsed.StackFrame == nil {
		t.Fatal("expected stack frame, got nil")
	}

	if parsed.StackFrame.Function != "main.processData" {
		t.Errorf("expected function 'main.processData', got '%s'", parsed.StackFrame.Function)
	}

	if parsed.StackFrame.File != "/home/user/project/main.go" {
		t.Errorf("expected file '/home/user/project/main.go', got '%s'", parsed.StackFrame.File)
	}

	if parsed.StackFrame.Line != "42" {
		t.Errorf("expected line '42', got '%s'", parsed.StackFrame.Line)
	}
}

func TestParse_IndexOutOfRange(t *testing.T) {
	input := `panic: runtime error: index out of range [5] with length 3

goroutine 1 [running]:
main.getElement(...)
	/app/handlers.go:89 +0x45
main.main()
	/app/main.go:20 +0x32
`

	parsed, err := Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if parsed == nil {
		t.Fatal("expected parsed result, got nil")
	}

	if parsed.ErrorType != "index_out_of_range" {
		t.Errorf("expected error type 'index_out_of_range', got '%s'", parsed.ErrorType)
	}

	if parsed.StackFrame == nil {
		t.Fatal("expected stack frame, got nil")
	}

	if parsed.StackFrame.Function != "main.getElement" {
		t.Errorf("expected function 'main.getElement', got '%s'", parsed.StackFrame.Function)
	}

	if parsed.StackFrame.Line != "89" {
		t.Errorf("expected line '89', got '%s'", parsed.StackFrame.Line)
	}
}

func TestParse_RuntimeError(t *testing.T) {
	input := `panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.divide(0x5, 0x0)
	/src/calc.go:15 +0x8b
`

	parsed, err := Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if parsed == nil {
		t.Fatal("expected parsed result, got nil")
	}

	if parsed.ErrorType != "runtime_error" {
		t.Errorf("expected error type 'runtime_error', got '%s'", parsed.ErrorType)
	}
}

func TestParse_NoPanic(t *testing.T) {
	input := `2024-01-15 10:30:45 INFO Starting server on port 8080
2024-01-15 10:30:46 INFO Server ready
`

	parsed, err := Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if parsed != nil {
		t.Errorf("expected nil result for non-panic input, got %+v", parsed)
	}
}

func TestParse_EmptyInput(t *testing.T) {
	input := ""

	parsed, err := Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if parsed != nil {
		t.Errorf("expected nil result for empty input, got %+v", parsed)
	}
}

func TestClassifyError(t *testing.T) {
	tests := []struct {
		message  string
		expected string
	}{
		{"invalid memory address or nil pointer dereference", "nil_pointer_dereference"},
		{"runtime error: nil pointer dereference", "nil_pointer_dereference"},
		{"invalid memory address", "invalid_memory_address"},
		{"index out of range [5] with length 3", "index_out_of_range"},
		{"runtime error: integer divide by zero", "runtime_error"},
		{"something unexpected happened", "panic"},
	}

	for _, tc := range tests {
		result := classifyError(tc.message)
		if result != tc.expected {
			t.Errorf("classifyError(%q) = %q, want %q", tc.message, result, tc.expected)
		}
	}
}
