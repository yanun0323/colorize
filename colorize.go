package colorize

import (
	"bytes"
	"io"
	"sync"
)

var (
	_colorPrefix = []byte("\x1b[")
	_colorReset  = []byte("\x1b[0m")
)

const (
	Black   = "30m"
	Red     = "31m"
	Green   = "32m"
	Yellow  = "33m"
	Blue    = "34m"
	Magenta = "35m"
	Cyan    = "36m"
	White   = "37m"

	BlackReversed   = "40m"
	RedReversed     = "41m"
	GreenReversed   = "42m"
	YellowReversed  = "43m"
	BlueReversed    = "44m"
	MagentaReversed = "45m"
	CyanReversed    = "46m"
	WhiteReversed   = "47m"

	BrightBlack   = "90m"
	BrightRed     = "91m"
	BrightGreen   = "92m"
	BrightYellow  = "93m"
	BrightBlue    = "94m"
	BrightMagenta = "95m"
	BrightCyan    = "96m"
	BrightWhite   = "97m"

	BrightBlackReversed   = "100m"
	BrightRedReversed     = "101m"
	BrightGreenReversed   = "102m"
	BrightYellowReversed  = "103m"
	BrightBlueReversed    = "104m"
	BrightMagentaReversed = "105m"
	BrightCyanReversed    = "106m"
	BrightWhiteReversed   = "107m"
)

var (
	bufferPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, 1024))
		},
	}
)

// Colorize colorize string
func Colorize(s string, c string) string {
	buf := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buf)
	buf.Reset()
	buf.Grow(len(s) + len(_colorPrefix) + len(c) + len(_colorReset))
	ColorizeToBuffer(buf, []byte(s), []byte(c))
	return buf.String()
}

// ColorizeToBuffer write colorized string to buffer
func ColorizeToBuffer(buf io.Writer, s []byte, c []byte) {
	buf.Write(_colorPrefix)
	buf.Write(c)
	buf.Write(s)
	buf.Write(_colorReset)
}
