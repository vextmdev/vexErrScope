package parse

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

// ParsedError holds the extracted error information from a log or stack trace.
type ParsedError struct {
	RawMessage string
	ErrorType  string
	StackFrame *StackFrame
}

// StackFrame represents a single frame in a stack trace.
type StackFrame struct {
	Function string
	File     string
	Line     string
}

var (
	panicRegex    = regexp.MustCompile(`(?i)panic:\s*(.+)`)
	funcRegex     = regexp.MustCompile(`^(\S+)\(.*\)$`)
	fileLineRegex = regexp.MustCompile(`^\s*(.+\.go):(\d+)`)
)

// Parse reads input and extracts error information.
func Parse(r io.Reader) (*ParsedError, error) {
	scanner := bufio.NewScanner(r)
	result := &ParsedError{}

	var foundPanic bool
	var lookingForStack bool

	for scanner.Scan() {
		line := scanner.Text()

		if matches := panicRegex.FindStringSubmatch(line); matches != nil {
			result.RawMessage = strings.TrimSpace(matches[1])
			result.ErrorType = classifyError(result.RawMessage)
			foundPanic = true
			lookingForStack = true
			continue
		}

		if lookingForStack {
			if strings.HasPrefix(line, "goroutine ") {
				continue
			}

			if result.StackFrame != nil && result.StackFrame.File == "" {
				if fileMatches := fileLineRegex.FindStringSubmatch(line); fileMatches != nil {
					result.StackFrame.File = fileMatches[1]
					result.StackFrame.Line = fileMatches[2]
					break
				}
			}

			if result.StackFrame == nil {
				trimmed := strings.TrimSpace(line)
				if funcMatches := funcRegex.FindStringSubmatch(trimmed); funcMatches != nil {
					result.StackFrame = &StackFrame{
						Function: funcMatches[1],
					}
					continue
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if !foundPanic {
		return nil, nil
	}

	return result, nil
}

// classifyError determines the error type from the message.
func classifyError(msg string) string {
	lower := strings.ToLower(msg)

	switch {
	case strings.Contains(lower, "nil pointer dereference"):
		return "nil_pointer_dereference"
	case strings.Contains(lower, "invalid memory address"):
		return "invalid_memory_address"
	case strings.Contains(lower, "index out of range"):
		return "index_out_of_range"
	case strings.Contains(lower, "runtime error"):
		return "runtime_error"
	default:
		return "panic"
	}
}
