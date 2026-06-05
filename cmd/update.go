/*
Update command is to update a already registered Secret on the DB, the user can change the
name, the key, but not it's ID.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/peterh/liner"
	"github.com/Alvesafk/agopass/color"
	"github.com/Alvesafk/agopass/storage"
)

// Update function receives a DB connection and the arguments string array, it returns 
// nothing.
func Update(db storage.DB, args []string) {
	// Check the amount of arguments, it must be 3.
	CheckAmountArguments(args)

	// Auth the user, we have private content being changed and printed.
	mk := Authenticate(db)

	// Open the reader.
	reader := bufio.NewReader(os.Stdin)

	// Tries to get secrety by name, using the third argument passed, the one after the
	// command.
	to_change_secret, err := db.GetByName(args[2])
	if err != nil {
		// If it's not possible to get the Secret based on the name given on the
		// command call the functions CheckArgumentSpelling is called, it receives
		// the arguments and a DB connection and tries to "guess" what the user
		// wanted based on what they wrote.
		probable_secret, err := CheckArgumentSpelling(args, db)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Prompts the user with the return value of the CheckArgumentSpelling function.
		fmt.Printf(color.Yellow("Wasn't able to retrieve a exact match, did you mean %s? y/N : ", "bold", 0), probable_secret.Name)
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print(color.Red("Wasn't able to retrieve the reponse input, aborting.", "bold", 1))
			os.Exit(1)
		}

		// Checks if they want to continue with this Secret or no.
		switch strings.TrimSpace(response) {
		case "Yes", "yes", "YES", "y", "Y":
			to_change_secret = &probable_secret
		default:
			fmt.Println("Ok! Exiting the program.")
			os.Exit(1)
		}
	}

	// Decrypt the key so it can be printed on the input.
	decrypted_key, err := storage.Decrypt(to_change_secret.Key, mk)
	if err != nil {
		fmt.Println("Was not possible to decrypt secret key, exiting.")
		os.Exit(1)
	}

	fmt.Println("---------------~Update~----------------")
	fmt.Printf(color.Green("Updating %s secret.", "bold", 1), to_change_secret.Name)

	// Liner is so we can print the Name and the Key of the secret on the Stdin, that
	// way the user can interact with it, they can just press enter and leave how it
	// already is, or delete it and write what they want.
	line := liner.NewLiner()

	// Gets the new name of the secret.
	new_name, err := line.PromptWithSuggestion("New name: ", to_change_secret.Name, -1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	// Gets the new key of the secret.
	new_key, err := line.PromptWithSuggestion("New key: ", decrypted_key, -1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	// Close the liner, it's closed here because this function make so the terminal is
	// in Raw Mode, if we don't close here the other stuff that we do with the normal
	// reader variable of Golang would be 'wrong', because the terminal would still be
	// in Raw Mode.
	line.Close()

	// Prompts the user if they are sure about the changes.
	fmt.Print("Are you sure you want to make this changes? y/N : ")
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(color.Red("Wasn't able to retrieve the reponse input, aborting.", "bold", 1))
		os.Exit(1)
		return
	}

	switch strings.TrimSpace(response) {
	// If they are, a new Secret instance will be created, the encryption of the key
	// is handled inside the Update method.
	case "Yes", "yes", "YES", "y", "Y":
		new_secret := storage.Secret{ID: to_change_secret.ID, Name: new_name, Key: new_key, Key_Length: len(new_key)}
		err = db.Update(to_change_secret.ID, new_secret, mk)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}

		fmt.Print(color.Green("Secret was updated on DB.", "bold", 1))
		os.Exit(0)
		return
	// If they are not sure the program is finished, maybe i will change this after,
	// prompt the user if they want to re-change.
	default:
		fmt.Printf(color.AddMod("Ok! Not updating the %s secret. Exiting.", "bold"), to_change_secret.Name)
		os.Exit(0)
		return
	}
}

/*
Index:
func Update(db storage.DB, args []string)
*/
