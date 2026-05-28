/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>
*/

/*
format lib, cstring.go, the intent of this package is to format strings in a easier and
more legible way, instead of using multiple escape sequences, making the string illegible
you can just call format.<Color>(string-to-format, modification<bold, underline...>, 
amount of new line escape).
*/

package color

func AddMod(s string, mod string) string {
	switch mod {
	case "bold":
		s = "\033[1m" + s
	case "underline":
		s = "\033[4m" + s
	case "strike":
		s = "\033[9m" + s
	case "italic":
		s = "\033[3m" + s
	case "none":
		return s
	default:
		return s
	}

	return s
}

func Red(s string, mod string, escape int) string {
	r := AddMod(s, mod)

	r = "\033[31m" + r + "\033[0m"

	if escape > 0 {
		for i := 0; i < escape; i++ {
			r += "\n"
		}
	}

	return r
}

func Green(s string, mod string, escape int) string {
	r := AddMod(s, mod)

	r = "\033[32m" + r + "\033[0m"

	if escape > 0 {
		for i := 0; i < escape; i++ {
			r += "\n"
		}
	}

	return r
}

func Yellow(s string, mod string, escape int) string {
	r := AddMod(s, mod)

	r = "\033[33m" + r + "\033[0m"

	if escape > 0 {
		for i := 0; i < escape; i++ {
			r += "\n"
		}
	}

	return r
}

func Blue(s string, mod string, escape int) string {
	r := AddMod(s, mod)

	r = "\033[34m" + r + "\033[0m"

	if escape > 0 {
		for i := 0; i < escape; i++ {
			r += "\n"
		}
	}

	return r
}

func Purple(s string, mod string, escape int) string {
	r := AddMod(s, mod)

	r = "\033[35m" + r + "\033[0m"

	if escape > 0 {
		for i := 0; i < escape; i++ {
			r += "\n"
		}
	}

	return r
}

func Cyan(s string, mod string, escape int) string {
	r := AddMod(s, mod)

	r = "\033[36m" + r + "\033[0m"

	if escape > 0 {
		for i := 0; i < escape; i++ {
			r += "\n"
		}
	}

	return r
}

func White(s string, mod string, escape int) string {
	r := AddMod(s, mod)

	r = "\033[37m" + r + "\033[0m"

	if escape > 0 {
		for i := 0; i < escape; i++ {
			r += "\n"
		}
	}

	return r
}

/*
func Red(s string, mod string, escape int) string {}
func Green(s string, mod string, escape int) string {}
func Yellow(s string, mod string, escape int) string {}
func Blue(s string, mod string, escape int) string {}
func Purple(s string, mod string, escape int) string {}
func Cyan(s string, mod string, escape int) string {}
func White(s string, mod string, escape int) string {}
*/
