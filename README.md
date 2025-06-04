# Colorize

A simple, lightweight Go library for terminal text colorization using ANSI escape codes.

<img src="https://raw.githubusercontent.com/yanun0323/assets/refs/heads/master/colorize.split.four.png" alt="colorize" style="border-radius: 8px;" />

## Features

- Simple and intuitive API
- Support for basic colors, bright colors, and background colors
- Memory efficient with buffer pooling
- Zero external dependencies

## Installation

```bash
go get github.com/yanun0323/colorize
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/yanun0323/colorize"
)

func main() {
    fmt.Println(colorize.String("Hello World", colorize.Red))
    fmt.Println(colorize.String("Success!", colorize.Green))
    fmt.Println(colorize.String("Warning", colorize.Yellow))
    fmt.Println(colorize.String("Error", colorize.BrightRed))
}
```

## Available Colors

### Basic Colors

`Black`, `Red`, `Green`, `Yellow`, `Blue`, `Magenta`, `Cyan`, `White`

### Bright Colors

`BrightBlack`, `BrightRed`, `BrightGreen`, `BrightYellow`, `BrightBlue`, `BrightMagenta`, `BrightCyan`, `BrightWhite`

### Background Colors

`BlackReversed`, `RedReversed`, `GreenReversed`, `YellowReversed`, `BlueReversed`, `MagentaReversed`, `CyanReversed`, `WhiteReversed`

### Bright Background Colors

`BrightBlackReversed`, `BrightRedReversed`, `BrightGreenReversed`, `BrightYellowReversed`, `BrightBlueReversed`, `BrightMagentaReversed`, `BrightCyanReversed`, `BrightWhiteReversed`

## API

### `Colorize(s string, c string) string`

Colorizes a string with the specified color.

```go
colored := colorize.Colorize("Hello", colorize.Blue)
fmt.Println(colored)
```

### `ColorizeToBuffer(buf io.Writer, s []byte, c []byte)`

Writes colorized string directly to a buffer for better performance.

```go
var buf bytes.Buffer
colorize.ColorizeToBuffer(&buf, []byte("Hello"), []byte(colorize.Green))
fmt.Print(buf.String())
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
