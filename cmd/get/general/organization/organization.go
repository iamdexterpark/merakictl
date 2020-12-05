package organization

import (
	"github.com/spf13/cobra"
)

// OrgCmd - Root for all organization CLI commands.
var OrgCmd = &cobra.Command{
	Use:   "org",
	Short: "organization level commands.",
	Long:  "List of available organization specific API calls exposed via CLI.",
}

// init - Entrypoint for OrgCmd sub-commands.
func init() {
	OrgCmd.AddCommand(list)
	OrgCmd.AddCommand(detail)
	OrgCmd.AddCommand(actionbatches)
	OrgCmd.AddCommand(actionbatch)
	OrgCmd.AddCommand(admins)
	OrgCmd.AddCommand(brandingPolicies)
	OrgCmd.AddCommand(brandingPolicy)
	OrgCmd.AddCommand(networks)
	OrgCmd.AddCommand(devices)
}
