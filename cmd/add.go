package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Alvesafk/gopass/color"
	"github.com/Alvesafk/gopass/storage"
)

type Secret struct {
	Name string
	Key string
}

func Add(db storage.DB) {
	fmt.Print(color.Green("Add a Secret:", "bold", 1))

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(color.Yellow("Name of the secret: ", "none", 0))
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(color.Yellow("Secret word: ", "none", 0))
	secret, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	formated_name := strings.ToLower(strings.TrimSpace(name))

	_, err = db.Insert(formated_name, strings.TrimSpace(secret))
	if err != nil {
		fmt.Print(color.Red("Error: Could not insert into Db.", "bold", 1))
		fmt.Println(err)
		return
	}

	fmt.Print(color.Green("Secret was saved!", "bold", 1))
}
