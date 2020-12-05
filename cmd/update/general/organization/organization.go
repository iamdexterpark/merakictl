package organization

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/organizations/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

// Update an existing organization
var Org = &cobra.Command{
	Use:   "org",
	Short: "",
	Long:  `update org {NAME} -o {ORG_ID}`,
	Run: func(cmd *cobra.Command, args []string) {

		org := cmd.Flag("organization").Value.String()
		var name = args[0]

		metadata := configure.UpdateOrganization(org, name)
		shell.Display(metadata, "organization", cmd.Flags())
	},
}
