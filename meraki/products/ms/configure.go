package ms

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/switch/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)


var GetAccessControlLists  = &cobra.Command{
Use:   "AccessControlLists",
Short: "Return the access control lists for a MS network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetAccessControlLists(networkId)
shell.Display(metadata, "AccessControlLists", cmd.Flags())
},
}

var GetAccessPolicies = &cobra.Command{
Use:   "AccessPolicies",
Short: "List the access policies for a ms network. Only returns access policies with 'my RADIUS server' as authentication method.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetAccessPolicies(networkId)
shell.Display(metadata, "AccessPolicies", cmd.Flags())
},
}

var GetAccessPolicy = &cobra.Command{
Use:   "AccessPolicy",
Short: "Return a specific access policy for a ms network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	accessPolicyNumber := args[0]
metadata := configure.GetAccessPolicy(networkId, accessPolicyNumber)
shell.Display(metadata, "AccessPolicy", cmd.Flags())
},
}

var GetSwitchPortsProfiles = &cobra.Command{
Use:   "SwitchPortsProfiles",
Short: "Return a ms profile port.",
Run: func(cmd *cobra.Command, args []string) {
	orgId, _, _ := shell.ResolveFlags(cmd.Flags())
	if orgId == "" {
	orgId = args[3]
	}

	configTemplateId := args[0]
	profileId := args[1]

metadata := configure.GetSwitchPortProfiles(orgId,
	configTemplateId, profileId)
shell.Display(metadata, "SwitchPortsProfiles", cmd.Flags())
},
}

var GetSwitchPortProfile = &cobra.Command{
Use:   "SwitchPortProfile",
Short: "Return all the ports of a ms profile.",
Run: func(cmd *cobra.Command, args []string) {
orgId, _, _ := shell.ResolveFlags(cmd.Flags())
if orgId == "" {
	orgId = args[0]
}

	configTemplateId := args[0]
	profileId := args[1]
	portId := args[2]

metadata := configure.GetSwitchPortProfile(orgId, configTemplateId, profileId, portId)
shell.Display(metadata, "SwitchPortProfile", cmd.Flags())
},
}

var GetSwitchProfiles = &cobra.Command{
Use:   "SwitchProfiles",
Short: "List the ms profiles for your ms template configuration.",
Run: func(cmd *cobra.Command, args []string) {
	orgId, _, _ := shell.ResolveFlags(cmd.Flags())
	if orgId == "" {
		orgId = args[1]
	}

	configTemplateId := args[0]
	profileId := args[1]
metadata := configure.GetSwitchPortProfiles(orgId, configTemplateId, profileId)
shell.Display(metadata, "SwitchProfiles", cmd.Flags())
},
}

var GetDHCPServerPolicy = &cobra.Command{
Use:   "DHCPServerPolicy",
Short: "Return the DHCP server policy.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetDHCPServerPolicy(networkId)
shell.Display(metadata, "DHCPServerPolicy", cmd.Flags())
},
}

var GetDSCP = &cobra.Command{
Use:   "DSCP",
Short: "Return the DSCP to CoS mappings.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetDscpToCosMappings(networkId)
shell.Display(metadata, "DSCP", cmd.Flags())
},
}

var GetLinkAggregations = &cobra.Command{
Use:   "LinkAggregations",
Short: "List link aggregation groups.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetLinkAggregations(networkId)
shell.Display(metadata, "LinkAggregations", cmd.Flags())
},
}

var GetMTU = &cobra.Command{
Use:   "MTU",
Short: "Return the MTU configuration.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetMTU(networkId)
shell.Display(metadata, "MTU", cmd.Flags())
},
}

var GetPortSchedules = &cobra.Command{
Use:   "PortSchedules",
Short: "List ms port schedules.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetPortSchedules(networkId)
shell.Display(metadata, "PortSchedules", cmd.Flags())
},
}

var GetSwitchPorts = &cobra.Command{
Use:   "SwitchPorts",
Short: "List the ms ports for a ms.",
Run: func(cmd *cobra.Command, args []string) {
_, _, serial := shell.ResolveFlags(cmd.Flags())
if serial == "" {
	serial = args[0]
}
metadata := configure.GetPorts(serial)
shell.Display(metadata, "SwitchPorts", cmd.Flags())
},
}

var GetSwitchPort = &cobra.Command{
Use:   "SwitchPort",
Short: "Return a single ms port.",
Run: func(cmd *cobra.Command, args []string) {
	_, _, serial := shell.ResolveFlags(cmd.Flags())
	if serial == "" {
		serial = args[0]
	}
	portId := args[0]
metadata := configure.GetSwitchPort(serial, portId)
shell.Display(metadata, "SwitchPort", cmd.Flags())
},
}

var GetQoSRuleOrder = &cobra.Command{
Use:   "QoSRuleOrder",
Short: "Return the quality of service rule IDs by order in which they will be processed by the ms.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetQoSRuleOrder(networkId)
shell.Display(metadata, "QoSRuleOrder", cmd.Flags())
},
}

var GetQoSRules = &cobra.Command{
Use:   "qosrules",
Short: "first quality of service rules.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetQoSRules(networkId)
shell.Display(metadata, "QoSRules", cmd.Flags())
},
}

var GetQoSRule = &cobra.Command{
Use:   "QoSRule",
Short: "Return a quality of service rule.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	qosRuleId := args[0]
metadata := configure.GetQoSRule(networkId, qosRuleId)
shell.Display(metadata, "QoSRule", cmd.Flags())
},
}

var GetDHCP = &cobra.Command{
Use:   "DHCP",
Short: "Return a layer 3 interface DHCP configuration for a ms.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetDHCPServerPolicy(networkId)
shell.Display(metadata, "DHCP", cmd.Flags())
},
}

var GetL3Interfaces = &cobra.Command{
Use:   "L3Interfaces",
Short: "List layer 3 interfaces for a ms.",
Run: func(cmd *cobra.Command, args []string) {
	_, _, serial := shell.ResolveFlags(cmd.Flags())
	if serial == "" {
		serial = args[0]
	}
metadata := configure.GetInterfaces(serial)
shell.Display(metadata, "L3Interfaces", cmd.Flags())
},
}

var GetL3Interface = &cobra.Command{
	Use:   "L3Interface",
	Short: "Return a layer 3 interface for a ms.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[1]
		}

		interfaceId := args[0]
		metadata := configure.GetInterface(serial, interfaceId)
		shell.Display(metadata, "L3Interface", cmd.Flags())
	},
}

var GetRendezvousPoints = &cobra.Command{
Use:   "RendezvousPoints",
Short: "List multicast rendezvous points.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetRendezvousPoints(networkId)
shell.Display(metadata, "RendezvousPoints", cmd.Flags())
},
}

var GetRendezvousPoint = &cobra.Command{
	Use:   "RendezvousPoint",
	Short: "List multicast rendezvous points.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		rendezvousPointId := args[0]
		metadata := configure.GetRendezvousPoint(networkId, rendezvousPointId)
		shell.Display(metadata, "RendezvousPoint", cmd.Flags())
	},
}

var GetMulticast = &cobra.Command{
Use:   "Multicast",
Short: "Return multicast settings for a network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetMulticast(networkId)
shell.Display(metadata, "Multicast", cmd.Flags())
},
}

var GetOSPF = &cobra.Command{
Use:   "OSPF",
Short: "Return layer 3 OSPF routing configuration.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}
metadata := configure.GetOSPF(networkId)
shell.Display(metadata, "OSPF", cmd.Flags())
},
}

var GetStaticRoutes = &cobra.Command{
Use:   "StaticRoutes",
Short: "List layer 3 static routes for a ms.",
Run: func(cmd *cobra.Command, args []string) {
	_, _, serial := shell.ResolveFlags(cmd.Flags())
	if serial == "" {
		serial = args[0]
	}
metadata := configure.GetStaticRoutes(serial)
shell.Display(metadata, "StaticRoutes", cmd.Flags())
},
}

var GetStaticRoute = &cobra.Command{
	Use:   "StaticRoute",
	Short: "Return a layer 3 static route for a ms.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[1]
		}
		staticRouteId := args[0]
		metadata := configure.GetStaticRoute(serial, staticRouteId)
		shell.Display(metadata, "StaticRoute", cmd.Flags())
	},
}

var GetSettings = &cobra.Command{
Use:   "Settings",
Short: "Returns the ms network settings.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}
metadata := configure.GetSettings(networkId)
shell.Display(metadata, "Settings", cmd.Flags())
},
}

var GetStackDHCP = &cobra.Command{
Use:   "StackDHCP",
Short: "Return a layer 3 interface DHCP configuration for a ms stack.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[2]
}

	switchStackId := args[0]
	interfaceId := args[1]
metadata := configure.GetStackDHCP(networkId, switchStackId, interfaceId)
shell.Display(metadata, "StackDHCP", cmd.Flags())
},
}

var GetStackL3Interfaces = &cobra.Command{
Use:   "StackL3Interfaces",
Short: "List layer 3 interfaces for a ms stack.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[1]
}
	switchStackId := args[0]
metadata := configure.GetStackInterfaces(networkId, switchStackId)
shell.Display(metadata, "StackL3Interfaces", cmd.Flags())
},
}

var GetStackL3Interface = &cobra.Command{
	Use:   "StackL3Interface",
	Short: "List layer 3 interfaces for a ms stack.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		switchStackId := args[0]
		interfaceId := args[1]
		metadata := configure.GetStackInterface(networkId, switchStackId, interfaceId)
		shell.Display(metadata, "StackL3Interface", cmd.Flags())
	},
}


var GetStackStaticRoutes = &cobra.Command{
Use:   "StackStaticRoutes",
Short: "List layer 3 static routes for a ms stack.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[1]
}

	switchStackId := args[0]

metadata := configure.GetStackStaticRoutes(networkId, switchStackId)
shell.Display(metadata, "StackStaticRoutes", cmd.Flags())
},
}

var GetStackStaticRoute = &cobra.Command{
	Use:   "StackStaticRoute",
	Short: "Return a layer 3 static route for a ms stack.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[2]
		}

		switchStackId := args[0]
		staticRouteId := args[1]
		metadata := configure.GetStackStaticRoute(networkId, switchStackId, staticRouteId)
		shell.Display(metadata, "StackStaticRoute", cmd.Flags())
	},
}


var GetSwitchStacks = &cobra.Command{
Use:   "SwitchStacks",
Short: "List the ms stacks in a network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}
metadata := configure.GetStacks(networkId)
shell.Display(metadata, "SwitchStacks", cmd.Flags())
},
}

var GetSwitchStack = &cobra.Command{
	Use:   "SwitchStack",
	Short: "Show a ms stack.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		switchStackId := args[0]
		metadata := configure.GetStack(networkId, switchStackId)
		shell.Display(metadata, "SwitchStack", cmd.Flags())
	},
}

var GetStormControl = &cobra.Command{
Use:   "StormControl",
Short: "Return the storm control configuration for a ms network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}
metadata := configure.GetStormControl(networkId)
shell.Display(metadata, "StormControl", cmd.Flags())
},
}

var GetSTP = &cobra.Command{
Use:   "STP",
Short: "Return warm spare configuration for a ms.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}
metadata := configure.GetSTP(networkId)
shell.Display(metadata, "STP", cmd.Flags())
},
}

var GetWarmSpare = &cobra.Command{
Use:   "WarmSpare",
Short: "Return warm spare configuration for a ms.",
Run: func(cmd *cobra.Command, args []string) {
_, _, serial := shell.ResolveFlags(cmd.Flags())
if serial == "" {
	serial = args[0]
}
metadata := configure.GetWarmSpare(serial)
shell.Display(metadata, "WarmSpare", cmd.Flags())
},
}