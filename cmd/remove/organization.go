package remove

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/organizations/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

// Delete an existing organization
var org = &cobra.Command{
	Use:   "org",
	Short: "",
	Long:  `remove org -o {ORG_ID}`,
	Run: func(cmd *cobra.Command, args []string) {
		org := cmd.Flag("organization").Value.String()
		metadata := configure.DeleteOrganization(org)
		shell.Display(metadata, "organization", cmd.Flags())
	},
}
