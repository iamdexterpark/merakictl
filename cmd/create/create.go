package create

import (
	"github.com/spf13/cobra"
)

// CreateCmd
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "",
	Long:  "Base for all HTTP POST Calls to Meraki Dashboard.",
}

func init() {
	CreateCmd.AddCommand(org)
}
