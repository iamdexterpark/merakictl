package create

import (
	"github.com/ddexterpark/merakictl/api/organizations"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)

// Create a new organization
var org = &cobra.Command{
	Use:   "org",
	Short: "",
	Long: `create org {NAME}`,
	Run: func(cmd *cobra.Command, args []string) {
		var name = args[0]
		organization, traceback := organizations.CreateOrganization(name)
		shell.Display(organization, traceback, "organization", cmd.Flags())
	},
}