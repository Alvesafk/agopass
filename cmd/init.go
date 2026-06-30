/*
This is the first stop on the cmd package, is also the first command that should be run
when using agopass, with it you define you Master Key, the only password that you
need to remember, this password is prompted whenever you want to use a command that
MODIFIES the content of DB, that means, List command will not need the auth via master key
because it doesn't change anything on the DB, it doesn't show the password either, just
the name with a mock hidden passwords, to get your password you use the command Get, this
is one that needs auth, the following commands need Auth : add, get, delete. Auth will
be called before this functions, after you use Init atleas one time, you can forget it
exists, you will not be able to change your password, because the encryption of you other
keys is done using the Master Key hash as salt, so if you forget you better of just
removing the DB out of existence, than you can start over again.
*/
package cmd

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"syscall"

	// Color lib is for easier colored strings withou the need to manually add every
	// escape sequence in history.
	"github.com/Alvesafk/agopass/storage"
	"github.com/Alvesafk/scolor/ansi"
	"golang.org/x/term"
)

// Init function, is the one that register your Master Key, it needs a DB connection that
// is passed through it's argumments.
func Init(db storage.DB) {
	// Checks if the DB already has a registered Master Key.
	_, err := db.MasterKeyExists()
	switch err {
	// This error is passed when no row is found with a Master Key, meaning, it has to
	// create a new one.
	case sql.ErrNoRows:
		reader := bufio.NewReader(os.Stdin)

		// That's one of the uses of the color lib, you have one function for every
		// color that can be made with escape sequences on print, you first put
		// your string, than a "mod", like bold or undeline, and the amount of new
		// lines on the end of the string.
		ansi.Green.FgPrintln("No master key was found, create one: ")
		// term.ReadPassword it's used to disable the echo, therefore when you type
		// your password nothing is shown.
		mk, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			ansi.Red.FgPrintln("\nError:", err)
			os.Exit(1)
		}
		fmt.Println()

		fmt.Println("This key can NOT be retrieved from the Database.")
		fmt.Println("This key will be used before the <add/delete/get> commands.")
		fmt.Printf("Make sure you don't forget it! =D\n\n")

		ansi.Yellow.FgPrint("Knowing this, are you sure that this is the password you want to define? y/N : ")
		response, err := reader.ReadString('\n')
		if err != nil {
			ansi.Red.FgPrintln("Error:", err)
			os.Exit(1)
		}

		switch strings.TrimSpace(response) {
		case "Y", "y", "yes", "Yes", "YES":
			// db.AddMasterKey() registers your master key based on the string
			// that is passed, it will hash before adding to DB.
			_, err = db.AddMasterKey(strings.TrimSpace(string(mk)))
			if err != nil {
				ansi.Red.FgPrintln("Error:", err)
				os.Exit(1)
			}

			ansi.Green.FgPrintln("Master key was saved on the DB!")
		default:
			fmt.Println("Ok, Master key was not saved!")
			os.Exit(0)
		}
	// This case happens when no error is returned from MasterKeyExists(), it means
	// that the key already exists.
	case nil:
		ansi.Yellow.FgPrintln("Master key is already on the DB, you didn't forget it? Right?")
		os.Exit(0)
	// Normal error if something strange go wrong for whatever reason.
	default:
		ansi.Red.FgPrintln("Was not possible to verify existence of master key on DB, aborting.")
		os.Exit(1)
	}

	autocomplete_exists, err := AutocompleteExists()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if !autocomplete_exists {
		fmt.Println()
		fmt.Println("Want to setup the autocomplete script? It will autocomplete the commands for you when you press <tab>")
		fmt.Println("Note: It WILL be appended to yout .rc (.bashrc, .zshrc), you can take this script to another file and source it, or leave it there.")
		ansi.Yellow.FgPrint("Knowing this, you want to setup the autocomplete script? y/N : ")

		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		if err != nil {
			ansi.Red.FgPrintln("Error:", err)
			os.Exit(1)
		}

		switch strings.TrimSpace(response) {
		case "Yes", "YES", "yes", "Y", "y":
			if err := InitAutocomplete(); err != nil {
				ansi.Red.FgPrintln("Error:", err)
			}
		default:
			fmt.Println("Ok! Autocomplete was NOT setup.")
		}
	}
}

/*
Index:
func Init(db storage.DB)
*/
