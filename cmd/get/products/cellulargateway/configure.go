package cellulargateway

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/cellulargateway/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)


var connectivitymonitor = &cobra.Command{
	Use:   "connectivitymonitor",
	Short: "Return the connectivity testing destinations for an MG network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetConnectivityTesting(networkId)
		shell.Display(metadata, "connectivitymonitor", cmd.Flags())
	},
}

var dhcp = &cobra.Command{
	Use:   "dhcp",
	Short: "List common DHCP settings of MGs.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetDHCPSettings(networkId)
		shell.Display(metadata, "dhcp", cmd.Flags())
	},
}

var lan = &cobra.Command{
	Use:   "lan",
	Short: "Show the LAN Settings of a MG.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := configure.GetLANSettings(serial)
		shell.Display(metadata, "lan", cmd.Flags())
	},
}

var portforwardingrules = &cobra.Command{
	Use:   "portforwardingrules",
	Short: "Returns the port forwarding rules for a single MG.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := configure.GetPortForwardingRules(serial)
		shell.Display(metadata, "portforwardingrules", cmd.Flags())
	},
}

var subnetpool = &cobra.Command{
	Use:   "subnetpool",
	Short: "Return the subnet pool and mask configured for MGs in the network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetSubnetPool(networkId)
		shell.Display(metadata, "subnetpool", cmd.Flags())
	},
}

var uplink = &cobra.Command{
	Use:   "uplink",
	Short: "Returns the uplink settings for your MG network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetUplinkSettings(networkId)
		shell.Display(metadata, "uplink", cmd.Flags())
	},
}