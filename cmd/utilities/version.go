package utilities

import (
	"fmt"
	"github.com/spf13/cobra"
)

// version
var Version = &cobra.Command{
	Use:   "version",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Merakictl Alpha Version")

	},
}
