# vexErrScope

A CLI tool that scans logs and stack traces, detects common error patterns, and outputs structured, human-readable explanations.

**Current Version:** v0.1.0 (initial release)

## Features

- Detects Go panic patterns:
  - Nil pointer dereference
  - Invalid memory address
  - Index out of range
  - General runtime errors
- Extracts stack trace information (function, file, line number)
- Provides root cause explanations and suggested fixes
- Confidence scoring (high/medium/low)

## Requirements

- Go 1.21 or later

## Installation

Build from source:

```bash
# Navigate to the project directory
cd vexErrScope

# Build the binary
go build -o vexErrScope ./cmd/vexErrScope

# (Optional) Move to your PATH
mv vexErrScope /usr/local/bin/
```

## Quick Start

```bash
# Analyze a log file
./vexErrScope explain path/to/crash.log

# Read from stdin
cat error.log | ./vexErrScope explain

# Show help
./vexErrScope help
```

---

## Tutorial

### Step 1: Create a Test Error Log

Create a file called `test-panic.log` with a Go panic:

```bash
cat > test-panic.log << 'EOF'
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x1092a40]

goroutine 1 [running]:
main.processUser(0x0)
	/home/dev/myapp/handlers/user.go:42 +0x26
main.main()
	/home/dev/myapp/main.go:15 +0x1a
EOF
```

### Step 2: Analyze the Error

Run vexErrScope on the file:

```bash
./vexErrScope explain test-panic.log
```

### Step 3: Read the Output

You'll see structured output like this:

```
Error Type:        Nil Pointer Dereference
Detected Language: Go
Root Cause:        Attempted to access a field or method on a nil pointer. This typically happens when a pointer variable was never initialized or was explicitly set to nil.
Suggested Fix:     Add nil checks before dereferencing pointers. Ensure the pointer is properly initialized before use. Consider using the comma-ok idiom for map lookups or type assertions.
Confidence:        high

Location:
  Function: main.processUser
  File:     /home/dev/myapp/handlers/user.go
  Line:     42
```

### Step 4: Using with Pipes

You can pipe errors directly from other commands:

```bash
# Pipe from a failed go run
go run main.go 2>&1 | ./vexErrScope explain

# Pipe from docker logs
docker logs mycontainer 2>&1 | ./vexErrScope explain

# Pipe from a log file with grep
grep -A 20 "panic:" application.log | ./vexErrScope explain
```

### Step 5: Try Different Error Types

**Index out of range:**

```bash
echo 'panic: runtime error: index out of range [5] with length 3

goroutine 1 [running]:
main.getItem(0x5)
	/app/main.go:25 +0x45' | ./vexErrScope explain
```

**Generic panic:**

```bash
echo 'panic: connection refused

goroutine 1 [running]:
main.connect()
	/app/db.go:100 +0x80' | ./vexErrScope explain
```

---

## Supported Patterns

| Pattern | Detection | Confidence |
|---------|-----------|------------|
| `nil pointer dereference` | Yes | High |
| `invalid memory address` | Yes | High |
| `index out of range` | Yes | High |
| `runtime error` (generic) | Yes | Medium |
| Any `panic:` | Yes | Low |

## Project Structure

```
vexErrScope/
├── cmd/vexErrScope/
│   └── main.go           # CLI entry point
├── internal/
│   ├── parse/            # Parses error messages and stack traces
│   ├── analyze/          # Maps patterns to explanations
│   ├── explain/          # Orchestrates parse → analyze
│   └── render/           # Formats output
├── testdata/
│   └── sample.log        # Example panic for testing
└── docs/
    ├── ARCHITECTURE.md   # Internal design documentation
    └── RULES.md          # Error pattern definitions
```

## Running Tests

```bash
go test ./...
```

## Limitations

- Currently only supports Go panic patterns
- Outputs plain text only (no JSON/YAML yet)
- Detects first panic in input only

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

MIT License - see [LICENSE](LICENSE) for details.

## Roadmap

Future improvements (not yet implemented):

- [ ] Python traceback support
- [ ] JavaScript error support
- [ ] JSON output format
- [ ] Multiple error detection in single input
- [ ] Custom pattern definitions via config file
