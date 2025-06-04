package colorize

import (
	"bytes"
	"strings"
	"testing"
)

func TestColorize(t *testing.T) {
	type args struct {
		s string
		c string
	}

	tests := []struct {
		args args
		want string
	}{
		// Basic colors
		{
			args{s: "Black color", c: Black},
			"\x1b[30mBlack color\x1b[0m",
		},
		{
			args{s: "Red color", c: Red},
			"\x1b[31mRed color\x1b[0m",
		},
		{
			args{s: "Green color", c: Green},
			"\x1b[32mGreen color\x1b[0m",
		},
		{
			args{s: "Yellow color", c: Yellow},
			"\x1b[33mYellow color\x1b[0m",
		},
		{
			args{s: "Blue color", c: Blue},
			"\x1b[34mBlue color\x1b[0m",
		},
		{
			args{s: "Magenta color", c: Magenta},
			"\x1b[35mMagenta color\x1b[0m",
		},
		{
			args{s: "Cyan color", c: Cyan},
			"\x1b[36mCyan color\x1b[0m",
		},
		{
			args{s: "White color", c: White},
			"\x1b[37mWhite color\x1b[0m",
		},
		// Reversed colors
		{
			args{s: "BlackReversed color", c: BlackReversed},
			"\x1b[40mBlackReversed color\x1b[0m",
		},
		{
			args{s: "RedReversed color", c: RedReversed},
			"\x1b[41mRedReversed color\x1b[0m",
		},
		{
			args{s: "GreenReversed color", c: GreenReversed},
			"\x1b[42mGreenReversed color\x1b[0m",
		},
		{
			args{s: "YellowReversed color", c: YellowReversed},
			"\x1b[43mYellowReversed color\x1b[0m",
		},
		{
			args{s: "BlueReversed color", c: BlueReversed},
			"\x1b[44mBlueReversed color\x1b[0m",
		},
		{
			args{s: "MagentaReversed color", c: MagentaReversed},
			"\x1b[45mMagentaReversed color\x1b[0m",
		},
		{
			args{s: "CyanReversed color", c: CyanReversed},
			"\x1b[46mCyanReversed color\x1b[0m",
		},
		{
			args{s: "WhiteReversed color", c: WhiteReversed},
			"\x1b[47mWhiteReversed color\x1b[0m",
		},
		// Bright colors
		{
			args{s: "BrightBlack color", c: BrightBlack},
			"\x1b[90mBrightBlack color\x1b[0m",
		},
		{
			args{s: "BrightRed color", c: BrightRed},
			"\x1b[91mBrightRed color\x1b[0m",
		},
		{
			args{s: "BrightGreen color", c: BrightGreen},
			"\x1b[92mBrightGreen color\x1b[0m",
		},
		{
			args{s: "BrightYellow color", c: BrightYellow},
			"\x1b[93mBrightYellow color\x1b[0m",
		},
		{
			args{s: "BrightBlue color", c: BrightBlue},
			"\x1b[94mBrightBlue color\x1b[0m",
		},
		{
			args{s: "BrightMagenta color", c: BrightMagenta},
			"\x1b[95mBrightMagenta color\x1b[0m",
		},
		{
			args{s: "BrightCyan color", c: BrightCyan},
			"\x1b[96mBrightCyan color\x1b[0m",
		},
		{
			args{s: "BrightWhite color", c: BrightWhite},
			"\x1b[97mBrightWhite color\x1b[0m",
		},
		// Bright reverse colors
		{
			args{s: "BrightBlackReversed color", c: BrightBlackReversed},
			"\x1b[100mBrightBlackReversed color\x1b[0m",
		},
		{
			args{s: "BrightRedReversed color", c: BrightRedReversed},
			"\x1b[101mBrightRedReversed color\x1b[0m",
		},
		{
			args{s: "BrightGreenReversed color", c: BrightGreenReversed},
			"\x1b[102mBrightGreenReversed color\x1b[0m",
		},
		{
			args{s: "BrightYellowReversed color", c: BrightYellowReversed},
			"\x1b[103mBrightYellowReversed color\x1b[0m",
		},
		{
			args{s: "BrightBlueReversed color", c: BrightBlueReversed},
			"\x1b[104mBrightBlueReversed color\x1b[0m",
		},
		{
			args{s: "BrightMagentaReversed color", c: BrightMagentaReversed},
			"\x1b[105mBrightMagentaReversed color\x1b[0m",
		},
		{
			args{s: "BrightCyanReversed color", c: BrightCyanReversed},
			"\x1b[106mBrightCyanReversed color\x1b[0m",
		},
		{
			args{s: "BrightWhiteReversed color", c: BrightWhiteReversed},
			"\x1b[107mBrightWhiteReversed color\x1b[0m",
		},
		// Test with empty string
		{
			args{s: "", c: Red},
			"\x1b[31m\x1b[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.args.s, func(t *testing.T) {
			got := String(tt.args.c, tt.args.s)
			if got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}

			got = string(Bytes(tt.args.c, []byte(tt.args.s)))
			if string(got) != tt.want {
				t.Errorf("Bytes() = %q, want %q", string(got), tt.want)
			}
		})
	}
}

func TestColorizeMulti(t *testing.T) {
	buf := strings.Builder{}
	WriteString(&buf, Red, "Hello", ", World")
	if buf.String() != "\x1b[31mHello, World\x1b[0m" {
		t.Errorf("WriteString() = %q, want %q", buf.String(), "\x1b[31mHello, World\x1b[0m")
	}

	WriteBytes(&buf, Green, []byte("Hello"), []byte(", World"))
	if buf.String() != "\x1b[31mHello, World\x1b[0m\x1b[32mHello, World\x1b[0m" {
		t.Errorf("WriteBytes() = %q, want %q", buf.String(), "\x1b[31mHello, World\x1b[0m\x1b[32mHello, World\x1b[0m")
	}
}

func TestColorizeReset(t *testing.T) {
	buf := bytes.Buffer{}
	WriteString(&buf, Red, "Hello", ", World")
	if ResetString(buf.String()) != "Hello, World" {
		t.Errorf("ResetString() = %q, want %q", ResetString(buf.String()), "Hello, World")
	}

	buf.Reset()
	WriteBytes(&buf, Green, []byte("Hello"), []byte(", World"))
	if string(ResetBytes(buf.Bytes())) != "Hello, World" {
		t.Errorf("ResetBytes() = %q, want %q", ResetBytes(buf.Bytes()), []byte("Hello, World"))
	}
}

func TestColorizePrint(t *testing.T) {
	type args struct {
		s string
		c string
	}
	tests := []args{

		// Basic colors
		{},
		{s: " black ", c: Black},
		{s: " red ", c: Red},
		{s: " green ", c: Green},
		{s: " yellow ", c: Yellow},
		{s: " blue ", c: Blue},
		{s: " magenta ", c: Magenta},
		{s: " cyan ", c: Cyan},
		{s: " white ", c: White},
		// Reversed colors
		{},
		{s: " black reversed ", c: BlackReversed},
		{s: " red reversed ", c: RedReversed},
		{s: " green reversed ", c: GreenReversed},
		{s: " yellow reversed ", c: YellowReversed},
		{s: " blue reversed ", c: BlueReversed},
		{s: " magenta reversed ", c: MagentaReversed},
		{s: " cyan reversed ", c: CyanReversed},
		{s: " white reversed ", c: WhiteReversed},
		// Bright colors
		{},
		{s: " bright black ", c: BrightBlack},
		{s: " bright red ", c: BrightRed},
		{s: " bright green ", c: BrightGreen},
		{s: " bright yellow ", c: BrightYellow},
		{s: " bright blue ", c: BrightBlue},
		{s: " bright magenta ", c: BrightMagenta},
		{s: " bright cyan ", c: BrightCyan},
		{s: " bright white ", c: BrightWhite},
		// Bright reverse colors
		{},
		{s: " bright black reversed ", c: BrightBlackReversed},
		{s: " bright red reversed ", c: BrightRedReversed},
		{s: " bright green reversed ", c: BrightGreenReversed},
		{s: " bright yellow reversed ", c: BrightYellowReversed},
		{s: " bright blue reversed ", c: BrightBlueReversed},
		{s: " bright magenta reversed ", c: BrightMagentaReversed},
		{s: " bright cyan reversed ", c: BrightCyanReversed},
		{s: " bright white reversed ", c: BrightWhiteReversed},

		{},
	}

	println("=== String ===")
	for _, tt := range tests {
		println(String(tt.c, tt.s))
	}

	println("=== Bytes ===")
	for _, tt := range tests {
		println(string(Bytes(tt.c, []byte(tt.s))))
	}

	println("=== Reset String ===")
	for _, tt := range tests {
		println(ResetString(String(tt.c, tt.s)))
	}

	println("=== Reset Bytes ===")
	for _, tt := range tests {
		println(string(ResetBytes(Bytes(tt.c, []byte(tt.s)))))
	}
}
