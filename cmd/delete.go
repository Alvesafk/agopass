package cmd

import (
	"bufio"
	"fmt"
	"strings"
	"os"

	"github.com/Alvesafk/agopass/color"
	"github.com/Alvesafk/agopass/storage"
)

func Delete(db storage.DB, args []string) {
	Authenticate(db)

	CheckAmountArguments(args)

	to_get_name := strings.TrimSpace(args[2])

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
		err = db.Delete(to_get_secret.ID)
		if err != nil {
			fmt.Printf(color.Red("Error: Could not delete %s.", "bold", 1), to_get_name)
			fmt.Println(err)
			return
		}

		fmt.Printf(color.Green("Success! %s secret was deleted.", "bold", 1), to_get_name)

	case "N", "n", "no", "No", "NO":
		fmt.Printf(color.White("Ok! %s secret was not deleted.", "bold", 1), to_get_name)
		return 
	default:
		fmt.Printf(color.White("%s secret was not deleted.", "bold", 1), to_get_name)
		return 
	}
}
