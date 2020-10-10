package remove

import (
	"github.com/spf13/cobra"
)

// DeleteCmd
var DeleteCmd = &cobra.Command{
	Use:   "remove",
	Short: "",
	Long: "Base for all HTTP DEL Calls to Meraki Dashboard.",
}

func init() {
	DeleteCmd.AddCommand(org)
}
