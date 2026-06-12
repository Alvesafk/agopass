package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	BASH_COMP_SCRIPT = `
_agopass_completion() {
	local cur prev
	cur="${COMP_WORDS[COMP_CWORD]}"
	prev="${COMP_WORDS[COMP_CWORD-1]}"

	local commands="init add list delete get update make autocomplete"

	if [[ $COMP_CWORD -eq 1 ]]; then
		COMPREPLY=($(compgen -W "$commands" -- "$cur"))
		return
	fi
}
complete -F _agopass_completion agopass
	`

	ZSH_COMP_SCRIPT = `
_agopass_completion() {
	local -a commands
	commands=(
		'init:Create DB and prompt for master key'
		'add:Create a secret in DB'
		'list:List all registered secrets'
		'delete:Delete a secret'
		'get:Get the secret key'
		'update:Modify a registered secret'
		'make:Create a random 32 bit key'
		'auto':Setup autocomplete'
	)

	if [[ $CURRENT -eq 2 ]]; then
		_describe 'command' commands
		return
	fi
}

compdef _agopass_completion agopass
	`
)

func InitAutocomplete() error {
	shell := detectShell()

	var rc_file, script string
	switch shell {
	case "bash":
		rc_file = os.ExpandEnv("$home/.bashrc")
		script = BASH_COMP_SCRIPT
	case "zsh":
		rc_file = os.ExpandEnv("$home/.zshrc")
		script = ZSH_COMP_SCRIPT
	default:
		return fmt.Errorf("%s shell is not suported", shell)
	}

	marker := "# agopass-completion"
	content, err := os.ReadFile(rc_file)
	if err != nil {
		return fmt.Errorf("Error: %s", err)
	}

	if strings.Contains(string(content), marker) {
		fmt.Println("Autocomplete is already configured.")
		return nil
	}

	f, err := os.OpenFile(rc_file, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func (){
		if err = f.Close(); err != nil {
			fmt.Println("Error on closing the rc file.")
		}
	}()

	_, err = fmt.Fprintf(f, "\n%s\n%s\n", marker, script)
	if err != nil {
		return err
	}

	fmt.Printf("Autocomplete was configured on %s\n", rc_file)
	fmt.Println("Source your config or restart your shell to activate it.")
	return nil
}

func detectShell() string {
	shell := os.Getenv("SHELL")
	return filepath.Base(shell)
}

func AutocompleteExists() (bool, error){
	shell := detectShell()

	var rc_file string
	switch shell {
	case "bash":
		rc_file = os.ExpandEnv("$home/.bashrc")
	case "zsh":
		rc_file = os.ExpandEnv("$home/.zshrc")
	default:
		return false, fmt.Errorf("%s shell is not suported", shell)
	}

	marker := "# agopass-completion"
	content, err := os.ReadFile(rc_file)
	if err != nil {
		return false, fmt.Errorf("Error: %s", err)
	}

	if strings.Contains(string(content), marker) {
		fmt.Println("Autocomplete is already configured.")
		return true, nil
	}

	return false, fmt.Errorf("Couldn't find the rc file.")
}
