package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Alvesafk/agopass/color"
	"github.com/Alvesafk/agopass/storage"
)

func Update(db storage.DB, args []string) {
	CheckAmountArguments(args)

	mk := Authenticate(db)

	reader := bufio.NewReader(os.Stdin)

	to_change_secret, err := db.GetByName(args[2])
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
			to_change_secret = &probable_secret
		default:
			fmt.Println("Ok! Exiting the program.")
			os.Exit(1)
		}
	}

	fmt.Println("---------------~Update~----------------")
	fmt.Printf(color.Green("Updating %s secret.", "bold", 1), to_change_secret.Name)

	fmt.Println(string(mk), to_change_secret)
}
