package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// version
var version = &cobra.Command{
	Use:   "version",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Merakictl Alpha Version")

	},
}