package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/Alvesafk/gopass/color"
)

func List() {
	if !fileExists(config_path) {
		fmt.Print(color.Red("Secrets file does not exist, use <gopass init>, exiting.", "bold", 1))
		return
	}

	var all_secrets []Secret

	data, err := os.ReadFile(secrets_path)
	if err != nil {
		log.Fatal(color.Red("Could not read the secrets file.", "bold", 1))
	}

	json.Unmarshal(data, &all_secrets)
	
	for i, v := range all_secrets {
		name := color.Green(v.Name, "bold", 0)

		fmt.Println(color.Cyan("-------------------------------------------------------------------------------", "none", 0))
		fmt.Printf("%v - %s\n    %s\n", i + 1, name, hidePassword(v.Key))
	}
}

func hidePassword(pass string) string {
	var result string
	l := int(math.Min(float64(len(pass)), 25))

	for i := 0; i < l; i++ {
		result += "*"
	}

	if len(pass) > 25 {
		result = result + "..."
	}

	return result
}
