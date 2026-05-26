/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
	Use:   "gopass",
	Short: "Password CLI manager made in Go",
	Long: `gopass
Password CLI manager. Usage:
tired to do this shit
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
