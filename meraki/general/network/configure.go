package network

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/networks/configure"
	organizations "github.com/ddexterpark/dashboard-api-golang/api/general/organizations/configure"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)

var GetAlertSettings = &cobra.Command{
	Use:   "alertSettings",
	Short: "Return The Alert Configuration For This Network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetAlertSettings(networkId)
		shell.Display(metadata, "AlertSettings", cmd.Flags())
	},
}

var PutAlertSettings = &cobra.Command{
	Use:   "alertSettings",
	Short: "Update the alert configuration for this network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.AlertSettings
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutAlertSettings(networkId,  input)
		shell.Display(metadata, "AlertSettings", cmd.Flags())
	},
}

var GetClientPolicy = &cobra.Command{
	Use:   "clientPolicy",
	Short: "Return the policy assigned to a client on the network",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		clientId := args[0]
		metadata := configure.GetClientPolicy(networkId, clientId)
		shell.Display(metadata, "ClientPolicy", cmd.Flags())
	},
}

var PutClientPolicy = &cobra.Command{
	Use:   "clientPolicy",
	Short: "Update the policy assigned to a client on the network",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		clientId := args[0]
		var format configure.ClientPolicy
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutClientPolicy(networkId, clientId,  input)
		shell.Display(metadata, "ClientPolicy", cmd.Flags())
	},
}

var GetSplashAuthorizationStatus = &cobra.Command{
	Use:   "splashAuthorizationStatus",
	Short: "Return the splash authorization for a client",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		clientId := args[0]
		metadata := configure.GetSplashAuthorizationStatus(networkId, clientId)
		shell.Display(metadata, "SplashAuthorizationStatus", cmd.Flags())
	},
}

var PutSplashAuthorizationStatus = &cobra.Command{
	Use:   "splashAuthorizationStatus",
	Short: "Update a client's splash authorization.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		clientId := args[0]
		var format configure.SplashAuthorizationStatus
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutSplashAuthorizationStatus(networkId, clientId,  input)
		shell.Display(metadata, "SplashAuthorizationStatus", cmd.Flags())
	},
}

var PostProvisionClients = &cobra.Command{
	Use:   "provisionClients",
	Short: "Update a client's splash authorization.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.ProvisionClients
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostProvisionClients(networkId,  input)
		shell.Display(metadata, "ProvisionClients", cmd.Flags())
	},
}



var GetDevices = &cobra.Command{
	Use:   "devices",
	Short: "List the devices in a network",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetDevices(networkId)
		shell.Display(metadata, "Devices", cmd.Flags())
	},
}

var PostClaimDevices = &cobra.Command{
	Use:   "claimDevices",
	Short: "Claim devices into a network",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		var format interface{}
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostClaimDevices(networkId,  input)
		shell.Display(metadata, "ClaimSerials", cmd.Flags())
	},
}

var PostRemoveDevices = &cobra.Command{
	Use:   "removeDevices",
	Short: "Remove a single device",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		var format interface{}
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostRemoveDevices(networkId,  input)
		shell.Display(metadata, "RemoveDevices", cmd.Flags())
	},
}

var GetFirmwareUpgrades = &cobra.Command{
	Use:   "firmwareUpgrades",
	Short: "Get current maintenance window for a network",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetFirmwareUpgrades(networkId)
		shell.Display(metadata, "FirmwareUpgrades", cmd.Flags())
	},
}


var PutFirmwareUpgrades = &cobra.Command{
	Use:   "firmwareUpgrades",
	Short: "Update current maintenance window for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		var format configure.FirmwareUpgrades
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutFirmwareUpgrades(networkId,  input)
		shell.Display(metadata, "FirmwareUpgrades", cmd.Flags())
	},
}


var GetFloorPlans = &cobra.Command{
	Use:   "floorPlans",
	Short: "List the floor plans that belong to your network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetFloorPlans(networkId)
		shell.Display(metadata, "FloorPlans", cmd.Flags())
	},
}

var PostFloorPlan = &cobra.Command{
	Use:   "floorPlan",
	Short: "Upload a floor plan",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		var format configure.FloorPlan
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostFloorPlan(networkId,  input)
		shell.Display(metadata, "FloorPlan", cmd.Flags())
	},
}

var GetFloorPlan = &cobra.Command{
	Use:   "floorplan",
	Short: "Find a floor plan by ID",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		floorplanId := args[0]
		metadata := configure.GetFloorPlan(networkId, floorplanId)
		shell.Display(metadata, "FloorPlan", cmd.Flags())
	},
}

var PutFloorPlan = &cobra.Command{
	Use:   "floorPlan",
	Short: "Update a floor plan's geolocation and other meta data",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		floorplanId := args[0]
		var format configure.FloorPlan
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutFloorPlan(networkId, floorplanId,  input)
		shell.Display(metadata, "FloorPlan", cmd.Flags())
	},
}

var DelFloorPlan = &cobra.Command{
	Use:   "floorPlan",
	Short: "Destroy a floor plan",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		floorplanId := args[0]
		metadata := configure.DelFloorPlan(networkId, floorplanId)
		shell.Display(metadata, "FloorPlan", cmd.Flags())
	},
}

var GetGroupPolicies = &cobra.Command{
	Use:   "groupPolicies",
	Short: "List the group policies in a network",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetGroupPolicies(networkId)
		shell.Display(metadata, "GroupPolicies", cmd.Flags())
	},
}

var GetGroupPolicy = &cobra.Command{
	Use:   "groupPolicy ",
	Short: "Display a group policy.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		groupPolicyId := args[0]
		metadata := configure.GetGroupPolicy(networkId, groupPolicyId)
		shell.Display(metadata, "GroupPolicy", cmd.Flags())
	},
}

var PostGroupPolicy = &cobra.Command{
	Use:   "groupPolicy ",
	Short: "Create a group policy",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.GroupPolicy
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostGroupPolicy(networkId,  input)
		shell.Display(metadata, "GroupPolicy", cmd.Flags())
	},
}


var PutGroupPolicy = &cobra.Command{
	Use:   "groupPolicy ",
	Short: "Update a group policy",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		groupPolicyId := args[0]
		var format configure.GroupPolicy
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutGroupPolicy(networkId, groupPolicyId,  input)
		shell.Display(metadata, "GroupPolicy", cmd.Flags())
	},
}

var DelGroupPolicy = &cobra.Command{
	Use:   "groupPolicy ",
	Short: "Delete a group policy",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		groupPolicyId := args[0]
		metadata := configure.DelGroupPolicy(networkId, groupPolicyId)
		shell.Display(metadata, "GroupPolicy", cmd.Flags())
	},
}

var GetMerakiAuthUsers = &cobra.Command{
	Use:   "merakiAuthUsers",
	Short: "List the users configured under Meraki Authentication for a network",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetMerakiAuthUsers(networkId)
		shell.Display(metadata, "MerakiAuthUsers", cmd.Flags())
	},
}

var PostMerakiAuthUser = &cobra.Command{
	Use:   "merakiAuthUser",
	Short: "Create a user configured with Meraki Authentication for a network",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.MerakiAuthUser
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostMerakiAuthUser(networkId,  input)
		shell.Display(metadata, "MerakiAuthUser", cmd.Flags())
	},
}

var GetMerakiAuthUser = &cobra.Command{
	Use:   "merakiAuthUser",
	Short: "Return the Meraki Auth splash guest, RADIUS, or client VPN user.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		merakiAuthUserId := args[0]
		metadata := configure.GetMerakiAuthUser(networkId, merakiAuthUserId)
		shell.Display(metadata, "MerakiAuthUser", cmd.Flags())
	},
}

var DelMerakiAuthUser = &cobra.Command{
	Use:   "merakiAuthUser",
	Short: "Delete a user configured with Meraki Authentication",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		merakiAuthUserId := args[0]
		metadata := configure.DelMerakiAuthUser(networkId, merakiAuthUserId)
		shell.Display(metadata, "MerakiAuthUser", cmd.Flags())
	},
}

var PutMerakiAuthUser = &cobra.Command{
	Use:   "merakiAuthUser",
	Short: "Update a user configured with Meraki Authentication",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		var format configure.MerakiAuthUser
		input, _ := shell.ReadConfigFile(cmd, &format)
		merakiAuthUserId := args[0]
		metadata := configure.PutMerakiAuthUser(networkId, merakiAuthUserId,  input)
		shell.Display(metadata, "MerakiAuthUser", cmd.Flags())
	},
}




var GetMQTTBrokers = &cobra.Command{
	Use:   "mqttBrokers",
	Short: "List the MQTT brokers for this network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetMQTTBrokers(networkId)
		shell.Display(metadata, "MQTTBrokers", cmd.Flags())
	},
}


var GetMQTTBroker = &cobra.Command{
	Use:   "mqttBroker",
	Short: "Return an MQTT broker.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		mqttBrokerId := args[0]
		metadata := configure.GetMQTTBroker(networkId, mqttBrokerId)
		shell.Display(metadata, "MQTTBroker", cmd.Flags())
	},
}

var DelMQTTBroker = &cobra.Command{
	Use:   "mqttBroker",
	Short: "Delete a MQTT broker.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		mqttBrokerId := args[0]
		metadata := configure.DelMQTTBroker(networkId, mqttBrokerId)
		shell.Display(metadata, "MQTTBroker", cmd.Flags())
	},
}

var PutMQTTBroker = &cobra.Command{
	Use:   "mqttBroker",
	Short: "Return an MQTT broker.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		mqttBrokerId := args[0]
		var format configure.MQTTBroker
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutMQTTBroker(networkId, mqttBrokerId,  input)
		shell.Display(metadata, "MQTTBroker", cmd.Flags())
	},
}

var PostMQTTBroker = &cobra.Command{
	Use:   "mqttBroker",
	Short: "Create a MQTT broker.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.MQTTBroker
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostMQTTBroker(networkId,  input)
		shell.Display(metadata, "MQTTBroker", cmd.Flags())
	},
}

var GetNetflow = &cobra.Command{
	Use:   "netflow",
	Short: "Return the NetFlow traffic reporting settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetNetFlow(networkId)
		shell.Display(metadata, "Netflow", cmd.Flags())
	},
}

var PutNetflow = &cobra.Command{
	Use:   "netflow",
	Short: "Return the NetFlow traffic reporting settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.NetFlow
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutNetFlow(networkId,  input)
		shell.Display(metadata, "Netflow", cmd.Flags())
	},
}

var GetChannelUtilization = &cobra.Command{
	Use:   "channelUtilization",
	Short: "Get the channel utilization over each radio for all APs in a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		t0, _ := cmd.Flags().GetString("t0")
		t1, _ := cmd.Flags().GetString("t1")
		timespan, _ := cmd.Flags().GetString("timespan")
		resolution, _ := cmd.Flags().GetString("resolution")
		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")

		metadata := configure.GetChannelUtilization(networkId, t0, t1,
			timespan, resolution, perPage, startingAfter, endingBefore)
		shell.Display(metadata, "ChannelUtilization", cmd.Flags())
	},
}

var GetPIIRequests = &cobra.Command{
	Use:   "piiRequests",
	Short: "List the PII requests for this network or organization.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetPIIRequests(networkId)
		shell.Display(metadata, "PIIRequests", cmd.Flags())
	},
}

var GetPIIRequest = &cobra.Command{
	Use:   "piiRequest",
	Short: "Return a PII request.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		requestId := args[0]
		metadata := configure.GetPIIRequest(networkId, requestId)
		shell.Display(metadata, "PIIRequest", cmd.Flags())
	},
}

var DelPIIRequest = &cobra.Command{
	Use:   "piiRequest",
	Short: "Delete a PII request.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		requestId := args[0]
		metadata := configure.DelPIIRequest(networkId, requestId)
		shell.Display(metadata, "PIIRequest", cmd.Flags())
	},
}

var PostPIIRequest = &cobra.Command{
	Use:   "piiRequest",
	Short: "Create a PII request.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.PIIRequest
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostPIIRequest(networkId,  input)
		shell.Display(metadata, "PIIRequest", cmd.Flags())
	},
}

var GetSMDevices = &cobra.Command{
	Use:   "smDevices",
	Short: "Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		username, _ := cmd.Flags().GetString("username")
		email, _ := cmd.Flags().GetString("email")
		mac, _ := cmd.Flags().GetString("mac")
		serial, _ := cmd.Flags().GetString("serial")
		imei, _ := cmd.Flags().GetString("imei")
		bluetoothMac, _ := cmd.Flags().GetString("bluetoothMac")

		metadata := configure.GetSMDevices(networkId,
			username, email, mac, serial, imei, bluetoothMac)
		shell.Display(metadata, "SMDevices", cmd.Flags())
	},
}

var GetSMOwners = &cobra.Command{
	Use:   "smOwners",
	Short: "Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		username, _ := cmd.Flags().GetString("username")
		email, _ := cmd.Flags().GetString("email")
		mac, _ := cmd.Flags().GetString("mac")
		serial, _ := cmd.Flags().GetString("serial")
		imei, _ := cmd.Flags().GetString("imei")
		bluetoothMac, _ := cmd.Flags().GetString("bluetoothMac")

		metadata := configure.GetSMOwners(networkId,
			username, email, mac, serial, imei, bluetoothMac)
		shell.Display(metadata, "SMOwners", cmd.Flags())
	},
}


var GetSettings = &cobra.Command{
	Use:   "settings",
	Short: "Return the settings for a network.",
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
	Use:   "settings",
	Short: "Update the settings for a network.",
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

var GetSNMP = &cobra.Command{
	Use:   "snmp",
	Short: "Return the SNMP settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetSNMP(networkId)
		shell.Display(metadata, "SNMP", cmd.Flags())
	},
}

var PutSNMP = &cobra.Command{
	Use:   "snmp",
	Short: "Return the SNMP settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.SNMP
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutSNMP(networkId,  input)
		shell.Display(metadata, "SNMP", cmd.Flags())
	},
}

var GetSyslogServers = &cobra.Command{
	Use:   "syslogServers",
	Short: "List the syslog servers for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetSyslogServers(networkId)
		shell.Display(metadata, "SyslogServers", cmd.Flags())
	},
}

var PutSyslogServers = &cobra.Command{
	Use:   "syslogServers",
	Short: "Update the syslog servers for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.SyslogServers
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutSyslogServers(networkId,  input)
		shell.Display(metadata, "SyslogServers", cmd.Flags())
	},
}

var GetTrafficAnalysis = &cobra.Command{
	Use:   "trafficAnalysis",
	Short: "Return the traffic analysis settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetTrafficAnalysis(networkId)
		shell.Display(metadata, "TrafficAnalysis", cmd.Flags())
	},
}

var PutTrafficAnalysis = &cobra.Command{
	Use:   "trafficAnalysis",
	Short: "Update the traffic analysis settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.TrafficAnalysis
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutTrafficAnalysis(networkId,  input)
		shell.Display(metadata, "TrafficAnalysis", cmd.Flags())
	},
}

var GetTrafficShaping = &cobra.Command{
	Use:   "trafficShaping",
	Short: "Returns the application categories for traffic shaping rules.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetTrafficAnalysis(networkId)
		shell.Display(metadata, "TrafficShaping", cmd.Flags())
	},
}

var GetDSCPTaggingOptions = &cobra.Command{
	Use:   "dscpTaggingOptions",
	Short: "Returns the available DSCP tagging options for your traffic shaping rules.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetDSCPTaggingOptions(networkId)
		shell.Display(metadata, "DSCPTaggingOptions", cmd.Flags())
	},
}

var GetHTTPServers = &cobra.Command{
	Use:   "httpServers",
	Short: "List the HTTP servers for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetHTTPServers(networkId)
		shell.Display(metadata, "HTTPServers", cmd.Flags())
	},
}

var GetHTTPServer = &cobra.Command{
	Use:   "httpServer",
	Short: "Return an HTTP server for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		httpServerId := args[0]
		metadata := configure.GetHTTPServer(networkId, httpServerId)
		shell.Display(metadata, "HTTPServer", cmd.Flags())
	},
}

var PutHTTPServer = &cobra.Command{
	Use:   "httpServer",
	Short: "Update an HTTP server for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		var format configure.HTTPServer
		input, _ := shell.ReadConfigFile(cmd, &format)
		httpServerId := args[0]
		metadata := configure.PutHTTPServer(networkId, httpServerId,  input)
		shell.Display(metadata, "HTTPServer", cmd.Flags())
	},
}

var PostHTTPServer = &cobra.Command{
	Use:   "httpServer",
	Short: "Create a HTTP server for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		var format configure.HTTPServer
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostHTTPServer(networkId,  input)
		shell.Display(metadata, "HTTPServer", cmd.Flags())
	},
}

var DelHTTPServer = &cobra.Command{
	Use:   "httpServer",
	Short: "Delete an HTTP server for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		httpServerId := args[0]
		metadata := configure.DelHTTPServer(networkId, httpServerId)
		shell.Display(metadata, "HTTPServer", cmd.Flags())
	},
}

var GetWebhookTest = &cobra.Command{
	Use:   "webhookTest",
	Short: "Return the status of a webhook test for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		weebhookTestId := args[0]
		metadata := configure.GetWebhookTest(networkId, weebhookTestId)
		shell.Display(metadata, "WebhookTest", cmd.Flags())
	},
}

var PostWebhookTest = &cobra.Command{
	Use:   "webhookTest",
	Short: "Post  a webhook test for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.PostWebhookTest(networkId)
		shell.Display(metadata, "WebhookTest", cmd.Flags())
	},
}

var GetNetwork = &cobra.Command{
	Use:   "details",
	Short: "Return a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetNetwork(networkId)
		shell.Display(metadata, "Details", cmd.Flags())
	},
}

var PutNetwork = &cobra.Command{
	Use:   "details",
	Short: "Return a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.Network
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutNetwork(networkId,  input)
		shell.Display(metadata, "Details", cmd.Flags())
	},
}

var DelNetwork = &cobra.Command{
	Use:   "details",
	Short: "Delete a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.DelNetwork(networkId)
		shell.Display(metadata, "Details", cmd.Flags())
	},
}

var PostBindNetwork = &cobra.Command{
	Use:   "bindNetwork",
	Short: "Bind a network to a template.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.Network
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostBindNetwork(networkId,  input)
		shell.Display(metadata, "BindNetwork", cmd.Flags())
	},
}


var PostUnBindNetwork = &cobra.Command{
	Use:   "unbindNetwork",
	Short: "UnBind a network from a template.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.Network
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostUnBindNetwork(networkId,  input)
		shell.Display(metadata, "UnBindNetwork", cmd.Flags())
	},
}


var PostSplitNetwork = &cobra.Command{
	Use:   "splitNetwork",
	Short: "Split a network from a template.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.Network
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostSplitNetwork(networkId,  input)
		shell.Display(metadata, "SplitNetwork", cmd.Flags())
	},
}

var GetClients = &cobra.Command{
	Use:   "clients",
	Short: "Return the policy assigned to a client on the network",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		t0, _ := cmd.Flags().GetString("t0")
		t1, _ := cmd.Flags().GetString("t1")
		timespan, _ := cmd.Flags().GetString("timespan")
		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")

		metadata := organizations.GetClients(networkId, t0, t1,
			timespan, perPage, startingAfter, endingBefore)
		shell.Display(metadata, "clients", cmd.Flags())
	},
}