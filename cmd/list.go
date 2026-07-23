/*
List command shows all secrets that are registered on the DB, the keys that it shows are
just a mock based on the length of the plain text key, therefore the user does not need to
authenticate before using it.
*/
package cmd

import (
	"math"
	"os"

	"github.com/Alvesafk/agopass/storage"
	"github.com/Alvesafk/scolor/ansi"
	"github.com/jedib0t/go-pretty/v6/table"
)

// List function, accepts a DB connection.
func List(db storage.DB) {
	// db.List() method does not have any argument (only the implicit DB one), the
	// method returns a slice of initialized Secret structs.
	all_secrets, err := db.List()
	if err != nil {
		ansi.Red.FgPrintln("Error:", err)
		return
	}

	// If no secret exist it will tell the user to add them.
	if len(all_secrets) < 1 {
		ansi.Red.FgPrintln("No secret registered! Use <gopass add> to add secrets.")
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.SetStyle(table.StyleLight)

	t.AppendHeader(table.Row{"Secrets", "Name", "Key"})

	// Range over the slice of secrets, printing them into your terminal.
	for i, v := range all_secrets {
		t.AppendRow(table.Row{i + 1, v.Name, hidePassword(v.Key_Length)})
	}

	t.Render()
}

// hide password function accepts a integer representing the length of the password, the
// return is a string made of '*'. Ex.: Pass123 -> len = 7 -> *******
func hidePassword(kl int) string {
	// Initialize the result string.
	var result string

	// Get the smallest of this two, the limit length is 25 characters.
	l := int(math.Min(float64(kl), 25))

	// Range over l(int) adding a '*' onto the result string in every iteration.
	for range l {
		result += "*"
	}

	// If the kl(key length) it's greater than 25 add a "..." on the end of the string.
	if kl > 25 {
		result = result + "..."
	}

	return result
}

/*
Index:
func List(db storage.DB)
func hidePassword(kl int) string
*/
