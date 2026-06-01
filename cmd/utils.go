/*
Utils, this file has function that: are used on more than one file and/or i didn't know
where to put.
*/
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

// Prints the agopass usage.
func PrintUsage(args []string) {
	fmt.Printf("Usage instruction:\n%s <COMMAND>\nCOMMANDS:\ninit\nadd\nlist\ndelete\nget\n", args[0])
}

// Check the amount of arguments passed on a args string slice, usually on the commands
// that need arguments only 3 total are needed, the binary, the command itself and an
// argument, more than this it's reject, less then 3 is also reject for this function is
// only called on commands that accept arguments.
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

// IsMasterKeyHash function compares a string (that will be hashed by the function) with
// the hashed Master Key on DB, returning true if it is the Master Key, and false if not.
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

// Authenticate function, receives a DB and returns a slice of bytes that is the hashed
// Master Key, it checks if the key exist, if it exists will prompt the user for it's
// Master Key, the user has 3 tries to get it right, if it fails the program will abort.
// If they get it right the program follows the normal flow.
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

/*
Index:
func PrintUsage(args []string)
func CheckAmountArguments(args []string)
func IsMasterKeyHash(db storage.DB, s string) (bool, error
func Authenticate(db storage.DB) []byte
*/
