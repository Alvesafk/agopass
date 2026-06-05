/*
Get command is used to get the actual secret from the DB, it will search the secret that
you want based on the name passed on the command it is case sensitive, GITHUB is different
from github, if it finds will copy the key to your clipboard.
Ex.: agopass get Github
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Alvesafk/agopass/color"
	"github.com/Alvesafk/agopass/storage"
	"github.com/atotto/clipboard"
)

// Get function receives a DB connection and a arguments slice.
func Get(db storage.DB, args []string) {
	// Check the amount of arguments passed with the command.
	CheckAmountArguments(args)

	// Authenticate user based on master key, if it authenticates the hash of the MK
	// is returned.
	mk := Authenticate(db)

	// Get a initialized struct of the key by name.
	reader := bufio.NewReader(os.Stdin)
	to_get_secret, err := db.GetByName(args[2])
	if err != nil {
		probable_secret, err := CheckArgumentSpelling(args, db)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf(color.Yellow("Wasn't able to retrieve a exact match, did you mean %s? y/N : ", "bold", 0), probable_secret.Name)
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print(color.Red("Wasn't able to retrieve the reponse input, aborting.", "bold", 1))
			os.Exit(1)
		}

		switch strings.TrimSpace(response) {
		case "Yes", "yes", "YES", "y", "Y":
			to_get_secret = &probable_secret
		default:
			fmt.Println("Ok! Exiting the program.")
			os.Exit(1)
		}
	}

	// storage.Decrypt() accepts a string and the hashed Master Key, using them the
	// function decrypts the key returnin the real secret.
	decrypted_key, err := storage.Decrypt(to_get_secret.Key, mk)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Stores on the clipboard.
	err = clipboard.WriteAll(decrypted_key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf(color.Green("Success! %s key is in your clipboard.", "bold", 1), to_get_secret.Name)
}

/*
Index:
func Get(db storage.DB, args []string)
*/
