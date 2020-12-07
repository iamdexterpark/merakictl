package wireless

import "github.com/spf13/cobra"

// MRCmd - Root for all organization CLI commands.
var MRCmd = &cobra.Command{
	Use:   "mr",
	Short: "",
	Long:  "List of available wifi specific API calls exposed via CLI.",
}

// init - Entrypoint for OrgCmd sub-commands.
func init() {
	MRCmd.AddCommand(alternatemgmtinterface)
	MRCmd.AddCommand(bluetoothsettings)
	MRCmd.AddCommand(bluetoothdevicesettings)
	MRCmd.AddCommand(radiosettings)
	MRCmd.AddCommand(rfprofiles)
	MRCmd.AddCommand(rfprofile)
	MRCmd.AddCommand(wirelesssettings)
	MRCmd.AddCommand(l3FirewallRules)
	MRCmd.AddCommand(l7FirewallRules)
	MRCmd.AddCommand(identitypsks)
	MRCmd.AddCommand(identitypsk)
	MRCmd.AddCommand(splashsettings)
	MRCmd.AddCommand(trafficshapingrules)
	MRCmd.AddCommand(ssids)
	MRCmd.AddCommand(ssid)
	MRCmd.AddCommand(airmarshal)
	MRCmd.AddCommand(channelutilizationhistory)
	MRCmd.AddCommand(clientcounthistory)
	MRCmd.AddCommand(connectionstat)
	MRCmd.AddCommand(connectionstats)
	MRCmd.AddCommand(connectivityevents)
	MRCmd.AddCommand(latencyhistory)
	MRCmd.AddCommand(latencystat)
	MRCmd.AddCommand(latencystats)
	MRCmd.AddCommand(deviceconnectionstats)
	MRCmd.AddCommand(networkconnectionstats)
	MRCmd.AddCommand(dataratehistory)
	MRCmd.AddCommand(connectionstats)
	MRCmd.AddCommand(latencystats)
	MRCmd.AddCommand(failedconnections)
	MRCmd.AddCommand(latencyhistory)
	MRCmd.AddCommand(devicelatencystats)
	MRCmd.AddCommand(networklatencystats)
	MRCmd.AddCommand(meshstatuses)
	MRCmd.AddCommand(signalqualityhistory)
	MRCmd.AddCommand(status)
	MRCmd.AddCommand(usagehistory)
}
