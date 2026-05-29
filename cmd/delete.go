package cmd

import (
	"fmt"
	"strings"

	"github.com/Alvesafk/gopass/color"
	"github.com/Alvesafk/gopass/storage"
)

func Delete(db storage.DB, args []string) {
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

	err = db.Delete(to_get_secret.ID)
	if err != nil {
		fmt.Printf(color.Red("Error: Could not delete %s.", "bold", 1), to_get_name)
		fmt.Println(err)
		return
	}

	fmt.Printf(color.Green("Success! %s secret was deleted.", "bold", 1), args[2])
}

