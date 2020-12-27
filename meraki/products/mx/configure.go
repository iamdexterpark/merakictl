package mx

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/appliance/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)


var GetConnectivityMonitoringDestinations = &cobra.Command{
Use:   "ConnectivityMonitoringDestinations",
Short: "Return the connectivity testing destinations for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetConnectivityMonitoringDestinations(networkId)
shell.Display(metadata, "ConnectivityMonitoringDestinations", cmd.Flags())
},
}



var GetContentFilteringCategories = &cobra.Command{
Use:   "ContentFilteringCategories",
Short: "List all available content filtering categories for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetContentFilteringCategories(networkId)
shell.Display(metadata, "ContentFilteringCategories", cmd.Flags())
},
}



var GetContentFiltering = &cobra.Command{
Use:   "ContentFiltering",
Short: "Return the content filtering settings for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetContentFiltering(networkId)
shell.Display(metadata, "ContentFiltering", cmd.Flags())
},
}



var GetCellularFirewallRules = &cobra.Command{
Use:   "CellularFirewallRules",
Short: "Return the cellular firewall rules for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetCellularFirewallRules(networkId)
shell.Display(metadata, "CellularFirewallRules", cmd.Flags())
},
}



var GetFirewalledServices = &cobra.Command{
Use:   "FirewalledServices",
Short: "List the mx services and their accessibility rules.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetFirewalledServices(networkId)
shell.Display(metadata, "FirewalledServices", cmd.Flags())
},
}



var GetFirewalledService = &cobra.Command{
	Use:   "FirewalledService",
	Short: "Return the accessibility settings of the given service ('ICMP', 'web', or 'SNMP').",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		serviceId := args[0]
		metadata := configure.GetFirewalledService(networkId, serviceId)
		shell.Display(metadata, "FirewalledService", cmd.Flags())
	},
}



var GetInboundFirewallRules = &cobra.Command{
Use:   "InboundFirewallRules",
Short: "Return the inbound firewall rules for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetInboundFirewallRules(networkId)
shell.Display(metadata, "InboundFirewallRules", cmd.Flags())
},
}



var GetL3FirewallRules = &cobra.Command{
Use:   "L3FirewallRules",
Short: "Return the L3 firewall rules for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetL3FirewallRules(networkId)
shell.Display(metadata, "L3FirewallRules", cmd.Flags())
},
}



var GetL7ApplicationCategories = &cobra.Command{
Use:   "L7ApplicationCategories",
Short: "Return the L7 firewall application categories and their associated applications for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetL7FirewallApplicationCategories(networkId)
shell.Display(metadata, "L7ApplicationCategories", cmd.Flags())
},
}



var GetL7FirewallRules = &cobra.Command{
Use:   "L7FirewallRule",
Short: "List the MX L7 firewall rules for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetL7FirewallRules(networkId)
shell.Display(metadata, "L7FirewallRule", cmd.Flags())
},
}



var GetOneToManyNatRules = &cobra.Command{
Use:   "OneToManyNatRules",
Short: "Return the 1:Many NAT mapping rules for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetOneToManyNatRules(networkId)
shell.Display(metadata, "OneToManyNatRules", cmd.Flags())
},
}



var GetOneToOneNatRules = &cobra.Command{
Use:   "OneToOneNatRules",
Short: "Return the 1:1 NAT mapping rules for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetOneToOneNatRules(networkId)
shell.Display(metadata, "OneToOneNatRules", cmd.Flags())
},
}



var GetPortForwardingRules = &cobra.Command{
Use:   "PortForwardingRules",
Short: "Return the port forwarding rules for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetPortForwardingRules(networkId)
shell.Display(metadata, "PortForwardingRules", cmd.Flags())
},
}



var GetPorts = &cobra.Command{
Use:   "Ports",
Short: "List per-port VLAN settings for all ports of a MX.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetAppliancePorts(networkId)
shell.Display(metadata, "Ports", cmd.Flags())
},
}

var GetPort = &cobra.Command{
	Use:   "Port",
	Short: "\tReturn per-port VLAN settings for a single MX port.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		portId := args[0]
		metadata := configure.GetAppliancePort(networkId, portId)
		shell.Display(metadata, "Port", cmd.Flags())
	},
}


var GetNetworkIntrusion = &cobra.Command{
Use:   "NetworkIntrusion",
Short: "Returns all supported intrusion settings for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetNetworkSecurityIntrusion(networkId)
shell.Display(metadata, "NetworkIntrusion", cmd.Flags())
},
}



var GetOrganizationIntrusion = &cobra.Command{
Use:   "OrganizationIntrusion",
Short: "Returns all supported intrusion settings for an organization.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetOrganizationSecurityIntrusion(networkId)
shell.Display(metadata, "OrganizationIntrusion", cmd.Flags())
},
}



var GetMalware = &cobra.Command{
Use:   "Malware",
Short: "Returns all supported malware settings for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetMalwareSettings(networkId)
shell.Display(metadata, "Malware", cmd.Flags())
},
}



var GetSettings = &cobra.Command{
Use:   "Settings",
Short: "Return the mx settings for a network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetSettings(networkId)
shell.Display(metadata, "Settings", cmd.Flags())
},
}



var GetSingleLan = &cobra.Command{
Use:   "SingleLan",
Short: "Return single LAN configuration.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetSingleLan(networkId)
shell.Display(metadata, "SingleLan", cmd.Flags())
},
}



var GetStaticRoutes = &cobra.Command{
Use:   "StaticRoutes",
Short: "List the static routes for an MX or teleworker network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetStaticRoutes(networkId)
shell.Display(metadata, "StaticRoutes", cmd.Flags())
},
}


var GetStaticRoute = &cobra.Command{
	Use:   "StaticRoute",
	Short: "\tReturn a static route for an MX or teleworker network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		staticRouteId := args[0]
		metadata := configure.GetStaticRoute(networkId, staticRouteId)
		shell.Display(metadata, "StaticRoute", cmd.Flags())
	},
}


var GetCustomPerformanceClasses = &cobra.Command{
Use:   "CustomPerformanceClasses",
Short: "List all custom performance classes for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetPerformanceClasses(networkId)
shell.Display(metadata, "CustomPerformanceClasses", cmd.Flags())
},
}

var GetCustomPerformanceClass = &cobra.Command{
	Use:   "CustomPerformanceClass",
	Short: "Return a custom performance class for an MX network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		customPerformanceClassId := args[0]
		metadata := configure.GetPerformanceClass(networkId,
			customPerformanceClassId)
		shell.Display(metadata, "CustomPerformanceClass", cmd.Flags())
	},
}

var GetTrafficShapingRules = &cobra.Command{
Use:   "TrafficShapingRules",
Short: "Display the traffic shaping settings rules for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetTrafficShapingRules(networkId)
shell.Display(metadata, "TrafficShapingRules", cmd.Flags())
},
}

var GetUplinkBandwidth = &cobra.Command{
Use:   "UplinkBandwidth",
Short: "Returns the uplink bandwidth settings for your MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetUplinkBandwidth(networkId)
shell.Display(metadata, "UplinkBandwidth", cmd.Flags())
},
}



var GetUplinkSelection = &cobra.Command{
Use:   "UplinkSelection",
Short: "Show uplink selection settings for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetUplinkSelection(networkId)
shell.Display(metadata, "UplinkSelection", cmd.Flags())
},
}

var GetTrafficShaping = &cobra.Command{
Use:   "TrafficShaping",
Short: "Display the traffic shaping settings for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetTrafficShapingRules(networkId)
shell.Display(metadata, "TrafficShaping", cmd.Flags())
},
}



var GetVLANSettings = &cobra.Command{
Use:   "VLANSettings",
Short: "Returns the enabled status of VLANs for the network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetVLANSettings(networkId)
shell.Display(metadata, "VLANSettings", cmd.Flags())
},
}

var GetVLANs = &cobra.Command{
Use:   "VLANs",
Short: "List the VLANs for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetVLANs(networkId)
shell.Display(metadata, "VLANs", cmd.Flags())
},
}

var GetVLAN = &cobra.Command{
	Use:   "VLAN",
	Short: "Return a VLAN.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		vlanId := args[0]
		metadata := configure.GetVLAN(networkId, vlanId)
		shell.Display(metadata, "VLAN", cmd.Flags())
	},
}

var GetSiteToSiteVPN = &cobra.Command{
Use:   "SiteToSiteVPN",
Short: "Return the site-to-site VPN settings of a network. Only valid for MX networks.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetSiteToSiteVPN(networkId)
shell.Display(metadata, "SiteToSiteVPN", cmd.Flags())
},
}

var GetThirdPartyVPNPeers = &cobra.Command{
Use:   "ThirdPartyVPNPeers",
Short: "Return the third party VPN peers for an organization.",
Run: func(cmd *cobra.Command, args []string) {
orgId, _, _ := shell.ResolveFlags(cmd.Flags())
if orgId == "" {
orgId = args[0]
}

metadata := configure.GetThirdPartyVPNPeers(orgId)
shell.Display(metadata, "ThirdPartyVPNPeers", cmd.Flags())
},
}



var GetVPNFirewallRules = &cobra.Command{
Use:   "VPNFirewallRules",
Short: "Return the firewall rules for an organization's site-to-site VPN.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetCellularFirewallRules(networkId)
shell.Display(metadata, "VPNFirewallRules", cmd.Flags())
},
}

var GetWarmspare = &cobra.Command{
Use:   "Warmspare",
Short: "Return MX warm spare settings.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetWarmSpare(networkId)
shell.Display(metadata, "Warmspare", cmd.Flags())
},
}