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
	OrgCmd.AddCommand(configurationtemplates)
	OrgCmd.AddCommand(configurationtemplate)
	OrgCmd.AddCommand(devices)
	OrgCmd.AddCommand(inventory)
	OrgCmd.AddCommand(licences)
	OrgCmd.AddCommand(licence)
	OrgCmd.AddCommand(securitysettings)
	OrgCmd.AddCommand(networks)
	OrgCmd.AddCommand(ldps)
	OrgCmd.AddCommand(ldp)
	OrgCmd.AddCommand(saml)
	OrgCmd.AddCommand(samlroles)
	OrgCmd.AddCommand(samlrole)
	OrgCmd.AddCommand(snmp)
	OrgCmd.AddCommand(apiresponsecount)
	OrgCmd.AddCommand(apirequests)
	OrgCmd.AddCommand(configurationchanges)
	OrgCmd.AddCommand(devicestatus)
	OrgCmd.AddCommand(losslatency)
	OrgCmd.AddCommand(licenseoverview)
	OrgCmd.AddCommand(openapi)
	OrgCmd.AddCommand(webhooklogs)
}
