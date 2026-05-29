package cmd

import (
	"fmt"
)

func PrintUsage(args []string) {
	fmt.Printf("Usage instruction:\n%s <COMMAND>\nCOMMANDS:\ninit\nadd\nlist\ndelete\nget\n", args[0])
}

