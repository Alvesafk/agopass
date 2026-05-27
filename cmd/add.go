package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Alvesafk/gopass/format"
)

type SecretFile struct {
	Secrets []Secret
}

type Secret struct {
	Name string `json:"name"`
	Key string `json:"key"`
}

func Add() {
	if !fileExists(config_path) {
		fmt.Print(format.Red("Secrets file does not exist, use <gopass init>, exiting.", "bold", 1))
		return
	}

	fmt.Print(format.Green("Add a Secret:", "bold", 1))

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(format.Yellow("Name of the secret: ", "none", 0))

	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Print(format.Yellow("Secret word: ", "none", 0))

	secret, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	new_secret := Secret{name, secret}

	bytes, err := json.Marshal(new_secret)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))
}

