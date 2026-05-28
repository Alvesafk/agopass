package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Alvesafk/gopass/color"
)

type Secret struct {
	Name string `json:"name"`
	Key string `json:"key"`
}

func Add() {
	if !fileExists(config_path) {
		fmt.Print(color.Red("Secrets file does not exist, use <gopass init>, exiting.", "bold", 1))
		return
	}

	fmt.Print(color.Green("Add a Secret:", "bold", 1))

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(color.Yellow("Name of the secret: ", "none", 0))
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(color.Yellow("Secret word: ", "none", 0))
	secret, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	new_secret := Secret{strings.TrimSpace(name), strings.TrimSpace(secret)}

	secretsFileStats, err := os.Stat(secrets_path)
	if err != nil {
		fmt.Print(color.Red("Secrets file does not exist, use <gopass init>, exiting.", "bold", 1))
		return
	}

	if secretsFileStats.Size() <= 10 {
		all_secrets := []Secret{new_secret}
		data, err := json.Marshal(all_secrets)
		if err != nil {
			log.Fatal(color.Red("Could not serialize secret into JSON", "bold", 1))
		}

		err = os.WriteFile(secrets_path, data, 0644)
		if err != nil {
			log.Fatal(color.Red("Could not write into file: %s", "bold", 1), err)	
		}
	} else {
		var all_secrets []Secret

		data, err := os.ReadFile(secrets_path)
		if err != nil {
			log.Fatal(color.Red("Could not read the secrets file.", "bold", 1))
		}

		json.Unmarshal(data, &all_secrets)
		all_secrets = append(all_secrets, new_secret)

		writeData, err := json.Marshal(all_secrets)
		if err != nil {
			log.Fatal(color.Red("Could not serialize secret into JSON", "bold", 1))
		}

		err = os.WriteFile(secrets_path, writeData, 0644)
		if err != nil {
			log.Fatal(color.Red("Could not write into file: %s", "bold", 1), err)	
		}
	}

	fmt.Print(color.Green("Secret was saved!", "bold", 1))
}
