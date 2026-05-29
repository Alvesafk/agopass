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


	switch args[1] {
	case "Init", "init":
		cmd.Init()
	case "Add", "add":
		cmd.Add(*db)
	case "List", "list":
		cmd.List(*db)
	case "Delete", "delete":
		cmd.Delete(args)
	case "Get", "get":
		cmd.Get(*db ,args)
	default:
		cmd.PrintUsage(args)
		log.Fatal("Invalid command, try again")
	}
}
