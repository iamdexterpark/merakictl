package mg

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/cellulargateway/configure"
	shell "github.com/ddexterpark/merakictl/utilities"
	"github.com/spf13/cobra"
)


var GetConnectivityMonitor = &cobra.Command{
	Use:   "ConnectivityMonitor",
	Short: "Return the connectivity testing destinations for an MG network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetConnectivityMonitoringDestinations(networkId)
		shell.Display(metadata, "ConnectivityMonitor", cmd.Flags())
	},
}

var PutConnectivityMonitor = &cobra.Command{
	Use:   "ConnectivityMonitor",
	Short: "Update the connectivity testing destinations for an MG network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.ConnectivityTesting
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutConnectivityMonitoringDestinations(networkId,  input)
		shell.Display(metadata, "ConnectivityMonitor", cmd.Flags())
	},
}

var GetDHCP = &cobra.Command{
	Use:   "DHCP",
	Short: "List common DHCP settings of MGs.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetDhcp(networkId)
		shell.Display(metadata, "DHCP", cmd.Flags())
	},
}

var PutDHCP = &cobra.Command{
	Use:   "DHCP",
	Short: "List common DHCP settings of MGs.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.DHCP
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutDhcp(networkId,  input)
		shell.Display(metadata, "DHCP", cmd.Flags())
	},
}

var GetLan = &cobra.Command{
	Use:   "Lan",
	Short: "Show the LAN Settings of a MG.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := configure.GetLan(serial)
		shell.Display(metadata, "Lan", cmd.Flags())
	},
}

var PutLan = &cobra.Command{
	Use:   "Lan",
	Short: "Show the LAN Settings of a MG.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		var format configure.LAN
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutLan(serial,  input)
		shell.Display(metadata, "Lan", cmd.Flags())
	},
}

var GetPortForwardingRules = &cobra.Command{
	Use:   "PortForwardingRules",
	Short: "Returns the port forwarding rules for a single MG.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := configure.GetPortForwardingRules(serial)
		shell.Display(metadata, "PortForwardingRules", cmd.Flags())
	},
}

var PutPortForwardingRules = &cobra.Command{
	Use:   "PortForwardingRules",
	Short: "Returns the port forwarding rules for a single MG.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		var format configure.PortForwardingRules
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutPortForwardingRules(serial,  input)
		shell.Display(metadata, "PortForwardingRules", cmd.Flags())
	},
}

var GetSubnetPool = &cobra.Command{
	Use:   "SubnetPool",
	Short: "Return the subnet pool and mask configured for MGs in the network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetSubnetPool(networkId)
		shell.Display(metadata, "SubnetPool", cmd.Flags())
	},
}

var PutSubnetPool = &cobra.Command{
	Use:   "SubnetPool",
	Short: "Return the subnet pool and mask configured for MGs in the network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.SubnetPool
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutSubnetPool(networkId,  input)
		shell.Display(metadata, "SubnetPool", cmd.Flags())
	},
}

var GetUplink = &cobra.Command{
	Use:   "Uplink",
	Short: "Returns the uplink settings for your MG network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetUplink(networkId)
		shell.Display(metadata, "Uplink", cmd.Flags())
	},
}

var PutUplink = &cobra.Command{
	Use:   "Uplink",
	Short: "Returns the uplink settings for your MG network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.Uplink
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutUplink(networkId,  input)
		shell.Display(metadata, "Uplink", cmd.Flags())
	},
}