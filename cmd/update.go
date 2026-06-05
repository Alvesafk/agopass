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

	decrypted_key, err := storage.Decrypt(to_change_secret.Key, mk)
	if err != nil {
		fmt.Println("Was not possible to decrypt secret key, exiting.")
		os.Exit(1)
	}

	fmt.Println("---------------~Update~----------------")
	fmt.Printf(color.Green("Updating %s secret.", "bold", 1), to_change_secret.Name)

	line := liner.NewLiner()

	new_name, err := line.PromptWithSuggestion("New name: ", to_change_secret.Name, -1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	new_key, err := line.PromptWithSuggestion("New key: ", decrypted_key, -1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	line.Close()

	fmt.Print("Are you sure you want to make this changes? y/N : ")
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(color.Red("Wasn't able to retrieve the reponse input, aborting.", "bold", 1))
		os.Exit(1)
		return
	}

	switch strings.TrimSpace(response) {
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
	default:
		fmt.Printf(color.AddMod("Ok! Not updating the %s secret. Exiting.", "bold"), to_change_secret.Name)
		os.Exit(0)
		return
	}
}
