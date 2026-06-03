package cmd

import (
	"fmt"
	"os"

	"github.com/Alvesafk/agopass/color"
	"github.com/Alvesafk/agopass/storage"
)

func Update(db storage.DB, args []string) {
	CheckAmountArguments(args)

	mk := Authenticate(db)

	to_change_secret, err := db.GetByName(args[2])
	if err != nil {
		probable_secret, err := CheckArgumentSpelling(args, db)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(probable_secret.Name)
		os.Exit(1)
	}

	fmt.Println("---------------~Update~----------------")
	fmt.Printf(color.Green("Updating %s secret.", "bold", 1), args[2])

	fmt.Println(mk, to_change_secret)
}
