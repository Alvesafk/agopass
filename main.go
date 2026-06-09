/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>

Agopass (formerly gopass but someone had already stole my flow) is a CLI Password manager
made in Go, it uses a SQLite db located on this path /home/<user>/.agopass/, all the
contentes of the DB are hashed or encrypted, it's simple to use, few commands but it gets
the job done, to start with agopass you have to run <agopass init>, that will create the
DB and prompt you to create your Master Password, the only password you'l need to remember!
*/
package main

import (
	// Std go lib.
	"log"
	"os"
	"path/filepath"

	// Libs made by me
	// This one is for the commands, all of the real work happens there.
	"github.com/Alvesafk/agopass/cmd"
	// This one is for the DB management, create a db, connect to it, encrypt text, etc...
	"github.com/Alvesafk/agopass/storage"
)

func main() {
	// Getting the args, if it's less then 2, print the usage of agopass and kill the
	// program. (It's less than 2 when the user just prompts <agopass> alone.
	args := os.Args
	if len(args) < 2 {
		cmd.PrintUsage(args)
		os.Exit(1)
	}

	// Get's the home dir of the user to create the: application folder and them the DB.
	home, _ := os.UserHomeDir()
	db_path := filepath.Join(home, ".agopass", "secrets.db")
	os.MkdirAll(filepath.Dir(db_path), 0755)

	// storage.New() is a function that creates a connection with the existing DB, it
	// also creates the DB does not exist when the program it's called.
	db, err := storage.New(db_path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Switch to get what user want to do, i could have used Cobra or something like
	// that, but i prefer doing it mysel, i believe aswell that CLI should be simple
	// and i don't think excessive amounts of tags really help on your application
	// so that's why agopass have now tags whatsoever, for now atleast.
	switch args[1] {
	case "Init", "init", "I", "i":
		cmd.Init(*db)
	case "Add", "add", "A", "a":
		cmd.Add(*db, args)
	case "List", "list", "L", "l":
		cmd.List(*db)
	case "Delete", "delete", "D", "d":
		cmd.Delete(*db, args)
	case "Get", "get", "G", "g":
		cmd.Get(*db, args)
	case "Update", "update", "U", "u":
		cmd.Update(*db, args)
	case "Version", "version", "V", "v":
		cmd.Version()
	default:
		cmd.PrintUsage(args)
		log.Fatal("Invalid command, try again")
	}
}
