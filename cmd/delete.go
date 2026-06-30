/*
Delete command, it deletes a secret, the secret to be deleted must be passed with the
command itself, it's case sensitive, so GITHUB it's different from github.
Ex.: agopass delete Github
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Alvesafk/agopass/storage"
	"github.com/Alvesafk/scolor"
	"github.com/Alvesafk/scolor/ansi"
)

// Delete function, receives a db connection and the arguments.
func Delete(db storage.DB, args []string) {
	// Check the amount of arguments on the command, if it's less or more then 3 will
	// abort the program after printing the error to the user.
	CheckAmountArguments(args)

	// Auth function
	Authenticate(db)

	// db.GetByName() method returns a initialized Struct of a secret if is found by
	// the name that is used as argument.
	reader := bufio.NewReader(os.Stdin)
	to_get_secret, err := db.GetByName(args[2])
	if err != nil {
		probable_secret, err := CheckArgumentSpelling(args, db)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		ansi.Yellow.FgPrintf("Wasn't able to retrieve a exact match, did you mean %s? y/N : ", probable_secret.Name)
		response, err := reader.ReadString('\n')
		if err != nil {
			ansi.Red.FgPrintln("Error:", err)
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

	fmt.Println("---------------~Delete~----------------")

	fmt.Printf("Are you sure you want to delete %s? y/N : ", to_get_secret.Name)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	switch strings.TrimSpace(response) {
	case "Y", "y", "yes", "Yes", "YES":
		// db.Delete() method deletes a row of the DB based on the ID that is passed
		// that's why we first need to get the secret struct.
		err = db.Delete(to_get_secret.ID)
		if err != nil {
			ansi.Red.BgPrintln("Error:", err)
			fmt.Println(err)
			return
		}

		ansi.Green.FgPrintf("Success! %s secret was deleted.\n", to_get_secret.Name)

	default:
		fmt.Printf(scolor.AddMod("%s secret was not deleted.", scolor.Bold), to_get_secret.Name)

		return
	}
}

/*
Index:
func Delete(db storage.DB, args []string)
*/
