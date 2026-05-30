package cmd

import (
	"fmt"
	"os"

	"github.com/Alvesafk/gopass/color"
	"github.com/Alvesafk/gopass/storage"
)

func PrintUsage(args []string) {
	fmt.Printf("Usage instruction:\n%s <COMMAND>\nCOMMANDS:\ninit\nadd\nlist\ndelete\nget\n", args[0])
}

func CheckAmountArguments(args []string) {
	l := len(args)
	if l != 3 {
		if l > 3 {
			fmt.Print(color.Red("Error: Too many arguments.", "bold", 1))
			PrintUsage(args)
			os.Exit(1)
		}

		if l < 3 {
			fmt.Print(color.Red("Error: Missing arguments.", "bold", 1))
			PrintUsage(args)
			os.Exit(1)
		}
	}
}

func IsMasterKeyHash(db storage.DB, s string) (bool, error) {
	_, err := db.MasterKeyExists()
	if err != nil {
		return false, fmt.Errorf("Master key does not exist.")
	}
	
	mk, err := db.GetHashedMasterKey()
	if err != nil {
		return false, fmt.Errorf("Could not get master key.")
	}

	return mk.Key == s, nil
}
