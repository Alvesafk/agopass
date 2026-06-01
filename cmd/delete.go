/*
Delete command, it deletes a secret, the secret to be deleted must be passed with the
command itself, it's case sensitive, so GITHUB it's different from github.
Ex.: agopass delete Github
*/
package cmd

import (
	"bufio"
	"fmt"
	"strings"
	"os"

	"github.com/Alvesafk/agopass/color"
	"github.com/Alvesafk/agopass/storage"
)

// Delete function, receives a db connection and the arguments.
func Delete(db storage.DB, args []string) {
	// Check the amount of arguments on the command, if it's less or more then 3 will
	// abort the program after printing the error to the user.
	CheckAmountArguments(args)

	// Auth function
	Authenticate(db)

	to_get_name := strings.TrimSpace(args[2])

	// db.GetByName() method returns a initialized Struct of a secret if is found by
	// the name that is used as argument.
	to_get_secret, err := db.GetByName(to_get_name)
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

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
			fmt.Printf(color.Red("Error: Could not delete %s.", "bold", 1), to_get_name)
			fmt.Println(err)
			return
		}

		fmt.Printf(color.Green("Success! %s secret was deleted.", "bold", 1), to_get_name)

	default:
		fmt.Printf(color.White("%s secret was not deleted.", "bold", 1), to_get_name)
		return 
	}
}

/*
Index:
func Delete(db storage.DB, args []string)
*/
