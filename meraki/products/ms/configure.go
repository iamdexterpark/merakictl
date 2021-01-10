package ms

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/switch/configure"
	shell "github.com/ddexterpark/merakictl/utilities"
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

var PutAccessControlLists  = &cobra.Command{
	Use:   "AccessControlLists",
	Short: "Return the access control lists for a MS network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.AccessControlLists
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutAccessControlLists(networkId,  input)
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

var DelAccessPolicy = &cobra.Command{
	Use:   "AccessPolicy",
	Short: "Return a specific access policy for a ms network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		accessPolicyNumber := args[0]
		metadata := configure.DelAccessPolicy(networkId, accessPolicyNumber)
		shell.Display(metadata, "AccessPolicy", cmd.Flags())
	},
}

var PutAccessPolicy = &cobra.Command{
	Use:   "AccessPolicy",
	Short: "Return a specific access policy for a ms network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		accessPolicyNumber := args[0]
		var format configure.AccessPolicy
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutAccessPolicy(networkId, accessPolicyNumber,  input)
		shell.Display(metadata, "AccessPolicy", cmd.Flags())
	},
}

var PostAccessPolicy = &cobra.Command{
	Use:   "AccessPolicy",
	Short: "Return a specific access policy for a ms network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.AccessPolicy
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostAccessPolicy(networkId,  input)
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

var PutSwitchPortProfile = &cobra.Command{
	Use:   "SwitchPortProfile",
	Short: "Return all the ports of a ms profile.",
	Run: func(cmd *cobra.Command, args []string) {
		orgId, _, _ := shell.ResolveFlags(cmd.Flags())
		if orgId == "" {
			orgId = args[3]
		}

		configTemplateId := args[0]
		profileId := args[1]
		portId := args[2]
		var format configure.SwitchProfile
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutSwitchPortProfile(orgId, configTemplateId, profileId, portId,  input)
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

var PostClone = &cobra.Command{
	Use:   "SwitchPortProfile",
	Short: "Return all the ports of a ms profile.",
	Run: func(cmd *cobra.Command, args []string) {
		orgId, _, _ := shell.ResolveFlags(cmd.Flags())
		if orgId == "" {
			orgId = args[0]
		}
		var format configure.Clone
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostClone(orgId,  input)
		shell.Display(metadata, "SwitchPortProfile", cmd.Flags())
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

var PutDHCPServerPolicy = &cobra.Command{
	Use:   "DHCPServerPolicy",
	Short: "Return the DHCP server policy.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.DHCPServerPolicy
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutDHCPServerPolicy(networkId,  input)
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

var PutDSCP = &cobra.Command{
	Use:   "DSCP",
	Short: "Return the DSCP to CoS mappings.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.DscpToCosMappings
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutDscpToCosMappings(networkId,  input)
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

var DelLinkAggregations = &cobra.Command{
	Use:   "LinkAggregations",
	Short: "Split a link aggregation group into separate ports",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		linkAggregationId := args[0]
		metadata := configure.DelLinkAggregation(networkId, linkAggregationId)
		shell.Display(metadata, "LinkAggregations", cmd.Flags())
	},
}

var PutLinkAggregations = &cobra.Command{
	Use:   "LinkAggregations",
	Short: "Split a link aggregation group into separate ports",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		linkAggregationId := args[0]
		var format configure.LinkAggregation
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutLinkAggregation(networkId, linkAggregationId,  input)
		shell.Display(metadata, "LinkAggregations", cmd.Flags())
	},
}

var PostLinkAggregations = &cobra.Command{
	Use:   "LinkAggregations",
	Short: "Split a link aggregation group into separate ports",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.LinkAggregation
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostLinkAggregation(networkId,  input)
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

var PutMTU = &cobra.Command{
	Use:   "MTU",
	Short: "Return the MTU configuration.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.MTU
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutMTU(networkId,  input)
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

var DelPortSchedules = &cobra.Command{
	Use:   "PortSchedules",
	Short: "List ms port schedules.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		portScheduleId := args[0]
		metadata := configure.DelPortSchedules(networkId, portScheduleId )
		shell.Display(metadata, "PortSchedules", cmd.Flags())
	},
}

var PutPortSchedules = &cobra.Command{
	Use:   "PortSchedules",
	Short: "List ms port schedules.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		portScheduleId := args[0]
		var format configure.PortSchedules
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutPortSchedules(networkId, portScheduleId,  input)
		shell.Display(metadata, "PortSchedules", cmd.Flags())
	},
}

var PostPortSchedules = &cobra.Command{
	Use:   "PortSchedules",
	Short: "List ms port schedules.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.PortSchedules
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostPortSchedules(networkId,  input)
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

var PutSwitchPort = &cobra.Command{
	Use:   "SwitchPort",
	Short: "Return a single ms port.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[1]
		}
		portId := args[0]
		var format configure.Port
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutSwitchPort(serial, portId,  input)
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

var PutQoSRuleOrder = &cobra.Command{
	Use:   "QoSRuleOrder",
	Short: "Return the quality of service rule IDs by order in which they will be processed by the ms.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.QoSRuleOrder
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutQoSRuleOrder(networkId,  input)
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

var DelQoSRule = &cobra.Command{
	Use:   "QoSRule",
	Short: "Return a quality of service rule.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		qosRuleId := args[0]
		metadata := configure.DelQoSRule(networkId, qosRuleId)
		shell.Display(metadata, "QoSRule", cmd.Flags())
	},
}

var PutQoSRule = &cobra.Command{
	Use:   "QoSRule",
	Short: "Return a quality of service rule.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		qosRuleId := args[0]
		var format configure.QoSRule
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutQoSRule(networkId, qosRuleId,  input)
		shell.Display(metadata, "QoSRule", cmd.Flags())
	},
}

var PostQoSRule = &cobra.Command{
	Use:   "QoSRule",
	Short: "Return a quality of service rule.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.QoSRule
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostQoSRule(networkId,  input)
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

var PutDHCP = &cobra.Command{
	Use:   "DHCP",
	Short: "Return a layer 3 interface DHCP configuration for a ms.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.DHCP
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutDHCPServerPolicy(networkId,  input)
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

var DelL3Interface = &cobra.Command{
	Use:   "L3Interface",
	Short: "Return a layer 3 interface for a ms.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[1]
		}

		interfaceId := args[0]
		metadata := configure.DelInterface(serial, interfaceId)
		shell.Display(metadata, "L3Interface", cmd.Flags())
	},
}

var PutL3Interface = &cobra.Command{
	Use:   "L3Interface",
	Short: "Return a layer 3 interface for a ms.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[1]
		}

		interfaceId := args[0]
		var format configure.Interface
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutInterface(serial, interfaceId,  input)
		shell.Display(metadata, "L3Interface", cmd.Flags())
	},
}

var PostL3Interface = &cobra.Command{
	Use:   "L3Interface",
	Short: "Return a layer 3 interface for a ms.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		var format configure.Interface
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostInterface(serial,  input)
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

//FIX IN DASHBOARD-API-GOLANG
var DelRendezvousPoint = &cobra.Command{
	Use:   "RendezvousPoint",
	Short: "Delete a multicast rendezvous point.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		//rendezvousPointId := args[0]
		//metadata := configure.DelRendezvousPoint(networkId, rendezvousPointId)
		//shell.Display(metadata, "RendezvousPoint", cmd.Flags())
	},
}

var PutRendezvousPoint = &cobra.Command{
	Use:   "RendezvousPoint",
	Short: "List multicast rendezvous points.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		rendezvousPointId := args[0]
		var format configure.RendezvousPoint
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutRendezvousPoint(networkId, rendezvousPointId,  input)
		shell.Display(metadata, "RendezvousPoint", cmd.Flags())
	},
}

var PostRendezvousPoint = &cobra.Command{
	Use:   "RendezvousPoint",
	Short: "List multicast rendezvous points.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.RendezvousPoint
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostRendezvousPoint(networkId,  input)
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

var PutMulticast = &cobra.Command{
	Use:   "Multicast",
	Short: "Return multicast settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.Multicast
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutMulticast(networkId,  input)
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

var PutOSPF = &cobra.Command{
	Use:   "OSPF",
	Short: "Return layer 3 OSPF routing configuration.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.OSPF
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutOSPF(networkId,  input)
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

var DelStaticRoute = &cobra.Command{
	Use:   "StaticRoute",
	Short: "Return a layer 3 static route for a ms.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[1]
		}
		staticRouteId := args[0]
		metadata := configure.DelStaticRoute(serial, staticRouteId)
		shell.Display(metadata, "StaticRoute", cmd.Flags())
	},
}

var PutStaticRoute = &cobra.Command{
	Use:   "StaticRoute",
	Short: "Return a layer 3 static route for a ms.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[1]
		}
		staticRouteId := args[0]
		var format configure.StaticRoute
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutStaticRoute(serial, staticRouteId,  input)
		shell.Display(metadata, "StaticRoute", cmd.Flags())
	},
}

var PostStaticRoute = &cobra.Command{
	Use:   "StaticRoute",
	Short: "Return a layer 3 static route for a ms.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		var format configure.StaticRoute
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostStaticRoute(serial,  input)
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

var PutSettings = &cobra.Command{
	Use:   "Settings",
	Short: "Returns the ms network settings.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.Settings
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutSettings(networkId,  input)
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

var PutStackDHCP = &cobra.Command{
	Use:   "StackDHCP",
	Short: "Return a layer 3 interface DHCP configuration for a ms stack.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[2]
		}

		switchStackId := args[0]
		interfaceId := args[1]
		var format configure.StackDHCP
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutStackDHCP(networkId, switchStackId, interfaceId,  input)
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
			networkId = args[2]
		}
		switchStackId := args[0]
		interfaceId := args[1]
		metadata := configure.GetStackInterface(networkId, switchStackId, interfaceId)
		shell.Display(metadata, "StackL3Interface", cmd.Flags())
	},
}

var DelStackL3Interface = &cobra.Command{
	Use:   "StackL3Interface",
	Short: "List layer 3 interfaces for a ms stack.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[2]
		}
		switchStackId := args[0]
		interfaceId := args[1]
		metadata := configure.DelStackInterface(networkId, switchStackId, interfaceId)
		shell.Display(metadata, "StackL3Interface", cmd.Flags())
	},
}

var PutStackL3Interface = &cobra.Command{
	Use:   "StackL3Interface",
	Short: "List layer 3 interfaces for a ms stack.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[2]
		}
		switchStackId := args[0]
		interfaceId := args[1]
		var format configure.StackInterfaces
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutStackInterface(networkId, switchStackId, interfaceId,  input)
		shell.Display(metadata, "StackL3Interface", cmd.Flags())
	},
}

var PostStackL3Interface = &cobra.Command{
	Use:   "StackL3Interface",
	Short: "List layer 3 interfaces for a ms stack.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		switchStackId := args[0]
		var format configure.StackInterface
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostStackInterface(networkId, switchStackId,  input)
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

var DelStackStaticRoute = &cobra.Command{
	Use:   "StackStaticRoute",
	Short: "Return a layer 3 static route for a ms stack.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[2]
		}

		switchStackId := args[0]
		staticRouteId := args[1]
		metadata := configure.DelStackStaticRoute(networkId, switchStackId, staticRouteId)
		shell.Display(metadata, "StackStaticRoute", cmd.Flags())
	},
}

var PutStackStaticRoute = &cobra.Command{
	Use:   "StackStaticRoute",
	Short: "Return a layer 3 static route for a ms stack.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[2]
		}

		switchStackId := args[0]
		staticRouteId := args[1]
		var format configure.StackStaticRoute
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutStackStaticRoute(networkId, switchStackId, staticRouteId,  input)
		shell.Display(metadata, "StackStaticRoute", cmd.Flags())
	},
}

var PostStackStaticRoute = &cobra.Command{
	Use:   "StackStaticRoute",
	Short: "Return a layer 3 static route for a ms stack.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		switchStackId := args[0]
		var format configure.StackStaticRoute
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostStackStaticRoute(networkId, switchStackId,  input)
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


var PostAddToStack = &cobra.Command{
	Use:   "AddToStack",
	Short: "Show a ms stack.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		switchStackId := args[0]
		var format configure.Stacks
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostAddToStack(networkId, switchStackId,  input)
		shell.Display(metadata, "AddToStack", cmd.Flags())
	},
}

var PostSwitchStack = &cobra.Command{
	Use:   "SwitchStack",
	Short: "Show a ms stack.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		switchStackId := args[0]
		metadata := configure.PostStack(networkId, switchStackId)
		shell.Display(metadata, "SwitchStack", cmd.Flags())
	},
}

var DelSwitchStack = &cobra.Command{
	Use:   "SwitchStack",
	Short: "Show a ms stack.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		switchStackId := args[0]
		metadata := configure.DelStack(networkId, switchStackId)
		shell.Display(metadata, "SwitchStack", cmd.Flags())
	},
}

var PostRemoveFromStack = &cobra.Command{
	Use:   "RemoveFromStack",
	Short: "Remove a switch from a stack.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		switchStackId := args[0]
		var format configure.Stack
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostRemoveFromStack(networkId, switchStackId,  input)
		shell.Display(metadata, "RemoveFromStack", cmd.Flags())
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

var PutStormControl = &cobra.Command{
	Use:   "StormControl",
	Short: "Return the storm control configuration for a ms network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.StormControl
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutStormControl(networkId,  input)
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

var PutSTP = &cobra.Command{
	Use:   "STP",
	Short: "Return warm spare configuration for a ms.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.STP
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutSTP(networkId,  input)
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

var PutWarmSpare = &cobra.Command{
	Use:   "WarmSpare",
	Short: "Return warm spare configuration for a ms.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		var format configure.WarmSpare
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutWarmSpare(serial,  input)
		shell.Display(metadata, "WarmSpare", cmd.Flags())
	},
}