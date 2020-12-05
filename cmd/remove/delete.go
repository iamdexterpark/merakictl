package remove

import (
	"github.com/ddexterpark/merakictl/cmd/remove/general/organization"
	"github.com/spf13/cobra"
)

// DeleteCmd
var DeleteCmd = &cobra.Command{
	Use:   "remove",
	Short: "",
	Long:  "Base for all HTTP DEL Calls to Meraki Dashboard.",
}

func init() {
	DeleteCmd.AddCommand(organization.Org)
}
