package update

import (
	"github.com/ddexterpark/merakictl/api/general/organizations/configure"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)

// Update an existing organization
var org = &cobra.Command{
	Use:   "org",
	Short: "",
	Long: `update org {NAME} -o {ORG_ID}`,
	Run: func(cmd *cobra.Command, args []string) {

		org := cmd.Flag("organization").Value.String()
		var name = args[0]

		organization, traceback := configure.UpdateOrganization(org, name)
		shell.Display(organization, traceback, "organization", cmd.Flags())
	},
}
