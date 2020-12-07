package cellulargateway


import "github.com/spf13/cobra"

// DeviceCmd - Root for all organization CLI commands.
var MGCmd = &cobra.Command{
	Use:   "mg",
	Short: "",
	Long:  "Entrypoint for cellular gateway subcommands.",
}

// init - Entrypoint for DeviceCmd sub-commands.
func init() {
	MGCmd.AddCommand(connectivitymonitor)
	MGCmd.AddCommand(dhcp)
	MGCmd.AddCommand(lan)
	MGCmd.AddCommand(portforwardingrules)
	MGCmd.AddCommand(subnetpool)
	MGCmd.AddCommand(uplink)

}