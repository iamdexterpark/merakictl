package appliance

import "github.com/spf13/cobra"

var MXCmd = &cobra.Command{
	Use:   "mx",
	Short: "",
	Long:  "Entrypoint for mx subcommands.",
}

// init - Entrypoint for DeviceCmd sub-commands.
func init() {
	MXCmd.AddCommand(connectivitymonitoringdestinations)
	MXCmd.AddCommand(contentfilteringcategories)
	MXCmd.AddCommand(contentfiltering)
	MXCmd.AddCommand(cellularfirewallrules)
	MXCmd.AddCommand(firewalledservices)
	MXCmd.AddCommand(firewalledservice)
	MXCmd.AddCommand(inboundfirewallrules)
	MXCmd.AddCommand(l3firewallrules)
	MXCmd.AddCommand(l7applicationcategories)
	MXCmd.AddCommand(l7firewallrules)
	MXCmd.AddCommand(onetomanynatrules)
	MXCmd.AddCommand(onetoonenatrules)
	MXCmd.AddCommand(portforwardingrules)
	MXCmd.AddCommand(ports)
	MXCmd.AddCommand(port)
	MXCmd.AddCommand(networkintrusion)
	MXCmd.AddCommand(organizationintrusion)
	MXCmd.AddCommand(malware)
	MXCmd.AddCommand(settings)
	MXCmd.AddCommand(singlelan)
	MXCmd.AddCommand(staticroutes)
	MXCmd.AddCommand(staticroute)
	MXCmd.AddCommand(customperformanceclasses)
	MXCmd.AddCommand(customperformanceclass)
	MXCmd.AddCommand(trafficshapingrules)
	MXCmd.AddCommand(uplinkbandwidth)
	MXCmd.AddCommand(uplinkselection)
	MXCmd.AddCommand(trafficshaping)
	MXCmd.AddCommand(vlansettings)
	MXCmd.AddCommand(vlans)
	MXCmd.AddCommand(vlan)
	MXCmd.AddCommand(sitetositevpn)
	MXCmd.AddCommand(thirdpartyvpnpeers)
	MXCmd.AddCommand(vpnfirewallrules)
	MXCmd.AddCommand(warmspare)
	MXCmd.AddCommand(securityevents)
	MXCmd.AddCommand(dhcpsubnets)
	MXCmd.AddCommand(performance)
	MXCmd.AddCommand(networksecurityevents)
	MXCmd.AddCommand(organizationsecurityevents)
	MXCmd.AddCommand(uplinkstatuses)
	MXCmd.AddCommand(vpnstats)
	MXCmd.AddCommand(vpnstatuses)
}
