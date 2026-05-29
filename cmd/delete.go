package cmd

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/Alvesafk/gopass/color"
)

func Delete(args []string) {
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

	index := math.MinInt

	for i, v := range all_secrets {
		if strings.ToLower(v.Name) == to_get_name {
			index = i
		}
	}

	if index == math.MinInt {
		fmt.Print(color.Red("Could not find this key, check your typing.", "bold", 1))
		return 
	}

	all_secrets = RemoveSecret(all_secrets, index)

	writeData, err := json.Marshal(all_secrets)
	if err != nil {
		fmt.Print(color.Red("Could not serialize secret into JSON", "bold", 1))
		return
	}

	err = os.WriteFile(secrets_path, writeData, 0644)
	if err != nil {
		fmt.Print(color.Red("Could not write into file: %s", "bold", 1), err)	
		return
	}
}

