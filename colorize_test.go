package colorize

import (
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
			"\x1b[30mtest\x1b[0m",
		},
		{
			args{s: "Red color", c: Red},
			"\x1b[31mtest\x1b[0m",
		},
		{
			args{s: "Green color", c: Green},
			"\x1b[32mtest\x1b[0m",
		},
		{
			args{s: "Yellow color", c: Yellow},
			"\x1b[33mtest\x1b[0m",
		},
		{
			args{s: "Blue color", c: Blue},
			"\x1b[34mtest\x1b[0m",
		},
		{
			args{s: "Magenta color", c: Magenta},
			"\x1b[35mtest\x1b[0m",
		},
		{
			args{s: "Cyan color", c: Cyan},
			"\x1b[36mtest\x1b[0m",
		},
		{
			args{s: "White color", c: White},
			"\x1b[37mtest\x1b[0m",
		},
		// Reversed colors
		{
			args{s: "BlackReversed color", c: BlackReversed},
			"\x1b[40mtest\x1b[0m",
		},
		{
			args{s: "RedReversed color", c: RedReversed},
			"\x1b[41mtest\x1b[0m",
		},
		{
			args{s: "GreenReversed color", c: GreenReversed},
			"\x1b[42mtest\x1b[0m",
		},
		{
			args{s: "YellowReversed color", c: YellowReversed},
			"\x1b[43mtest\x1b[0m",
		},
		{
			args{s: "BlueReversed color", c: BlueReversed},
			"\x1b[44mtest\x1b[0m",
		},
		{
			args{s: "MagentaReversed color", c: MagentaReversed},
			"\x1b[45mtest\x1b[0m",
		},
		{
			args{s: "CyanReversed color", c: CyanReversed},
			"\x1b[46mtest\x1b[0m",
		},
		{
			args{s: "WhiteReversed color", c: WhiteReversed},
			"\x1b[47mtest\x1b[0m",
		},
		// Bright colors
		{
			args{s: "BrightBlack color", c: BrightBlack},
			"\x1b[90mtest\x1b[0m",
		},
		{
			args{s: "BrightRed color", c: BrightRed},
			"\x1b[91mtest\x1b[0m",
		},
		{
			args{s: "BrightGreen color", c: BrightGreen},
			"\x1b[92mtest\x1b[0m",
		},
		{
			args{s: "BrightYellow color", c: BrightYellow},
			"\x1b[93mtest\x1b[0m",
		},
		{
			args{s: "BrightBlue color", c: BrightBlue},
			"\x1b[94mtest\x1b[0m",
		},
		{
			args{s: "BrightMagenta color", c: BrightMagenta},
			"\x1b[95mtest\x1b[0m",
		},
		{
			args{s: "BrightCyan color", c: BrightCyan},
			"\x1b[96mtest\x1b[0m",
		},
		{
			args{s: "BrightWhite color", c: BrightWhite},
			"\x1b[97mtest\x1b[0m",
		},
		// Bright reverse colors
		{
			args{s: "BrightBlackReversed color", c: BrightBlackReversed},
			"\x1b[100mtest\x1b[0m",
		},
		{
			args{s: "BrightRedReversed color", c: BrightRedReversed},
			"\x1b[101mtest\x1b[0m",
		},
		{
			args{s: "BrightGreenReversed color", c: BrightGreenReversed},
			"\x1b[102mtest\x1b[0m",
		},
		{
			args{s: "BrightYellowReversed color", c: BrightYellowReversed},
			"\x1b[103mtest\x1b[0m",
		},
		{
			args{s: "BrightBlueReversed color", c: BrightBlueReversed},
			"\x1b[104mtest\x1b[0m",
		},
		{
			args{s: "BrightMagentaReversed color", c: BrightMagentaReversed},
			"\x1b[105mtest\x1b[0m",
		},
		{
			args{s: "BrightCyanReversed color", c: BrightCyanReversed},
			"\x1b[106mtest\x1b[0m",
		},
		{
			args{s: "BrightWhiteReversed color", c: BrightWhiteReversed},
			"\x1b[107mtest\x1b[0m",
		},
		// Test with empty string
		{
			args{s: "Empty string with Red color", c: Red},
			"\x1b[31m\x1b[0m",
		},
		// Test with different text
		{
			args{s: "Different text with Blue color", c: Blue},
			"\x1b[34mHello World\x1b[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.args.s, func(t *testing.T) {
			got := Colorize(tt.args.s, tt.args.c)
			if got != tt.want {
				t.Errorf("Colorize() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestColorizeCheck(t *testing.T) {
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

	for _, tt := range tests {
		println(Colorize(tt.s, tt.c))
	}
}
