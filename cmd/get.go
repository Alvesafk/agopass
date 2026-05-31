package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Alvesafk/gopass/color"
	"github.com/Alvesafk/gopass/storage"
	"github.com/atotto/clipboard"
)

func Get(db storage.DB, args []string) {
	mk := Authenticate(db)

	CheckAmountArguments(args)

	to_get_name := strings.TrimSpace(args[2])

	to_get_secret, err := db.GetByName(to_get_name)
	if err != nil {
		fmt.Println(err)
		return
	}

	decrypted_key, err := storage.Decrypt(to_get_secret.Key, mk)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = clipboard.WriteAll(decrypted_key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf(color.Green("Success! %s key is in your clipboard.", "bold", 1), to_get_secret.Name)
}

