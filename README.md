# Colorize

A lightweight Go library for terminal text colorization using ANSI escape codes.

## Features

- Simple and intuitive API
- Support for multiple color schemes (basic, bright, background colors)
- Memory efficient with sync.Pool buffer management
- High-performance operations with zero external dependencies
- Compatible with any io.Writer interface

## Installation

```bash
go get github.com/yanun0323/colorize
```

## Quick Start

```go
package main

import (
    "fmt"
    "strings"

    "github.com/yanun0323/colorize"
)

func main() {
    // Basic string colorization
    fmt.Println(colorize.String(colorize.ColorRed, "Error message"))
    fmt.Println(colorize.Sprint(colorize.ColorGreen, "Success:", " operation completed"))
    fmt.Println(colorize.Sprintf(colorize.ColorBlue, "User %s logged in", "john"))

    // Write to buffer
    var buf strings.Builder
    colorize.Fprint(&buf, colorize.ColorYellow, "Warning:", " low disk space")
    colorize.Fprintf(&buf, colorize.ColorCyan, " Available: %d%%", 15)
    fmt.Print(buf.String())

    // Remove ANSI codes
    coloredText := colorize.String(colorize.ColorRed, "Error")
    plainText := colorize.Reset(coloredText)
    fmt.Println(plainText) // Output: "Error" (without colors)
}
```

## API

### Core Functions

- `String(c Color, str ...string) string` - Colorize multiple strings
- `Sprint(c Color, args ...any) string` - Colorize with fmt.Sprint behavior
- `Sprintf(c Color, format string, args ...any) string` - Colorize with formatting
- `Fprint(w Writer, c Color, args ...any) (int, error)` - Write colorized text to writer
- `Fprintf(w Writer, c Color, format string, args ...any) (int, error)` - Write formatted colorized text
- `Reset(s string) string` - Remove ANSI escape codes from string

### Available Colors

<img src="https://raw.githubusercontent.com/yanun0323/assets/refs/heads/master/colorize.split.four.png" alt="colorize" style="border-radius: 8px;" />

## Benchmarks

Benchmarks results on a 2022 MacBook Air 15" with an Apple M2 chip.

```bash
BenchmarkColorizeString-8       39335316                30.01 ns/op           24 B/op          1 allocs/op
BenchmarkColorizeSprint-8       23177600                51.25 ns/op           24 B/op          1 allocs/op
BenchmarkColorizeSprintf-8      22460238                53.26 ns/op           24 B/op          1 allocs/op
BenchmarkColorizeFprint-8       36809578                33.07 ns/op            0 B/op          0 allocs/op
BenchmarkColorizeFprintf-8      34705401                35.12 ns/op            0 B/op          0 allocs/op
BenchmarkColorizeReset-8        25729447                45.86 ns/op           16 B/op          1 allocs/op
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
