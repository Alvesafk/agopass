package cmd

import (
	"fmt"
	"math"

	"github.com/Alvesafk/gopass/color"
	"github.com/Alvesafk/gopass/storage"
)

func List(db storage.DB) {
	all_secrets, err := db.List()
	if err != nil {
		fmt.Print(color.Red("Could not list secrets from db", "bold", 1))
		return
	}

	fmt.Println("---------------~Secrets~---------------")

	for _, v := range all_secrets {
		fmt.Printf("Name: %s\n", v.Name)
		fmt.Printf("Key: %s\n", hidePassword(v.Key))
		fmt.Println("---------------------------------------")
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
