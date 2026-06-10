/*
The Make command creates a random 32 bit long key, if you compile it from source you can
change the lenght, will add this eventually, the command uses the crypto/rand package to
generate random numbers, crypto is used because it generates ranmdom numbers using a lot
of entropy, making the 'randomnes' true, is not biased as well.
*/
package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"

	"github.com/Alvesafk/agopass/color"
	"github.com/atotto/clipboard"
)

// consts, CHARS is the string used to generate the random key, DEFAULT_LENGHT is the len
// of the generated key.
const (
	CHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*(),.<>;:/?{}[]"
	DEFAULT_LENGHT = 32
)

// Make function receives nothing and return nothing, it generate the random key and paste
// it on your clipboard.
func Make() {
	fmt.Print(color.AddMod("Making a new random key, putting on your clipboard.\n", "bold"))

	var result string
	for range DEFAULT_LENGHT {
		i, _ := rand.Int(rand.Reader, big.NewInt(int64(len(CHARS))))

		result += string(CHARS[i.Int64()])
	}

	err := clipboard.WriteAll(result)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print(color.Green("Success! New random key is in your clipboard.", "bold", 1))
}

/*
const CHARS string
const DEFAULT_LENGHT int
func Make()
*/
