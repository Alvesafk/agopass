package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Alvesafk/agopass/color"
	"github.com/Alvesafk/agopass/storage"
)

func Add(db storage.DB) {
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
