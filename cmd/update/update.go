package update

import (
	"github.com/ddexterpark/merakictl/cmd/update/wireless"
	"github.com/spf13/cobra"
)

// UpdateCmd
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates configurations.",
	Long:  "Base for all HTTP PUT Calls to Meraki Dashboard.",
}

func init() {
	UpdateCmd.AddCommand(org)
	UpdateCmd.AddCommand(wireless.MRCmd)
}
