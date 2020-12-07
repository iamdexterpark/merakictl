package sm

import "github.com/spf13/cobra"

// DeviceCmd - Root for all organization CLI commands.
var SMCmd = &cobra.Command{
	Use:   "insight",
	Short: "",
	Long:  "Entrypoint for Insight subcommands.",
}

// init - Entrypoint for DeviceCmd sub-commands.
func init() {
	SMCmd.AddCommand(apnscert)
	SMCmd.AddCommand(bypassactivationlockattempts)
	SMCmd.AddCommand(certs)
	SMCmd.AddCommand(deviceprofiles)
	SMCmd.AddCommand(networkadapters)
	SMCmd.AddCommand(restrictions)
	SMCmd.AddCommand(securitycenters)
	SMCmd.AddCommand(devicesoftware)
	SMCmd.AddCommand(wlanlists)
	SMCmd.AddCommand(devices)
	SMCmd.AddCommand(profiles)
	SMCmd.AddCommand(targetgroups)
	SMCmd.AddCommand(targetgroup)
	SMCmd.AddCommand(deviceprofiles)
	SMCmd.AddCommand(users)
	SMCmd.AddCommand(usersoftware)
	SMCmd.AddCommand(vppaccount)
	SMCmd.AddCommand(vppaccounts)
	SMCmd.AddCommand(cellularusagehistory)
	SMCmd.AddCommand(connectivity)
	SMCmd.AddCommand(desktoplogs)
	SMCmd.AddCommand(devicecommandlogs)
	SMCmd.AddCommand(performancehistory)
}