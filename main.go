/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>
*/
package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Alvesafk/gopass/cmd"
	"github.com/Alvesafk/gopass/storage"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		cmd.PrintUsage(args)
		log.Fatal("Invalid operation, try again.")
	}

	home, _ := os.UserHomeDir()
	db_path := filepath.Join(home, ".gopass", "secrets.db")
	os.MkdirAll(filepath.Dir(db_path), 0755)

	db, err := storage.New(db_path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// todo :: rewrite init function, it will ask for a master password, the function
	// must tell the user that this password is not saved anywhere on the db of the app
	// so the user must remember or store the password some where
	switch args[1] {
	case "Init", "init", "I", "i":
		cmd.Init(*db)
	case "Add", "add", "A", "a":
		cmd.Add(*db)
	case "List", "list", "L", "l":
		cmd.List(*db)
	case "Delete", "delete", "D", "d":
		cmd.Delete(*db, args)
	case "Get", "get", "G", "g":
		cmd.Get(*db, args)
	default:
		cmd.PrintUsage(args)
		log.Fatal("Invalid command, try again")
	}
}
