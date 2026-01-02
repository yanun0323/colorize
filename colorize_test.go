package colorize

import (
	"bytes"
	"strings"
	"testing"
)

func TestColorize(t *testing.T) {
	type args struct {
		s string
		c Color
	}

	tests := []struct {
		args args
		want string
	}{
		// Basic colors
		{
			args{s: "Black color", c: ColorBlack},
			"\x1b[30mBlack color\x1b[0m",
		},
		{
			args{s: "Red color", c: ColorRed},
			"\x1b[31mRed color\x1b[0m",
		},
		{
			args{s: "Green color", c: ColorGreen},
			"\x1b[32mGreen color\x1b[0m",
		},
		{
			args{s: "Yellow color", c: ColorYellow},
			"\x1b[33mYellow color\x1b[0m",
		},
		{
			args{s: "Blue color", c: ColorBlue},
			"\x1b[34mBlue color\x1b[0m",
		},
		{
			args{s: "Magenta color", c: ColorMagenta},
			"\x1b[35mMagenta color\x1b[0m",
		},
		{
			args{s: "Cyan color", c: ColorCyan},
			"\x1b[36mCyan color\x1b[0m",
		},
		{
			args{s: "White color", c: ColorWhite},
			"\x1b[37mWhite color\x1b[0m",
		},
		// Reversed colors
		{
			args{s: "BlackReversed color", c: ColorBlackReversed},
			"\x1b[40mBlackReversed color\x1b[0m",
		},
		{
			args{s: "RedReversed color", c: ColorRedReversed},
			"\x1b[41mRedReversed color\x1b[0m",
		},
		{
			args{s: "GreenReversed color", c: ColorGreenReversed},
			"\x1b[42mGreenReversed color\x1b[0m",
		},
		{
			args{s: "YellowReversed color", c: ColorYellowReversed},
			"\x1b[43mYellowReversed color\x1b[0m",
		},
		{
			args{s: "BlueReversed color", c: ColorBlueReversed},
			"\x1b[44mBlueReversed color\x1b[0m",
		},
		{
			args{s: "MagentaReversed color", c: ColorMagentaReversed},
			"\x1b[45mMagentaReversed color\x1b[0m",
		},
		{
			args{s: "CyanReversed color", c: ColorCyanReversed},
			"\x1b[46mCyanReversed color\x1b[0m",
		},
		{
			args{s: "WhiteReversed color", c: ColorWhiteReversed},
			"\x1b[47mWhiteReversed color\x1b[0m",
		},
		// Bright colors
		{
			args{s: "BrightBlack color", c: ColorBrightBlack},
			"\x1b[90mBrightBlack color\x1b[0m",
		},
		{
			args{s: "BrightRed color", c: ColorBrightRed},
			"\x1b[91mBrightRed color\x1b[0m",
		},
		{
			args{s: "BrightGreen color", c: ColorBrightGreen},
			"\x1b[92mBrightGreen color\x1b[0m",
		},
		{
			args{s: "BrightYellow color", c: ColorBrightYellow},
			"\x1b[93mBrightYellow color\x1b[0m",
		},
		{
			args{s: "BrightBlue color", c: ColorBrightBlue},
			"\x1b[94mBrightBlue color\x1b[0m",
		},
		{
			args{s: "BrightMagenta color", c: ColorBrightMagenta},
			"\x1b[95mBrightMagenta color\x1b[0m",
		},
		{
			args{s: "BrightCyan color", c: ColorBrightCyan},
			"\x1b[96mBrightCyan color\x1b[0m",
		},
		{
			args{s: "BrightWhite color", c: ColorBrightWhite},
			"\x1b[97mBrightWhite color\x1b[0m",
		},
		// Bright reverse colors
		{
			args{s: "BrightBlackReversed color", c: ColorBrightBlackReversed},
			"\x1b[100mBrightBlackReversed color\x1b[0m",
		},
		{
			args{s: "BrightRedReversed color", c: ColorBrightRedReversed},
			"\x1b[101mBrightRedReversed color\x1b[0m",
		},
		{
			args{s: "BrightGreenReversed color", c: ColorBrightGreenReversed},
			"\x1b[102mBrightGreenReversed color\x1b[0m",
		},
		{
			args{s: "BrightYellowReversed color", c: ColorBrightYellowReversed},
			"\x1b[103mBrightYellowReversed color\x1b[0m",
		},
		{
			args{s: "BrightBlueReversed color", c: ColorBrightBlueReversed},
			"\x1b[104mBrightBlueReversed color\x1b[0m",
		},
		{
			args{s: "BrightMagentaReversed color", c: ColorBrightMagentaReversed},
			"\x1b[105mBrightMagentaReversed color\x1b[0m",
		},
		{
			args{s: "BrightCyanReversed color", c: ColorBrightCyanReversed},
			"\x1b[106mBrightCyanReversed color\x1b[0m",
		},
		{
			args{s: "BrightWhiteReversed color", c: ColorBrightWhiteReversed},
			"\x1b[107mBrightWhiteReversed color\x1b[0m",
		},
		// Test with empty string
		{
			args{s: "", c: ColorRed},
			"\x1b[31m\x1b[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.args.s, func(t *testing.T) {
			got := Sprint(tt.args.c, tt.args.s)
			if got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}

			got = Sprintf(tt.args.c, "%s", tt.args.s)
			if string(got) != tt.want {
				t.Errorf("Bytes() = %q, want %q", string(got), tt.want)
			}
		})
	}
}

func TestColorizeMulti(t *testing.T) {
	buf := strings.Builder{}
	Fprint(&buf, ColorRed, "Hello, World")
	if buf.String() != "\x1b[31mHello, World\x1b[0m" {
		t.Errorf("WriteString() = %q, want %q", buf.String(), "\x1b[31mHello, World\x1b[0m")
	}

	Fprintf(&buf, ColorGreen, "%s", "Hello, World")
	if buf.String() != "\x1b[31mHello, World\x1b[0m\x1b[32mHello, World\x1b[0m" {
		t.Errorf("WriteBytes() = %q, want %q", buf.String(), "\x1b[31mHello, World\x1b[0m\x1b[32mHello, World\x1b[0m")
	}
}

func TestColorizeReset(t *testing.T) {
	buf := bytes.Buffer{}
	Fprint(&buf, ColorRed, "Hello, World")
	if Reset(buf.String()) != "Hello, World" {
		t.Errorf("ResetString() = %q, want %q", Reset(buf.String()), "Hello, World")
	}
}

func TestColorizePrint(t *testing.T) {
	type args struct {
		s string
		c Color
	}
	tests := []args{

		// Basic colors
		{},
		{s: " black ", c: ColorBlack},
		{s: " red ", c: ColorRed},
		{s: " green ", c: ColorGreen},
		{s: " yellow ", c: ColorYellow},
		{s: " blue ", c: ColorBlue},
		{s: " magenta ", c: ColorMagenta},
		{s: " cyan ", c: ColorCyan},
		{s: " white ", c: ColorWhite},
		// Reversed colors
		{},
		{s: " black reversed ", c: ColorBlackReversed},
		{s: " red reversed ", c: ColorRedReversed},
		{s: " green reversed ", c: ColorGreenReversed},
		{s: " yellow reversed ", c: ColorYellowReversed},
		{s: " blue reversed ", c: ColorBlueReversed},
		{s: " magenta reversed ", c: ColorMagentaReversed},
		{s: " cyan reversed ", c: ColorCyanReversed},
		{s: " white reversed ", c: ColorWhiteReversed},
		// Bright colors
		{},
		{s: " bright black ", c: ColorBrightBlack},
		{s: " bright red ", c: ColorBrightRed},
		{s: " bright green ", c: ColorBrightGreen},
		{s: " bright yellow ", c: ColorBrightYellow},
		{s: " bright blue ", c: ColorBrightBlue},
		{s: " bright magenta ", c: ColorBrightMagenta},
		{s: " bright cyan ", c: ColorBrightCyan},
		{s: " bright white ", c: ColorBrightWhite},
		// Bright reverse colors
		{},
		{s: " bright black reversed ", c: ColorBrightBlackReversed},
		{s: " bright red reversed ", c: ColorBrightRedReversed},
		{s: " bright green reversed ", c: ColorBrightGreenReversed},
		{s: " bright yellow reversed ", c: ColorBrightYellowReversed},
		{s: " bright blue reversed ", c: ColorBrightBlueReversed},
		{s: " bright magenta reversed ", c: ColorBrightMagentaReversed},
		{s: " bright cyan reversed ", c: ColorBrightCyanReversed},
		{s: " bright white reversed ", c: ColorBrightWhiteReversed},

		{},
	}

	println("=== Sprint ===")
	for _, tt := range tests {
		println(Sprint(tt.c, tt.s))
	}

	println("=== Reset ===")
	for _, tt := range tests {
		println(Reset(Sprint(tt.c, tt.s)))
	}
}

func BenchmarkColorizeString(b *testing.B) {
	for b.Loop() {
		String(ColorRed, "Hello, World")
	}
}

func BenchmarkColorizeReset(b *testing.B) {
	s := String(ColorRed, "Hello, World")
	for b.Loop() {
		Reset(s)
	}
}

func BenchmarkColorizeSprint(b *testing.B) {
	for b.Loop() {
		Sprint(ColorRed, "Hello, World")
	}
}

func BenchmarkColorizeSprintf(b *testing.B) {
	for b.Loop() {
		Sprintf(ColorRed, "%s", "Hello, World")
	}
}

func BenchmarkColorizeFprint(b *testing.B) {
	buf := bytes.Buffer{}
	for b.Loop() {
		buf.Reset()
		Fprint(&buf, ColorRed, "Hello, World")
	}
}

func BenchmarkColorizeFprintf(b *testing.B) {
	buf := bytes.Buffer{}
	for b.Loop() {
		buf.Reset()
		Fprintf(&buf, ColorRed, "%s", "Hello, World")
	}
}

func BenchmarkColorizeAppend(b *testing.B) {
	buf := make([]byte, 4<<10)
	for b.Loop() {
		buf = buf[:0]
		buf = append(buf, ColorRed...)
		buf = append(buf, "Hello, World"...)
		buf = append(buf, ColorReset...)
	}
}
