package organization

import (
	"github.com/spf13/cobra"
)


// OrgCmd - Root for all organization CLI commands.
var OrgCmd = &cobra.Command{
	Use:   "org",
	Short: "",
	Long: "List of available organization specific API calls exposed via CLI.",
}

// init - Entrypoint for OrgCmd sub-commands.
func init() {
	OrgCmd.AddCommand(orgList)
	OrgCmd.AddCommand(single)
	OrgCmd.AddCommand(networks)
	OrgCmd.AddCommand(devices)
	OrgCmd.AddCommand(actionbatchlist)
	OrgCmd.AddCommand(admins)
	OrgCmd.AddCommand(brandingPolicies)
	OrgCmd.AddCommand(brandingPolicy)
}