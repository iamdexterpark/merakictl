package mg

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/cellulargateway/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
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
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutConnectivityMonitoringDestinations(networkId, format)
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
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutDhcp(networkId, format)
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
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutLan(serial, format)
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
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutPortForwardingRules(serial, format)
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
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutSubnetPool(networkId, format)
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
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutUplink(networkId, format)
		shell.Display(metadata, "Uplink", cmd.Flags())
	},
}