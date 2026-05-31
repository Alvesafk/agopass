package cmd

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/Alvesafk/gopass/color"
	"github.com/Alvesafk/gopass/storage"
	"golang.org/x/term"
)

func Init(db storage.DB) {
	_, err := db.MasterKeyExists()
	switch err {
	case sql.ErrNoRows:
		reader := bufio.NewReader(os.Stdin)

		fmt.Print(color.Green("No master key was found, create one: ", "bold", 0))
		mk, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println()
			fmt.Print(color.Red("Could not read the master key input.", "bold", 1))
			os.Exit(1)
		}
		fmt.Println()

		fmt.Println("This key can NOT be retrieved from the Database.")
		fmt.Println("This key will be used before the <add/delete/get> commands.")
		fmt.Printf("Make sure you don't forget it! =D\n\n")

		fmt.Print(color.Yellow("Knowing this, are you sure that is the password you want to define? y/N : ", "bold", 0))
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print(color.Red("Could not read the response input.", "bold", 1))
			os.Exit(1)
		}

		switch strings.TrimSpace(response) {
		case "Y", "y", "yes", "Yes", "YES":
			_, err = db.AddMasterKey(strings.TrimSpace(string(mk)))
			if err != nil {
				fmt.Print(color.Red("Could not save the master key, try again.", "bold", 1))
				os.Exit(1)
			}

			fmt.Print(color.Green("Master key was saved on the DB!", "bold", 1))
		default:
			fmt.Println("Ok, Master key was not saved!")
			os.Exit(0)
		}
	case nil:
		fmt.Print(color.Yellow("Master key is already on the DB, you didn't forget it? Right?", "underline", 1))
		os.Exit(0)
	default:
		fmt.Print(color.Red("Was not possible to verify existance of master key on DB, aborting", "bold", 1))	
		os.Exit(1)
	}
}
