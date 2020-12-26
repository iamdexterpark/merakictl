package appliance

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/appliance/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)


var connectivitymonitoringdestinations = &cobra.Command{
Use:   "connectivitymonitoringdestinations",
Short: "Return the connectivity testing destinations for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetConnectivityMonitoringDestinations(networkId)
shell.Display(metadata, "", cmd.Flags())
},
}



var contentfilteringcategories = &cobra.Command{
Use:   "contentfilteringcategories",
Short: "List all available content filtering categories for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetContentFilteringCategories(networkId)
shell.Display(metadata, "contentfilteringcategories", cmd.Flags())
},
}



var contentfiltering = &cobra.Command{
Use:   "contentfiltering",
Short: "Return the content filtering settings for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetContentFiltering(networkId)
shell.Display(metadata, "contentfiltering", cmd.Flags())
},
}



var cellularfirewallrules = &cobra.Command{
Use:   "cellularfirewallrules",
Short: "Return the cellular firewall rules for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetCellularFirewallRules(networkId)
shell.Display(metadata, "cellularfirewallrules", cmd.Flags())
},
}



var firewalledservices = &cobra.Command{
Use:   "firewalledservices",
Short: "List the appliance services and their accessibility rules.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetFirewalledServices(networkId)
shell.Display(metadata, "firewalledservices", cmd.Flags())
},
}



var firewalledservice = &cobra.Command{
	Use:   "firewalledservice",
	Short: "Return the accessibility settings of the given service ('ICMP', 'web', or 'SNMP').",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		serviceId := args[0]
		metadata := configure.GetFirewalledService(networkId, serviceId)
		shell.Display(metadata, "firewalledservice", cmd.Flags())
	},
}



var inboundfirewallrules = &cobra.Command{
Use:   "inboundfirewallrules",
Short: "Return the inbound firewall rules for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetInboundFirewallRules(networkId)
shell.Display(metadata, "inboundfirewallrules", cmd.Flags())
},
}



var  l3firewallrules = &cobra.Command{
Use:   "l3firewallrules",
Short: "Return the L3 firewall rules for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetL3FirewallRules(networkId)
shell.Display(metadata, "l3firewallrules", cmd.Flags())
},
}



var l7applicationcategories = &cobra.Command{
Use:   "l7applicationcategories",
Short: "Return the L7 firewall application categories and their associated applications for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetL7FirewallApplicationCategories(networkId)
shell.Display(metadata, "l7applicationcategories", cmd.Flags())
},
}



var l7firewallrules = &cobra.Command{
Use:   "l7firewallrules",
Short: "List the MX L7 firewall rules for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetL7FirewallRules(networkId)
shell.Display(metadata, "l7firewallrules", cmd.Flags())
},
}



var onetomanynatrules = &cobra.Command{
Use:   "onetomanynatrules",
Short: "Return the 1:Many NAT mapping rules for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetOneToManyNatRules(networkId)
shell.Display(metadata, "onetomanynatrules", cmd.Flags())
},
}



var onetoonenatrules = &cobra.Command{
Use:   "onetoonenatrules",
Short: "Return the 1:1 NAT mapping rules for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetOneToOneNatRules(networkId)
shell.Display(metadata, "onetoonenatrules", cmd.Flags())
},
}



var portforwardingrules = &cobra.Command{
Use:   "portforwardingrules",
Short: "Return the port forwarding rules for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetPortForwardingRules(networkId)
shell.Display(metadata, "portforwardingrules", cmd.Flags())
},
}



var ports = &cobra.Command{
Use:   "ports",
Short: "List per-port VLAN settings for all ports of a MX.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetAppliancePorts(networkId)
shell.Display(metadata, "ports", cmd.Flags())
},
}

var port = &cobra.Command{
	Use:   "port",
	Short: "\tReturn per-port VLAN settings for a single MX port.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		portId := args[0]
		metadata := configure.GetAppliancePort(networkId, portId)
		shell.Display(metadata, "port", cmd.Flags())
	},
}


var networkintrusion = &cobra.Command{
Use:   "networkintrusion",
Short: "Returns all supported intrusion settings for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetNetworkSecurityIntrusion(networkId)
shell.Display(metadata, "networkintrusion", cmd.Flags())
},
}



var organizationintrusion = &cobra.Command{
Use:   "organizationintrusion",
Short: "Returns all supported intrusion settings for an organization.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetOrganizationSecurityIntrusion(networkId)
shell.Display(metadata, "organizationintrusion", cmd.Flags())
},
}



var malware = &cobra.Command{
Use:   "malware",
Short: "Returns all supported malware settings for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetMalwareSettings(networkId)
shell.Display(metadata, "malware", cmd.Flags())
},
}



var settings = &cobra.Command{
Use:   "settings",
Short: "Return the appliance settings for a network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetSettings(networkId)
shell.Display(metadata, "settings", cmd.Flags())
},
}



var singlelan = &cobra.Command{
Use:   "singlelan",
Short: "Return single LAN configuration.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetSingleLan(networkId)
shell.Display(metadata, "singlelan", cmd.Flags())
},
}



var staticroutes = &cobra.Command{
Use:   "staticroutes",
Short: "List the static routes for an MX or teleworker network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetStaticRoutes(networkId)
shell.Display(metadata, "staticroutes", cmd.Flags())
},
}


var staticroute = &cobra.Command{
	Use:   "staticroute",
	Short: "\tReturn a static route for an MX or teleworker network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		staticRouteId := args[0]
		metadata := configure.GetStaticRoute(networkId, staticRouteId)
		shell.Display(metadata, "staticroute", cmd.Flags())
	},
}


var customperformanceclasses = &cobra.Command{
Use:   "customperformanceclasses",
Short: "List all custom performance classes for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetPerformanceClasses(networkId)
shell.Display(metadata, "customperformanceclasses", cmd.Flags())
},
}

var customperformanceclass = &cobra.Command{
	Use:   "customperformanceclass",
	Short: "Return a custom performance class for an MX network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		customPerformanceClassId := args[0]
		metadata := configure.GetPerformanceClass(networkId,
			customPerformanceClassId)
		shell.Display(metadata, "customperformanceclass", cmd.Flags())
	},
}

var trafficshapingrules = &cobra.Command{
Use:   "trafficshapingrules",
Short: "Display the traffic shaping settings rules for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetTrafficShapingRules(networkId)
shell.Display(metadata, "trafficshapingrules", cmd.Flags())
},
}

var uplinkbandwidth = &cobra.Command{
Use:   "uplinkbandwidth",
Short: "Returns the uplink bandwidth settings for your MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetUplinkBandwidth(networkId)
shell.Display(metadata, "uplinkbandwidth", cmd.Flags())
},
}



var uplinkselection = &cobra.Command{
Use:   "uplinkselection",
Short: "Show uplink selection settings for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetUplinkSelection(networkId)
shell.Display(metadata, "uplinkselection", cmd.Flags())
},
}

var trafficshaping = &cobra.Command{
Use:   "trafficshaping",
Short: "Display the traffic shaping settings for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetTrafficShapingRules(networkId)
shell.Display(metadata, "trafficshaping", cmd.Flags())
},
}



var vlansettings = &cobra.Command{
Use:   "vlansettings",
Short: "Returns the enabled status of VLANs for the network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetVLANSettings(networkId)
shell.Display(metadata, "vlansettings", cmd.Flags())
},
}

var vlans = &cobra.Command{
Use:   "vlans",
Short: "List the VLANs for an MX network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetVLANs(networkId)
shell.Display(metadata, "vlans", cmd.Flags())
},
}

var vlan = &cobra.Command{
	Use:   "vlan",
	Short: "Return a VLAN.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		vlanId := args[0]
		metadata := configure.GetVLAN(networkId, vlanId)
		shell.Display(metadata, "vlan", cmd.Flags())
	},
}

var sitetositevpn = &cobra.Command{
Use:   "sitetositevpn",
Short: "Return the site-to-site VPN settings of a network. Only valid for MX networks.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetSiteToSiteVPN(networkId)
shell.Display(metadata, "sitetositevpn", cmd.Flags())
},
}

var thirdpartyvpnpeers = &cobra.Command{
Use:   "thirdpartyvpnpeers",
Short: "Return the third party VPN peers for an organization.",
Run: func(cmd *cobra.Command, args []string) {
orgId, _, _ := shell.ResolveFlags(cmd.Flags())
if orgId == "" {
orgId = args[0]
}

metadata := configure.GetThirdPartyVPNPeers(orgId)
shell.Display(metadata, "thirdpartyvpnpeers", cmd.Flags())
},
}



var vpnfirewallrules = &cobra.Command{
Use:   "vpnfirewallrules",
Short: "Return the firewall rules for an organization's site-to-site VPN.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetCellularFirewallRules(networkId)
shell.Display(metadata, "vpnfirewallrules", cmd.Flags())
},
}

var warmspare = &cobra.Command{
Use:   "warmspare",
Short: "Return MX warm spare settings.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

metadata := configure.GetWarmSpare(networkId)
shell.Display(metadata, "warmspare", cmd.Flags())
},
}