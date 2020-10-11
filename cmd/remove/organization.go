package remove

import (
	"github.com/ddexterpark/merakictl/api/general/organizations"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)

// Delete an existing organization
var org = &cobra.Command{
	Use:   "org",
	Short: "",
	Long: `delete org -o {ORG_ID}`,
	Run: func(cmd *cobra.Command, args []string) {
		org := cmd.Flag("organization").Value.String()
		organization, traceback := organizations.DeleteOrganization(org)
		shell.Display(organization, traceback, "organization", cmd.Flags())
	},
}