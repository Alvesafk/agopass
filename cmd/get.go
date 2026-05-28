package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Alvesafk/gopass/color"
	"github.com/atotto/clipboard"
)

func Get(args []string) {
	if !fileExists(config_path) {
		fmt.Print(color.Red("Secrets file does not exist, use <gopass init>, exiting.", "bold", 1))
		return
	}
	
	l := len(args)
	if l != 3 {
		if l > 3 {
			fmt.Print(color.Red("Error: Too many arguments.", "bold", 1))
			PrintUsage(args)
			return
		}

		if l < 3 {
			fmt.Print(color.Red("Error: Missing arguments.", "bold", 1))
			PrintUsage(args)
			return
		}
	}

	to_get_name := strings.ToLower(strings.TrimSpace(args[2]))

	var all_secrets []Secret

	data, err := os.ReadFile(secrets_path)
	if err != nil {
		fmt.Print(color.Red("Could not read the secrets file.", "bold", 1))
	}

	json.Unmarshal(data, &all_secrets)

	var to_get_secret Secret

	for _, v := range all_secrets {
		if strings.ToLower(v.Name) == to_get_name {
			to_get_secret = v
		}
	}

	if to_get_secret.Key == "" && to_get_secret.Name == "" {
		fmt.Print(color.Red("Could not find this key, check your typing.", "bold", 1))
		return 
	}

	err = clipboard.WriteAll(to_get_secret.Key)
	if err != nil {
		panic(err)
	}

	fmt.Printf(color.Green("Success! %s key is in your clipboard.", "bold", 1), to_get_secret.Name)
}

