package _switch

import "github.com/spf13/cobra"

// DeviceCmd - Root for all organization CLI commands.
var MSCmd = &cobra.Command{
	Use:   "ms",
	Short: "",
	Long:  "Entrypoint for switch subcommands.",
}

// init - Entrypoint for DeviceCmd sub-commands.
func init() {
	MSCmd.AddCommand(accesscontrollists)
	MSCmd.AddCommand(accesspolicies)
	MSCmd.AddCommand(accesspolicy)
	MSCmd.AddCommand(switchportsprofiles)
	MSCmd.AddCommand(switchportprofile)
	MSCmd.AddCommand(switchprofiles)
	MSCmd.AddCommand(dhcpserverpolicy)
	MSCmd.AddCommand(dscp)
	MSCmd.AddCommand(linkaggregations)
	MSCmd.AddCommand(mtu)
	MSCmd.AddCommand(portschedules)
	MSCmd.AddCommand(switchports)
	MSCmd.AddCommand(switchport)
	MSCmd.AddCommand(qosruleids)
	MSCmd.AddCommand(qosrules)
	MSCmd.AddCommand(qosrule)
	MSCmd.AddCommand(dhcp)
	MSCmd.AddCommand(l3interfaces)
	MSCmd.AddCommand(l3interface)
	MSCmd.AddCommand(rendezvouspoints)
	MSCmd.AddCommand(rendezvouspoint)
	MSCmd.AddCommand(multicast)
	MSCmd.AddCommand(ospf)
	MSCmd.AddCommand(staticroutes)
	MSCmd.AddCommand(staticroute)
	MSCmd.AddCommand(settings)
	MSCmd.AddCommand(stackdhcp)
	MSCmd.AddCommand(stackl3interfaces)
	MSCmd.AddCommand(stackl3interface)
	MSCmd.AddCommand(stackstaticroutes)
	MSCmd.AddCommand(stackstaticroute)
	MSCmd.AddCommand(switchstacks)
	MSCmd.AddCommand(switchstack)
	MSCmd.AddCommand(stormcontrol)
	MSCmd.AddCommand(stp)
	MSCmd.AddCommand(warmspare)
	MSCmd.AddCommand(packets)
	MSCmd.AddCommand(portsstatuses)
}