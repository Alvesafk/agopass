package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Alvesafk/gopass/color"
	"github.com/Alvesafk/gopass/storage"
)

func Add(db storage.DB) {
	Authenticate(db)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("-----------------~Add~-----------------")

	fmt.Print("Name of the secret: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(color.Red("Could not read the name input.", "bold", 1))
		return
	}

	fmt.Print("Secret word: ")
	secret, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(color.Red("Could not read the key input.", "bold", 1))
		return
	}

	_, err = db.Insert(strings.TrimSpace(name), strings.TrimSpace(secret))
	if err != nil {
		fmt.Print(color.Red("Error: Could not insert into DB.", "bold", 1))
		fmt.Println(err)
		return
	}

	fmt.Print(color.Green("Secret was saved!", "bold", 1))
}
