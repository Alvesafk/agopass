/*
List command shows all secrets that are registered on the DB, the keys that it shows are
just a mock based on the lenght of the plain text key, therefore the user does not need to
authenticate before using it.
*/
package cmd

import (
	"fmt"
	"math"

	"github.com/Alvesafk/agopass/color"
	"github.com/Alvesafk/agopass/storage"
)

// List function, accepts a DB connection.
func List(db storage.DB) {
	// db.List() method does not have any argument (only the implicit DB one), the 
	// method returns a slice of initialized Secret structs.
	all_secrets, err := db.List()
	if err != nil {
		fmt.Print(color.Red("Could not list secrets from db", "bold", 1))
		return
	}

	// If no secret exist it will tell the user to add them.
	if len(all_secrets) < 1 {
		fmt.Print(color.Red("No secret registered! Use <gopass add> to add secrets.", "bold", 1))
		return
	}

	fmt.Println("---------------~Secrets~---------------")

	// Range over the slice of secrets, printing them into your terminal.
	for _, v := range all_secrets {
		fmt.Printf("Name: %s\n", v.Name)
		fmt.Printf("Key:  %s\n", hidePassword(v.Key_Length))
		fmt.Println("---------------------------------------")
	}
}

// hide password function accepts a integer representing the lenght of the password, the
// return is a string made of '*'. Ex.: Pass123 -> len = 7 -> *******
func hidePassword(kl int) string {
	// Initialize the result string.
	var result string

	// Get the smallest of this two, the limit length is 25 characters.
	l := int(math.Min(float64(kl), 25))

	// Range over l(int) adding a '*' onto the result string in every iteration.
	for range l {
		result += "*"
	}

	// If the kl(key length) it's greater than 25 add a "..." on the end of the string.
	if kl > 25 {
		result = result + "..."
	}

	return result
}

/*
Index:
func List(db storage.DB)
func hidePassword(kl int) string
*/
