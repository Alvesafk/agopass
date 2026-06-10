package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"

	"github.com/Alvesafk/agopass/color"
	"github.com/atotto/clipboard"
)

const (
	CHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*(),.<>;:/?{}[]"
	DEFAULT_LENGHT = 32
)

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

