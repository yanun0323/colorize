package colorize

import (
	"bytes"
	"fmt"
	"io"
)

type Color string

func (c Color) String() string {
	return string(c)
}

const (
	ColorReset Color = "\x1b[0m"
)

const (
	ColorBlack   Color = "\x1b[30m"
	ColorRed     Color = "\x1b[31m"
	ColorGreen   Color = "\x1b[32m"
	ColorYellow  Color = "\x1b[33m"
	ColorBlue    Color = "\x1b[34m"
	ColorMagenta Color = "\x1b[35m"
	ColorCyan    Color = "\x1b[36m"
	ColorWhite   Color = "\x1b[37m"

	ColorBlackReversed   Color = "\x1b[40m"
	ColorRedReversed     Color = "\x1b[41m"
	ColorGreenReversed   Color = "\x1b[42m"
	ColorYellowReversed  Color = "\x1b[43m"
	ColorBlueReversed    Color = "\x1b[44m"
	ColorMagentaReversed Color = "\x1b[45m"
	ColorCyanReversed    Color = "\x1b[46m"
	ColorWhiteReversed   Color = "\x1b[47m"

	ColorBrightBlack   Color = "\x1b[90m"
	ColorBrightRed     Color = "\x1b[91m"
	ColorBrightGreen   Color = "\x1b[92m"
	ColorBrightYellow  Color = "\x1b[93m"
	ColorBrightBlue    Color = "\x1b[94m"
	ColorBrightMagenta Color = "\x1b[95m"
	ColorBrightCyan    Color = "\x1b[96m"
	ColorBrightWhite   Color = "\x1b[97m"

	ColorBrightBlackReversed   Color = "\x1b[100m"
	ColorBrightRedReversed     Color = "\x1b[101m"
	ColorBrightGreenReversed   Color = "\x1b[102m"
	ColorBrightYellowReversed  Color = "\x1b[103m"
	ColorBrightBlueReversed    Color = "\x1b[104m"
	ColorBrightMagentaReversed Color = "\x1b[105m"
	ColorBrightCyanReversed    Color = "\x1b[106m"
	ColorBrightWhiteReversed   Color = "\x1b[107m"
)

const (
	_defaultBufferSize = 4 << 10
)

// String colorize string
func String(c Color, str ...string) string {
	if len(str) == 0 {
		return ""
	}

	sz := len(c) + len(ColorReset)
	for i := range str {
		sz += len(str[i])
	}

	buf := make([]byte, sz)
	buf = buf[:0]
	buf = append(buf, c...)
	for i := range str {
		buf = append(buf, str[i]...)
	}
	buf = append(buf, ColorReset...)
	return string(buf)
}

// Sprint colorize and formats using the default formats for its operands and returns the resulting string. Spaces are added between operands when neither is a string.
func Sprint(c Color, args ...any) string {
	buf := bytes.Buffer{}
	Fprint(&buf, c, args...)
	return buf.String()
}

// Sprintf colorize and formats according to a format specifier and returns the resulting string.
func Sprintf(c Color, format string, args ...any) string {
	buf := bytes.Buffer{}
	Fprintf(&buf, c, format, args...)
	return buf.String()
}

// Fprint colorize and formats using the default formats for its operands and writes to w. Spaces are added between operands when neither is a string. It returns the number of bytes written and any write error encountered.
func Fprint(w io.Writer, c Color, args ...any) (int, error) {
	if len(args) == 0 {
		return 0, nil
	}

	var (
		result int
		n      int
		err    error
	)
	n, err = io.WriteString(w, string(c))
	if err != nil {
		return 0, err
	}
	result += n

	n, err = fmt.Fprint(w, args...)
	if err != nil {
		return 0, err
	}
	result += n

	n, err = io.WriteString(w, string(ColorReset))
	if err != nil {
		return 0, err
	}
	result += n

	return result, nil
}

// Fprintf colorize and formats using the default formats for its operands and writes to w. Spaces are added between operands when neither is a string. It returns the number of bytes written and any write error encountered.
func Fprintf(w io.Writer, c Color, format string, args ...any) (int, error) {
	if len(format) == 0 {
		return 0, nil
	}

	if len(args) == 0 {
		return Fprint(w, c, format)
	}

	var (
		result int
		n      int
		err    error
	)
	n, err = io.WriteString(w, string(c))
	if err != nil {
		return 0, err
	}
	result += n

	n, err = fmt.Fprintf(w, format, args...)
	if err != nil {
		return 0, err
	}
	result += n

	n, err = io.WriteString(w, string(ColorReset))
	if err != nil {
		return 0, err
	}
	result += n

	return result, nil
}

// Reset reset string color
func Reset(s string) string {
	if len(s) == 0 {
		return s
	}

	buf := make([]byte, 0, len(s))
	i := 0
	for i < len(s) {
		if s[i] == '\x1b' && i+1 < len(s) && s[i+1] == '[' {
			// Find the end of ANSI escape sequence
			j := i + 2
			for j < len(s) && s[j] != 'm' {
				j++
			}
			if j < len(s) {
				// Skip the entire escape sequence including 'm'
				i = j + 1
			} else {
				// Malformed escape sequence, keep the character
				buf = append(buf, s[i])
				i++
			}
		} else {
			buf = append(buf, s[i])
			i++
		}
	}

	return string(buf)
}
