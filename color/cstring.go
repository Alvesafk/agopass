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

// Red function, all the other functions follow the same format so i will not be writing
// stuff on them, the colors function accepts a base string, a modifier, and how many new
// lines you want on the end of the string, it returns a string formatted based on your
// choices.
func Red(s string, mod string, escape int) string {
	r := AddMod(s, mod)

	r = "\033[31m" + r + "\033[0m"

	if escape > 0 {
		for range escape {
			r += "\n"
		}
	}

	return r
}

func Green(s string, mod string, escape int) string {
	r := AddMod(s, mod)

	r = "\033[32m" + r + "\033[0m"

	if escape > 0 {
		for range escape {
			r += "\n"
		}
	}

	return r
}

func Yellow(s string, mod string, escape int) string {
	r := AddMod(s, mod)

	r = "\033[33m" + r + "\033[0m"

	if escape > 0 {
		for range escape {
			r += "\n"
		}
	}

	return r
}

func Blue(s string, mod string, escape int) string {
	r := AddMod(s, mod)

	r = "\033[34m" + r + "\033[0m"

	if escape > 0 {
		for range escape {
			r += "\n"
		}
	}

	return r
}

func Purple(s string, mod string, escape int) string {
	r := AddMod(s, mod)

	r = "\033[35m" + r + "\033[0m"

	if escape > 0 {
		for range escape {
			r += "\n"
		}
	}

	return r
}

func Cyan(s string, mod string, escape int) string {
	r := AddMod(s, mod)

	r = "\033[36m" + r + "\033[0m"

	if escape > 0 {
		for range escape {
			r += "\n"
		}
	}

	return r
}

func White(s string, mod string, escape int) string {
	r := AddMod(s, mod)

	r = "\033[37m" + r + "\033[0m"

	if escape > 0 {
		for range escape {
			r += "\n"
		}
	}

	return r
}

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
