package create

import (
	"github.com/ddexterpark/merakictl/api/general/organizations/configure"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)

// Create a new organization
var org = &cobra.Command{
	Use:   "org",
	Short: "",
	Long:  `create org {NAME}`,
	Run: func(cmd *cobra.Command, args []string) {
		var name = args[0]
		metadata := configure.CreateOrganization(name)
		shell.Display(metadata, "organization", cmd.Flags())
	},
}
