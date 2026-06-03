package cmd

import (
	"fmt"

	"github.com/Alvesafk/agopass/storage"
)

func Update(db storage.DB, args []string) {
	CheckAmountArguments(args)

	Authenticate(db)

	fmt.Println("Hello from the update command.")
}
