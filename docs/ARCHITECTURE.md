# Architecture

This document describes the internal architecture of vexErrScope.

## Overview

vexErrScope follows a pipeline architecture with four main stages:

```
Input → Parse → Analyze → Explain → Render → Output
```

## Components

### 1. Parser (`internal/parse`)

The parser extracts structured data from raw log input.

**Responsibilities:**
- Read input line by line
- Detect panic signatures using regex patterns
- Extract error messages
- Parse stack traces to identify the first stack frame

**Key Types:**
```go
type ParsedError struct {
    RawMessage string      // Original error message
    ErrorType  string      // Classified error type
    StackFrame *StackFrame // First stack frame
}

type StackFrame struct {
    Function string // Function name (e.g., "main.processUser")
    File     string // Source file path
    Line     string // Line number
}
```

### 2. Analyzer (`internal/analyze`)

The analyzer maps parsed errors to known patterns and generates explanations.

**Responsibilities:**
- Match error types to predefined patterns
- Provide root cause descriptions
- Suggest fixes
- Assign confidence scores

**Pattern Structure:**
```go
type Pattern struct {
    ErrorType  string // Human-readable error name
    RootCause  string // Explanation of why this happens
    SuggestFix string // How to fix it
    Confidence string // high, medium, or low
}
```

### 3. Explainer (`internal/explain`)

The explainer orchestrates the pipeline, connecting parser and analyzer.

**Responsibilities:**
- Coordinate parsing and analysis
- Transform analysis results into explanation format
- Handle error propagation

### 4. Renderer (`internal/render`)

The renderer formats explanations for output.

**Responsibilities:**
- Format structured data as human-readable text
- Handle nil/empty cases gracefully

## Data Flow

```
┌─────────┐     ┌────────┐     ┌──────────┐     ┌─────────┐     ┌────────┐
│  Input  │────▶│ Parser │────▶│ Analyzer │────▶│Explainer│────▶│Renderer│
│(io.Reader)    │        │     │          │     │         │     │        │
└─────────┘     └────────┘     └──────────┘     └─────────┘     └────────┘
                    │               │                │               │
                    ▼               ▼                ▼               ▼
              ParsedError      Analysis        Explanation      Formatted
                                                                  Output
```

## Design Decisions

### Why a Pipeline?

- **Testability**: Each stage can be tested independently
- **Extensibility**: New patterns or output formats can be added without affecting other stages
- **Clarity**: Single responsibility for each component

### Why Internal Packages?

Using `internal/` prevents external packages from importing implementation details, ensuring a stable public API through `cmd/`.

### Why Regex for Parsing?

- Go panic formats are well-defined and consistent
- Regex provides sufficient power without additional dependencies
- Easy to extend with new patterns

## Future Considerations

### Multi-language Support

The architecture supports adding new languages by:
1. Adding language-specific patterns to the analyzer
2. Extending the parser to detect language signatures
3. No changes needed to explainer or renderer

### Output Formats

Additional output formats (JSON, YAML) can be added by:
1. Creating new renderer implementations
2. Adding format selection to CLI
