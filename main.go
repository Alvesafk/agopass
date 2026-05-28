/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>
*/
package main

import (
	"os"
	"log"

	"github.com/Alvesafk/gopass/cmd"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		cmd.PrintUsage(args)
		log.Fatal("Invalid operation, try again.")
	}


	switch args[1] {
	case "Init", "init":
		cmd.Init()
	case "Add", "add":
		cmd.Add()
	case "List", "list":
		cmd.List()
	case "Delete", "delete":
		cmd.Delete()
	case "Get", "get":
		cmd.Get(args)
	default:
		cmd.PrintUsage(args)
		log.Fatal("Invalid command, try again")
	}
}
