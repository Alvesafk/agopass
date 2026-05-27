package cmd

import (
	"fmt"
	"log"
	"os"
)

var (
	user_home = getUserHomeDir()
	config_path = user_home + "/.config/gopass"
)

func Init() {
	fmt.Printf("\033[1m\033[32mInitiating, checking if config dir already exists.\033[0m\n")
	if fileExists(config_path) {
		fmt.Printf("\033[1m\033[32mFile already exists!\033[0m\n")
		return
	} else {
		fmt.Printf("\033[33mDir does not exist, creating it.\033[0m\n")
		err := createConfigDir()
		if err != nil {
			fmt.Println(err)
		}

		if fileExists(config_path) {
			fmt.Printf("\033[1m\033[32mSuccess! Config dir was created\033[0m\n")
		} else {
			log.Fatal("File was not created! A problem ocurred.")
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

func getUserHomeDir() string {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	return home_dir
}

