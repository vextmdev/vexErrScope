# Error Pattern Rules

This document describes the error patterns recognized by vexErrScope.

## Go Patterns

### Nil Pointer Dereference

**Pattern:** `nil pointer dereference`

**Example:**
```
panic: runtime error: invalid memory address or nil pointer dereference
```

**Root Cause:**
Attempted to access a field or method on a nil pointer. This typically happens when a pointer variable was never initialized or was explicitly set to nil.

**Suggested Fix:**
- Add nil checks before dereferencing pointers
- Ensure pointers are properly initialized before use
- Use the comma-ok idiom for map lookups or type assertions

**Confidence:** High

---

### Invalid Memory Address

**Pattern:** `invalid memory address`

**Example:**
```
panic: runtime error: invalid memory address or nil pointer dereference
```

**Root Cause:**
Program attempted to access memory at an invalid address, typically due to a nil pointer or corrupted pointer value.

**Suggested Fix:**
- Check for nil pointers before dereferencing
- Verify that all pointers are properly initialized
- Check for race conditions in concurrent code

**Confidence:** High

---

### Index Out of Range

**Pattern:** `index out of range`

**Example:**
```
panic: runtime error: index out of range [5] with length 3
```

**Root Cause:**
Attempted to access a slice or array element at an index that exceeds its bounds. The index is either negative or greater than or equal to the length.

**Suggested Fix:**
- Add bounds checking before accessing elements
- Use `len()` to verify slice has enough elements
- Consider using range loops instead of index-based access

**Confidence:** High

---

### Runtime Error

**Pattern:** `runtime error`

**Example:**
```
panic: runtime error: integer divide by zero
```

**Root Cause:**
A general runtime error occurred during program execution.

**Suggested Fix:**
- Review the specific error message for details
- Common causes: nil pointers, out-of-bounds access, division by zero

**Confidence:** Medium

---

### Generic Panic

**Pattern:** `panic:` (fallback)

**Example:**
```
panic: something went wrong
```

**Root Cause:**
The program encountered an unrecoverable error and panicked.

**Suggested Fix:**
- Review the panic message and stack trace
- Consider adding `recover()` in critical sections

**Confidence:** Low

---

## Adding New Rules

To add a new error pattern:

1. **Parser** (`internal/parse/parser.go`):
   - Add pattern detection in `classifyError()` function

2. **Analyzer** (`internal/analyze/analyzer.go`):
   - Add pattern to `goPatterns` map:

```go
"your_pattern_id": {
    ErrorType:  "Human Readable Name",
    RootCause:  "Description of why this error occurs.",
    SuggestFix: "Steps to fix this error.",
    Confidence: "high", // or "medium", "low"
},
```

3. **Tests**:
   - Add test case in `parser_test.go`
   - Add test case in `analyzer_test.go`

4. **Documentation**:
   - Document the pattern in this file

## Confidence Levels

| Level | Meaning |
|-------|---------|
| High | Pattern is well-defined and uniquely identifiable |
| Medium | Pattern may have multiple causes |
| Low | Fallback pattern or insufficient information |
