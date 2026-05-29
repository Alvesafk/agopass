package cmd

import (
	"fmt"
	"strings"

	"github.com/Alvesafk/gopass/color"
	"github.com/Alvesafk/gopass/storage"
	"github.com/atotto/clipboard"
)

func Get(db storage.DB, args []string) {
	l := len(args)
	if l != 3 {
		if l > 3 {
			fmt.Print(color.Red("Error: Too many arguments.", "bold", 1))
			PrintUsage(args)
			return
		}

		if l < 3 {
			fmt.Print(color.Red("Error: Missing arguments.", "bold", 1))
			PrintUsage(args)
			return
		}
	}

	to_get_name := strings.TrimSpace(args[2])

	to_get_secret, err := db.GetByName(to_get_name)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = clipboard.WriteAll(to_get_secret.Key)
	if err != nil {
		panic(err)
	}

	fmt.Printf(color.Green("Success! %s key is in your clipboard.", "bold", 1), to_get_secret.Name)
}

