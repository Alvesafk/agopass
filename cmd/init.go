package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Alvesafk/gopass/color"
	"github.com/Alvesafk/agosh/utils"
)

var (
	user_home = utils.GetUserHomeDir()
	config_path = user_home + "/.config/gopass"
	secrets_path = config_path + "/secrets.json"
)

func Init() {
	fmt.Print(color.Green("Initializing, checking if config dir already exists.", "bold", 1))
	if fileExists(config_path) {
		fmt.Print(color.Green("Dir already exists! Exiting.", "bold", 1))

		if fileExists(secrets_path) {
			return
		} else {
			fmt.Print(color.Yellow("Secrets file don't exist! Creating.", "underline", 1))
			err := createConfigFile()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Print(color.Green("Config file was created!", "bold", 2))
			return 
		}

	} else {
		fmt.Print(color.Yellow("Config dir does not exist! Creating.", "underline", 1))
		err := createConfigDir()
		if err != nil {
			fmt.Println(err)
		}

		if fileExists(config_path) {
			fmt.Print(color.Green("Success! Config dir was created.", "bold", 1))
			fmt.Print(color.Yellow("Creating secrets file.", "underline", 1))
			err := createConfigFile()
			if err != nil {
				log.Fatal(err)
			}

			if fileExists(secrets_path) {
				fmt.Print(color.Green("Success! Secrets file was created.", "bold", 1))
			}

		} else {
			log.Fatal(color.Red("Dir was not created! A problem ocurred.", "bold", 1))
		}
	}
	
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func createConfigDir () error {
	err := os.MkdirAll(config_path, 0755)
	if err != nil {
		return fmt.Errorf("Creating config directory: %w", err)
	}

	return nil
}

func createConfigFile () error {
	file, err := os.Create(secrets_path)
	if err != nil {
		return fmt.Errorf("Creating secrets file: %w", err)
	}

	defer file.Close() 

	return nil
}
