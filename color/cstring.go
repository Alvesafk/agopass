/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>
*/

/*
color lib, cstring.go, the intent of this package is to format strings in a easier and
more legible way, instead of using multiple escape sequences, making the string illegible
you can just call color.<Color>(string-to-format, modification<bold, underline...>,
amount of new line escape).
*/

package color

import (
	"strings"
)

const (
	reset   = "\033[0m"
	cRed    = "\033[31m"
	cGreen  = "\033[32m"
	cYellow = "\033[33m"
	cBlue   = "\033[34m"
	cPurple = "\033[35m"
	cCyan   = "\033[36m"
	cWhite  = "\033[37m"
)

// AddMod function adds a "modifier" to a string, a modifier is just a escape sequence that
// changes the way the text is rendered, you have: "bold", "underline", "strike", "italic",
// for nothing.
func AddMod(s, mod string) string {
	switch mod {
	case "bold":
		s = "\033[1m" + s
	case "underline":
		s = "\033[4m" + s
	case "strike":
		s = "\033[9m" + s
	case "italic":
		s = "\033[3m" + s
	default:
		return s
	}

	return s
}

// All color functions are just a wrapper of this one, it's the one that adds the color and
// new line escape sequences.
func colorize(s, mod, code string, escape int) string {
	r := code + AddMod(s, mod) + reset
	if escape > 0 {
		r += strings.Repeat("\n", escape)
	}

	return r
}

func Red(s, mod string, escape int) string    { return colorize(s, mod, cRed, escape) }
func Green(s, mod string, escape int) string  { return colorize(s, mod, cGreen, escape) }
func Yellow(s, mod string, escape int) string { return colorize(s, mod, cYellow, escape) }
func Blue(s, mod string, escape int) string   { return colorize(s, mod, cBlue, escape) }
func Purple(s, mod string, escape int) string { return colorize(s, mod, cPurple, escape) }
func Cyan(s, mod string, escape int) string   { return colorize(s, mod, cCyan, escape) }
func White(s, mod string, escape int) string  { return colorize(s, mod, cWhite, escape) }
/*
func AddMod(s, mod string) string
func Red(s string, mod string, escape int) string
func Green(s string, mod string, escape int) string
func Yellow(s string, mod string, escape int) string
func Blue(s string, mod string, escape int) string
func Purple(s string, mod string, escape int) string
func Cyan(s string, mod string, escape int) string
func White(s string, mod string, escape int) string
*/
