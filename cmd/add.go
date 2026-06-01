/*
The Add command is for adding secrets into the DB, the encryption part is handled inside
the storage package.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Alvesafk/agopass/color"
	"github.com/Alvesafk/agopass/storage"
)

// Add function, it receives a DB connection.
func Add(db storage.DB) {
	// Authenticate function, it auth the user based on the MK, if the user get their
	// password wrong tree times in a row the program is aborted, if they get it right
	// the program continues and the MK Hash it's returned via the Authenticate func
	// it receives only one argument, a DB connection.
	mk := Authenticate(db)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("-----------------~Add~-----------------")

	fmt.Print("Name of the secret: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(color.Red("Could not read the name input.", "bold", 1))
		return
	}

	fmt.Print("Secret key: ")
	secret, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(color.Red("Could not read the key input.", "bold", 1))
		return
	}

	fmt.Println()
	fmt.Printf("Are you sure you want to add %s? y/N : ", strings.TrimSpace(name))
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	switch strings.TrimSpace(response) {
	case "Y", "y", "yes", "Yes", "YES":
		// db.Insert() is a method, receives 2 strings, the name and the key of the
		// new secret.
		_, err = db.Insert(strings.TrimSpace(name), strings.TrimSpace(secret), mk)
		if err != nil {
			fmt.Print(color.Red("Error: Could not insert into DB.", "bold", 1))
			fmt.Println(err)
			return
		}

	default:
		fmt.Print(color.White("Ok! Secret was not registered.", "bold", 1))
		return 
	}

	fmt.Print(color.Green("Secret was saved!", "bold", 1))
}

/*
Index:
func Add(db storage.DB)
*/
