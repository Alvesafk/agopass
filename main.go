/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>

Agopass (formerly gopass but someone had already stole my flow) is a CLI Password manager
made in Go, it uses a SQLite db located on this path /home/<user>/.agopass/, all the
contentes of the DB are hashed or encrypted, it's simple to use, few commands but it gets
the job done, to start with agopass you have to run <agopass init>, that will create the
DB and prompt you to create your Master Password, the only password you'l need to remember!
*/
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Alvesafk/agopass/cmd"
	"github.com/Alvesafk/agopass/storage"

	tea "charm.land/bubbletea/v2"
)

var (
	commands = []string{"Add secret", "List secrets", "Get secret", "Update secret", "Make secret"}
	args     = os.Args
)

type choice_model struct {
	choices  []string
	cursor   int
	selected string

	db storage.DB
}

func InitModel() choice_model {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	db_path := filepath.Join(home, ".agopass", "secrets.db")
	err = os.MkdirAll(filepath.Dir(db_path), 0755)
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.New(db_path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = db.Close(); err != nil {
			log.Println(err)
		}
	}()

	return choice_model{

		choices: commands,
		db:      *db,
	}
}

func (cm choice_model) Init() tea.Cmd {
	return nil
}

func (cm choice_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return cm, tea.Quit

		case "up", "k":
			if cm.cursor > 0 {
				cm.cursor--
			}

		case "down", "j":
			if cm.cursor < len(cm.choices)-1 {
				cm.cursor++
			}

		case "enter", "space":
			cm.selected = cm.choices[cm.cursor]
		}
	}

	if len(cm.selected) > 0 {
		switch cm.selected {
		case commands[0]:
			cmd.Add(cm.db, args)
		case commands[1]:
			cmd.List(cm.db)
		case commands[2]:
			cmd.Get(cm.db, args)
		case commands[3]:
			cmd.Update(cm.db, args)
		case commands[4]:
			cmd.Make()
		}
	}

	return cm, nil
}

func (cm choice_model) View() tea.View {
	s := "Agopass secrets manager\n\n"

	for i, choice := range cm.choices {
		cursor := " "
		if cm.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"

	return tea.NewView(s)
}

func main() {
	p := tea.NewProgram(InitModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
