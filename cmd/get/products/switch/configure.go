package _switch

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/switch/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)


var accesscontrollists  = &cobra.Command{
Use:   "accesscontrollists",
Short: "Return the access control lists for a MS network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetAccessControlLists(networkId)
shell.Display(metadata, "accesscontrollists", cmd.Flags())
},
}

var accesspolicies = &cobra.Command{
Use:   "accesspolicies",
Short: "List the access policies for a switch network. Only returns access policies with 'my RADIUS server' as authentication method.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetAccessPolicies(networkId)
shell.Display(metadata, "accesspolicies", cmd.Flags())
},
}

var accesspolicy = &cobra.Command{
Use:   "accesspolicy",
Short: "Return a specific access policy for a switch network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	accessPolicyNumber := args[0]
metadata := configure.GetAccessPolicy(networkId, accessPolicyNumber)
shell.Display(metadata, "accesspolicy", cmd.Flags())
},
}

var switchportsprofiles = &cobra.Command{
Use:   "switchportsprofile",
Short: "Return a switch profile port.",
Run: func(cmd *cobra.Command, args []string) {
	orgId, _, _ := shell.ResolveFlags(cmd.Flags())
	if orgId == "" {
	orgId = args[3]
	}

	configTemplateId := args[0]
	profileId := args[1]

metadata := configure.GetSwitchPortProfiles(orgId,
	configTemplateId, profileId)
shell.Display(metadata, "switchportsprofile", cmd.Flags())
},
}

var switchportprofile = &cobra.Command{
Use:   "switchportprofile",
Short: "Return all the ports of a switch profile.",
Run: func(cmd *cobra.Command, args []string) {
orgId, _, _ := shell.ResolveFlags(cmd.Flags())
if orgId == "" {
	orgId = args[0]
}

	configTemplateId := args[0]
	profileId := args[1]
	portId := args[2]

metadata := configure.GetSwitchPortProfile(orgId,configTemplateId, profileId, portId)
shell.Display(metadata, "switchportprofile", cmd.Flags())
},
}

var switchprofiles = &cobra.Command{
Use:   "switchprofiles",
Short: "List the switch profiles for your switch template configuration.",
Run: func(cmd *cobra.Command, args []string) {
	orgId, _, _ := shell.ResolveFlags(cmd.Flags())
	if orgId == "" {
		orgId = args[1]
	}

	configTemplateId := args[0]

metadata := configure.GetSwitchTemplateConfigProfiles(orgId, configTemplateId)
shell.Display(metadata, "switchprofiles", cmd.Flags())
},
}

var dhcpserverpolicy = &cobra.Command{
Use:   "dhcpserverpolicy",
Short: "Return the DHCP server policy.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetDHCPServerPolicy(networkId)
shell.Display(metadata, "dhcpserverpolicy", cmd.Flags())
},
}

var dscp = &cobra.Command{
Use:   "dscp",
Short: "Return the DSCP to CoS mappings.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetDSCPCOSMapping(networkId)
shell.Display(metadata, "dscp", cmd.Flags())
},
}

var linkaggregations = &cobra.Command{
Use:   "linkaggregations",
Short: "List link aggregation groups.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetLinkAggregationGroups(networkId)
shell.Display(metadata, "linkaggregations", cmd.Flags())
},
}

var mtu = &cobra.Command{
Use:   "mtu",
Short: "Return the MTU configuration.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetMTU(networkId)
shell.Display(metadata, "mtu", cmd.Flags())
},
}

var portschedules = &cobra.Command{
Use:   "portschedules",
Short: "List switch port schedules.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetSwitchPortSchedules(networkId)
shell.Display(metadata, "portschedules", cmd.Flags())
},
}

var switchports = &cobra.Command{
Use:   "ports",
Short: "List the switch ports for a switch.",
Run: func(cmd *cobra.Command, args []string) {
_, _, serial := shell.ResolveFlags(cmd.Flags())
if serial == "" {
	serial = args[0]
}
metadata := configure.GetSwitchPorts(serial)
shell.Display(metadata, "switchports", cmd.Flags())
},
}

var switchport = &cobra.Command{
Use:   "switchport",
Short: "Return a single switch port.",
Run: func(cmd *cobra.Command, args []string) {
	_, _, serial := shell.ResolveFlags(cmd.Flags())
	if serial == "" {
		serial = args[0]
	}
	portId := args[0]
metadata := configure.GetSwitchPort(serial, portId)
shell.Display(metadata, "switchport", cmd.Flags())
},
}

var qosruleids = &cobra.Command{
Use:   "qosruleids",
Short: "Return the quality of service rule IDs by order in which they will be processed by the switch.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetQoSRuleIds(networkId)
shell.Display(metadata, "qosruleids", cmd.Flags())
},
}

var qosrules = &cobra.Command{
Use:   "qosrules",
Short: "ist quality of service rules.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetQoSRules(networkId)
shell.Display(metadata, "qosrules", cmd.Flags())
},
}

var qosrule = &cobra.Command{
Use:   "qosrule",
Short: "Return a quality of service rule.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	qosRuleId := args[0]
metadata := configure.GetQoSRule(networkId, qosRuleId)
shell.Display(metadata, "qosrule", cmd.Flags())
},
}

var dhcp = &cobra.Command{
Use:   "dhcp",
Short: "Return a layer 3 interface DHCP configuration for a switch.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetDHCPServerPolicy(networkId)
shell.Display(metadata, "dhcp", cmd.Flags())
},
}

var l3interfaces = &cobra.Command{
Use:   "l3interfaces",
Short: "List layer 3 interfaces for a switch.",
Run: func(cmd *cobra.Command, args []string) {
	_, _, serial := shell.ResolveFlags(cmd.Flags())
	if serial == "" {
		serial = args[0]
	}
metadata := configure.GetLayer3Interfaces(serial)
shell.Display(metadata, "l3interfaces", cmd.Flags())
},
}

var l3interface = &cobra.Command{
	Use:   "l3interface",
	Short: "Return a layer 3 interface for a switch.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[1]
		}

		interfaceId := args[0]
		metadata := configure.GetLayer3Interface(serial, interfaceId)
		shell.Display(metadata, "l3interface", cmd.Flags())
	},
}

var rendezvouspoints = &cobra.Command{
Use:   "rendezvouspoints",
Short: "List multicast rendezvous points.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetMulticastRendezvousPoints(networkId)
shell.Display(metadata, "rendezvouspoints", cmd.Flags())
},
}

var rendezvouspoint = &cobra.Command{
	Use:   "rendezvouspoint",
	Short: "List multicast rendezvous points.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		rendezvousPointId := args[0]
		metadata := configure.GetMulticastRendezvousPoint(networkId, rendezvousPointId)
		shell.Display(metadata, "rendezvouspoint", cmd.Flags())
	},
}

var multicast = &cobra.Command{
Use:   "multicast",
Short: "Return multicast settings for a network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetMulticastSettings(networkId)
shell.Display(metadata, "multicast", cmd.Flags())
},
}

var ospf = &cobra.Command{
Use:   "ospf",
Short: "Return layer 3 OSPF routing configuration.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}
metadata := configure.GetOSPFRouting(networkId)
shell.Display(metadata, "ospf", cmd.Flags())
},
}

var staticroutes = &cobra.Command{
Use:   "staticroutes",
Short: "List layer 3 static routes for a switch.",
Run: func(cmd *cobra.Command, args []string) {
	_, _, serial := shell.ResolveFlags(cmd.Flags())
	if serial == "" {
		serial = args[0]
	}
metadata := configure.GetStaticRoutes(serial)
shell.Display(metadata, "staticroutes", cmd.Flags())
},
}

var staticroute = &cobra.Command{
	Use:   "staticroute",
	Short: "Return a layer 3 static route for a switch.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[1]
		}
		staticRouteId := args[0]
		metadata := configure.GetStaticRoute(serial, staticRouteId)
		shell.Display(metadata, "staticroute", cmd.Flags())
	},
}

var settings = &cobra.Command{
Use:   "settings",
Short: "Returns the switch network settings.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}
metadata := configure.GetSwitchNetworkSettings(networkId)
shell.Display(metadata, "settings", cmd.Flags())
},
}

var stackdhcp = &cobra.Command{
Use:   "stackdhcp",
Short: "Return a layer 3 interface DHCP configuration for a switch stack.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[2]
}

	switchStackId := args[0]
	interfaceId := args[1]
metadata := configure.GetInterfaceDHCPConfiguration(networkId, switchStackId, interfaceId)
shell.Display(metadata, "stackdhcp", cmd.Flags())
},
}

var stackl3interfaces = &cobra.Command{
Use:   "stackl3interfaces",
Short: "List layer 3 interfaces for a switch stack.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[1]
}
	switchStackId := args[0]
metadata := configure.GetStackLayer3Interfaces(networkId, switchStackId)
shell.Display(metadata, "stackl3interfaces", cmd.Flags())
},
}

var stackl3interface = &cobra.Command{
	Use:   "stackl3interface",
	Short: "List layer 3 interfaces for a switch stack.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		switchStackId := args[0]
		interfaceId := args[1]
		metadata := configure.GetStackLayer3Interface(networkId, switchStackId, interfaceId)
		shell.Display(metadata, "stackl3interface", cmd.Flags())
	},
}


var stackstaticroutes = &cobra.Command{
Use:   "stackstaticroutes",
Short: "List layer 3 static routes for a switch stack.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[1]
}

	switchStackId := args[0]

metadata := configure.GetStackStaticRoutes(networkId, switchStackId)
shell.Display(metadata, "stackstaticroutes", cmd.Flags())
},
}

var stackstaticroute = &cobra.Command{
	Use:   "stackstaticroute",
	Short: "Return a layer 3 static route for a switch stack.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[2]
		}

		switchStackId := args[0]
		staticRouteId := args[1]
		metadata := configure.GetStackStaticRoute(networkId, switchStackId, staticRouteId)
		shell.Display(metadata, "stackstaticroute", cmd.Flags())
	},
}


var switchstacks = &cobra.Command{
Use:   "switchstacks",
Short: "List the switch stacks in a network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}
metadata := configure.GetSwitchStacks(networkId)
shell.Display(metadata, "switchstacks", cmd.Flags())
},
}

var switchstack = &cobra.Command{
	Use:   "switchstack",
	Short: "Show a switch stack.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		switchStackId := args[0]
		metadata := configure.GetSwitchStack(networkId, switchStackId)
		shell.Display(metadata, "switchstack", cmd.Flags())
	},
}

var stormcontrol = &cobra.Command{
Use:   "stormcontrol",
Short: "Return the storm control configuration for a switch network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}
metadata := configure.GetStormControl(networkId)
shell.Display(metadata, "stormcontrol", cmd.Flags())
},
}

var stp = &cobra.Command{
Use:   "stp",
Short: "Return warm spare configuration for a switch.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}
metadata := configure.GetSTPSettings(networkId)
shell.Display(metadata, "stp", cmd.Flags())
},
}

var warmspare = &cobra.Command{
Use:   "warmspare",
Short: "Return warm spare configuration for a switch.",
Run: func(cmd *cobra.Command, args []string) {
_, _, serial := shell.ResolveFlags(cmd.Flags())
if serial == "" {
	serial = args[0]
}
metadata := configure.GetWarmSpare(serial)
shell.Display(metadata, "warmspare", cmd.Flags())
},
}