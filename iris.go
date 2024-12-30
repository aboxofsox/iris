package iris

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Color int

const Reset = 0

func (c Color) Int() int { return int(c) }

// SetColor sets the color of the given text using the hexidecimal representation of a color.
func SetColor(s, fg, bg string) string {
	fc, bc := convert(fg), convert(bg)
	return ansi(fc, bc) + s + ansi(Reset, Reset)
}

// SetFgColor sets the foreground color of the given text.
func SetFgColor(s, fg string) string {
	return ansiFg(convert(fg)) + s + ansi(Reset, Reset)
}

// SetBgColor sets teh background color of the given text.
func SetBgColor(s, bg string) string {
	return ansiBg(convert(bg)) + s + ansi(Reset, Reset)
}

// Strip removes all ANSI escape charactes from a string.
func Strip(s string) string {
	rgex := regexp.MustCompile(`\x1b\[[0-9;]*[mG]`)
	return rgex.ReplaceAllString(s, "")
}

func ansi(fg, bg Color) string {
	if fg == Reset || bg == Reset {
		return "\033[0m"
	}
	fgAnsi := fmt.Sprintf("\033[38;5;%dm", fg)
	bgAnsi := fmt.Sprintf("\033[48;5;%dm", bg)
	return fmt.Sprintf("%s%s", fgAnsi, bgAnsi)
}

func ansiFg(c Color) string {
	if c == Reset {
		return "\033[0m"
	}
	return fmt.Sprintf("\033[38;5;%dm", c)
}

func ansiBg(c Color) string {
	return fmt.Sprintf("\033[48;5;%dm", c)
}

func convert(hx string) Color {
	r, g, b := rgb(hx)
	idx := colorIndex(r, g, b)
	return Color(idx)
}

func colorIndex(r, g, b int) int {
	r, g, b = scale(r), scale(g), scale(b)
	return 16 + (36 * r) + (6 * g) + b
}

func scale(c int) int {
	return (c*5 + 127) / 255
}

/*
rgb converts a hexidecimal string to its respective RGB values.

	r, g, b := rgb("#FF0000")

The RGB value is determined by taking the hexdecimal representation of a color
and segmenting it by "nibbles". #FF0000 becomes (FF) (00) (00), which is then
converted to (255) (0) (0). The determintation is not 100% accurate, but would
probably be fine for most terminal emulators.
*/
func rgb(hx string) (r, g, b int) {
	hx = strings.TrimPrefix(hx, "#")
	if len(hx) != 6 {
		return
	}

	rs, gs, bs := hx[0:2], hx[2:4], hx[4:6]

	cints := make([]int, 3)
	for i, c := range []string{rs, gs, bs} {
		cint, err := strconv.ParseUint(c, 16, 8)
		if err != nil {
			return
		}
		cints[i] = int(cint)
	}
	r, g, b = cints[0], cints[1], cints[2]

	return
}
