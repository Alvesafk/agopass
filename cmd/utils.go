package cmd

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/Alvesafk/agopass/color"
	"github.com/Alvesafk/agopass/storage"
	"golang.org/x/term"
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

	return mk.Key == string(storage.HashMasterKey(s)), nil
}

func Authenticate(db storage.DB) []byte {
	_, err := db.MasterKeyExists()
	if err != nil {
		fmt.Print(color.Red("Master key does not exist! Use <gopass init> to add a master key.", "bold", 1))
		os.Exit(1)
		return nil
	}

	for range MAX_PASSWORD_RETRIES {
		fmt.Print(color.White("Enter with your master key: ", "bold", 0))
		password, err := term.ReadPassword(int(syscall.Stdin)) 
		if err != nil {
			fmt.Print(color.Red("Could not read the password input.", "bold", 1))
			os.Exit(1)
		}

		is, err := IsMasterKeyHash(db, strings.TrimSpace(string(password)))
		if err != nil {
			fmt.Print(color.Red("Could not access master key from db", "bold", 1))
			os.Exit(1)
		}

		if !is {
			fmt.Print(color.Red("Input doesn't match master key.", "bold", 1))
			continue
		} else {
			fmt.Print(color.Green("Authenticated.", "underline", 1))
			return storage.HashMasterKey(string(password))
		}
	}

	fmt.Print(color.Red("Could not authenticate, aborting.", "bold", 1))
	os.Exit(1)
	return nil
}
