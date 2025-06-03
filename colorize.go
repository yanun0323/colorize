package colorize

import (
	"bytes"
	"sync"
)

const (
	Reset = "\x1b[0m"
)

const (
	Black   = "\x1b[30m"
	Red     = "\x1b[31m"
	Green   = "\x1b[32m"
	Yellow  = "\x1b[33m"
	Blue    = "\x1b[34m"
	Magenta = "\x1b[35m"
	Cyan    = "\x1b[36m"
	White   = "\x1b[37m"

	BlackReversed   = "\x1b[40m"
	RedReversed     = "\x1b[41m"
	GreenReversed   = "\x1b[42m"
	YellowReversed  = "\x1b[43m"
	BlueReversed    = "\x1b[44m"
	MagentaReversed = "\x1b[45m"
	CyanReversed    = "\x1b[46m"
	WhiteReversed   = "\x1b[47m"

	BrightBlack   = "\x1b[90m"
	BrightRed     = "\x1b[91m"
	BrightGreen   = "\x1b[92m"
	BrightYellow  = "\x1b[93m"
	BrightBlue    = "\x1b[94m"
	BrightMagenta = "\x1b[95m"
	BrightCyan    = "\x1b[96m"
	BrightWhite   = "\x1b[97m"

	BrightBlackReversed   = "\x1b[100m"
	BrightRedReversed     = "\x1b[101m"
	BrightGreenReversed   = "\x1b[102m"
	BrightYellowReversed  = "\x1b[103m"
	BrightBlueReversed    = "\x1b[104m"
	BrightMagentaReversed = "\x1b[105m"
	BrightCyanReversed    = "\x1b[106m"
	BrightWhiteReversed   = "\x1b[107m"
)

var (
	bufferPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, 1024))
		},
	}
)

// String colorize string
func String(s string, c string) string {
	buf := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buf)
	buf.Reset()
	buf.Grow(len(s) + len(c) + len(Reset))
	Write(buf, s, c)
	return buf.String()
}

// Bytes colorize bytes
func Bytes(s []byte, c string) []byte {
	buf := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buf)
	buf.Reset()
	buf.Grow(len(s) + len(c) + len(Reset))
	Write(buf, string(s), c)
	return buf.Bytes()
}

// StringWriter is the interface that wraps the basic WriteString method.
type StringWriter interface {
	WriteString(s string) (n int, err error)
}

// Write write colorized string to buffer
func Write(buf StringWriter, s string, c string) {
	buf.WriteString(c)
	buf.WriteString(s)
	buf.WriteString(Reset)
}
