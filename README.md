# Go Reloaded

A Go program that processes text files applying various transformations:
- Convert hexadecimal and binary numbers to decimal
- Apply capitalization transformations (uppercase, lowercase, etc.)
- Fix articles (a/an)
- Format punctuation
- Align quotes

## Installation

Requires Go 1.24.3 or later.

Clone the repository and run:

```bash
go build
```

## Usage

```bash
go run cmd/main.go <input_file>
```

The program will read the input file, process it, and output the result to `result.txt`.

## Example

Given a sample.txt with content:
```
harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '
```

After processing, result.txt will contain:
```
Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'
```

## Running Tests

```bash
go test ./...
```

## Project Structure

- `cmd/main.go`: Entry point of the application
- `process/process.go`: Main processing logic
- `pkg/`: Individual transformation packages (alignquotes, article, baseConversion, capitalise, punctuation)
- `tests/`: Test files