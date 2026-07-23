/*
The Make command creates a random N (where N can be a user inputed value or the default
16 one) bit long key, if you compile it from source you can change the length, will add
this eventually, the command uses the crypto/rand package to generate random numbers,
crypto is used because it generates ranmdom numbers using a lot of entropy, making the
'randomnes' true, is not biased as well.
*/
package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/Alvesafk/scolor"
	"github.com/Alvesafk/scolor/ansi"
	"github.com/atotto/clipboard"
)

// consts, CHARS is the string used to generate the random key, DEFAULT_LENGTH is the default
// len of the generated key.
const (
	CHARS          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*(),.<>;:/?{}[]"
	DEFAULT_LENGTH = 16
)

// Make function receives nothing and return nothing, it generate the random key and paste
// it on your clipboard.
func Make(args []string) {
	fmt.Println(scolor.AddMod("Making a new random key, putting on your clipboard.", scolor.Bold))

	makeLen := DEFAULT_LENGTH
	if len(args) > 2 {
		temp, err := strconv.Atoi(args[2])
		if err == nil {
			makeLen = temp
		} else {
			ansi.Yellow.FgPrintf("Not a valid number. Using default length %v..", DEFAULT_LENGTH)
		}
	}

	var sb strings.Builder
	sb.Grow(makeLen)

	for range makeLen {
		i, _ := rand.Int(rand.Reader, big.NewInt(int64(len(CHARS))))

		sb.WriteByte(CHARS[i.Int64()])
	}

	result := sb.String()

	err := clipboard.WriteAll(result)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ansi.Green.FgPrintln("Success! New random key is in your clipboard.")
}

/*
const CHARS string
const DEFAULT_LENGTH int
func Make()
*/
