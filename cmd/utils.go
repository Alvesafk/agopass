package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Alvesafk/gopass/color"
	"github.com/Alvesafk/gopass/storage"
)

const (
	MAX_PASSWORD_RETRIES = 3
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

func Authenticate(db storage.DB) {
	reader := bufio.NewReader(os.Stdin)

	for range MAX_PASSWORD_RETRIES {
		fmt.Print(color.White("Enter with your master key: ", "bold", 0))
		password, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print(color.Red("Could not read the password input.", "bold", 1))
			os.Exit(1)
		}

		is, err := IsMasterKeyHash(db, strings.TrimSpace(password))
		if err != nil {
			fmt.Print(color.Red("Could not access master key from db", "bold", 1))
			os.Exit(1)
		}

		if !is {
			fmt.Print(color.Red("Input doesn't match master key.", "bold", 1))
			continue
		} else {
			fmt.Print(color.Green("Authenticated.", "underline", 1))
			return
		}
	}

	fmt.Print(color.Red("Could not authenticate, aborting.", "bold", 1))
	os.Exit(1)
}
