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

	"github.com/Alvesafk/agopass/storage"
	"github.com/Alvesafk/scolor"
	"github.com/Alvesafk/scolor/ansi"
	"golang.org/x/term"
)

const (
	MAX_PASSWORD_RETRIES = 3
)

// Prints the agopass usage.
func PrintUsage(args []string) {
	fmt.Println("Usage instructions:")
	fmt.Printf("%s <COMMAND> [ARG]\n", args[0])
	fmt.Println("Init / init / I / i     :: Create DB and prompt for master key.")
	fmt.Println("Add / add / A / a       :: Create a secret in DB.")
	fmt.Println("List / list / L / l     :: List all registered secrets.")
	fmt.Println("Delete / delete / D / d :: Delete a secret.")
	fmt.Println("Get / get / G / g       :: Get the secret key.")
	fmt.Println("Update / update / U / u :: Modify a registered secret.")
	fmt.Println("Make / make / M / m     :: Create a random N bit key, pass a number to N,\n                        :: if not it defaults to 16.")
	fmt.Println("Auto / auto / Au / au   :: Setup autocomplete.")
}

// Check the amount of arguments passed on a args string slice, usually on the commands
// that need arguments only 3 total are needed, the binary, the command itself and an
// argument, more than this it's reject, less then 3 is also reject for this function is
// only called on commands that accept arguments.
func CheckAmountArguments(args []string) {
	l := len(args)
	if l != 3 {
		if l > 3 {
			ansi.Red.FgPrintln("Error: Too many arguments.")
			PrintUsage(args)
			os.Exit(1)
		}

		if l < 3 {
			ansi.Red.FgPrintln("Error: Missing arguments.")
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
		return false, fmt.Errorf("master key does not exist")
	}

	mk, err := db.GetHashedMasterKey()
	if err != nil {
		return false, fmt.Errorf("could not get master key")
	}

	return mk.Key == string(storage.HashMasterKey(s)), nil
}

// Authenticate function, receives a DB and returns a slice of bytes that is the hashed
// Master Key, it checks if the key exist, if it exists will prompt the user for it's
// Master Key, the user has 3 tries to get it right, if it fails the program will abort.
// If they get it right the program follows the normal flow.
func Authenticate(db storage.DB) []byte {
	tmp_mk, err := loadTmpHash()
	if err == nil {
		return tmp_mk
	}

	_, err = db.MasterKeyExists()
	if err != nil {
		ansi.Red.FgPrintln("Master key does not exist! Use <gopass init> to add a master key.")
		os.Exit(1)
		return nil
	}

	for range MAX_PASSWORD_RETRIES {
		fmt.Print(scolor.AddMod("Enter with your master key: ", scolor.Bold))
		password, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			ansi.Red.FgPrintln("Error:", err)
			os.Exit(1)
		}

		is, err := IsMasterKeyHash(db, strings.TrimSpace(string(password)))
		if err != nil {
			ansi.Red.FgPrintln("Error:", err)
			os.Exit(1)
		}

		if !is {
			ansi.Red.FgPrintln("Input doesn't match master key.")
			continue
		} else {
			ansi.Green.FgPrintln("Authenticated.")
			hashed_mk := storage.HashMasterKey(string(password))

			err = saveTmpHash(hashed_mk)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			return hashed_mk
		}
	}

	ansi.Red.FgPrintln("Could not authenticate, aborting.")
	os.Exit(1)
	return nil
}

// CheckArgumentSpelling checks the spelling of the optional third argument, it checks if
// argument, when not a exact match, 'looks' like other Secrets on the DB, if it is it returns
// the instatieted Secret.
func CheckArgumentSpelling(args []string, db storage.DB) (storage.Secret, error) {
	query := strings.ToLower(args[2])
	all_secrets, err := db.List()
	if err != nil {
		return storage.Secret{}, err
	}

	var probable_secret int
	var compare_count int
	for index, secret := range all_secrets {
		if strings.Contains(query, strings.ToLower(secret.Name)) {
			return all_secrets[index], nil
		}

		var this_count int
		for _, c := range strings.ToLower(secret.Name) {
			if strings.ContainsRune(query, c) {
				this_count++
			}
		}

		if this_count > compare_count && this_count > 0 {
			compare_count = this_count
			probable_secret = index
		}
	}

	if compare_count <= 0 {
		return storage.Secret{}, fmt.Errorf("query wasn't close to anything in db")
	}

	return all_secrets[probable_secret], nil
}

// sessionFile() function returns the path to a temporary session file based on the ppid
// of the shell process, the intend is that once you authenticate in a shell process you
// won't need to do it again till you create a new shell process.
func sessionFile() string {
	ppid := os.Getppid()
	return fmt.Sprintf("/tmp/agopass_%d", ppid)
}

// saveTmpHash writes your hashed and encrypted master key in order to do the auto auth,
// it stays on the /tmp/ so all of this files are deleted on shutdown.
func saveTmpHash(hash []byte) error {
	return os.WriteFile(sessionFile(), hash, 0600)
}

// loadTmpHash tries to load the session file based on the ppid of the actual shell process
// if the file could not be found it returns an error, if the file was found, returns the
// hashed master key.
func loadTmpHash() ([]byte, error) {
	data, err := os.ReadFile(sessionFile())
	if err != nil {
		return nil, err
	}

	return data, nil
}

/*
Index:
func PrintUsage(args []string)
func CheckAmountArguments(args []string)
func IsMasterKeyHash(db storage.DB, s string) (bool, error
func Authenticate(db storage.DB) []byte
func CheckArgumentSpelling(args []string, db storage.DB) (storage.Secret, error)
func sessionFile() string
func saveTmpHash(hash []byte) error
func loadTmpHash() ([]byte, error)
*/
