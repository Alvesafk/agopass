package cmd

import (
	"fmt"

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
}

