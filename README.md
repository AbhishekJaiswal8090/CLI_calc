# CLI Calculator

A command-line calculator written in Go. This program allows users to perform basic arithmetic operations using command-line flags.

## Features
- Addition, subtraction, multiplication, and division
- Command-line interface with flags for input
- Written in Go

## Usage

1. Build the program:
   ```sh
   go build -o calc CLI_calc/calc.go
   ```

2. Run the calculator using command-line flags:
   ```sh
   go run CLI_calc/calc.go -a=10 -b=40 -op=mul
   ```
   or if built:
   ```sh
   ./calc -a=10 -b=40 -op=mul
   ```

   - `-a` : First number
   - `-b` : Second number
   - `-op` : Operation (`add`, `sub`, `mul`, `div`)

## Example

```
go run CLI_calc/calc.go -a=10 -b=40 -op=mul
Result: 400
```

## Requirements
- Go 1.16 or higher


