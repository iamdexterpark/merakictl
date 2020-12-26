package network

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/networks/configure"
	organizations "github.com/ddexterpark/dashboard-api-golang/api/general/organizations/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

var details = &cobra.Command{
	Use:   "details",
	Short: "Return a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetNetwork(networkId)
		shell.Display(metadata, "details", cmd.Flags())
	},
}

// clients
var clients = &cobra.Command{
	Use:   "clients",
	Short: "List the clients that have used this network in the timespan.",
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

var alertconfig = &cobra.Command{
	Use:   "alertconfig",
	Short: "Return The Alert Configuration For This Network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetAlertSettings(networkId)
		shell.Display(metadata, "alertconfig", cmd.Flags())
	},
}

var clientpolicy = &cobra.Command{
	Use:   "clientpolicy",
	Short: "Return the policy assigned to a client on the network. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		clientId := args[0]
		metadata := configure.GetClientPolicy(networkId, clientId)
		shell.Display(metadata, "clientpolicy", cmd.Flags())
	},
}

var clientsplashauthorization = &cobra.Command{
	Use:   "clientsplashauthorization",
	Short: "Return the splash authorization for a client, for each SSID they've associated with through splash. Only enabled SSIDs with Click-through splash enabled will be included. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		clientId := args[0]
		metadata := configure.GetSplashAuthorization(networkId, clientId)
		shell.Display(metadata, "clientsplashauthorization", cmd.Flags())
	},
}

var devices = &cobra.Command{
	Use:   "devices",
	Short: "List the devices in a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetDevices(networkId)
		shell.Display(metadata, "devices", cmd.Flags())
	},
}

var firmwareupgrades = &cobra.Command{
	Use:   "firmwareupgrades",
	Short: "Get current maintenance window for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetFirmwareUpgrades(networkId)
		shell.Display(metadata, "firmwareupgrades", cmd.Flags())
	},
}

var floorplans = &cobra.Command{
	Use:   "floorplans",
	Short: "List the floor plans that belong to your network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetFloorPlans(networkId)
		shell.Display(metadata, "floorplans", cmd.Flags())
	},
}

var floorplan = &cobra.Command{
	Use:   "floorplan",
	Short: "Find a floor plan by ID.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		floorplanId := args[0]
		metadata := configure.GetFloorPlan(networkId, floorplanId)
		shell.Display(metadata, "floorplan", cmd.Flags())
	},
}

var grouppolicies = &cobra.Command{
	Use:   "grouppolicies",
	Short: "List the group policies in a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetGroupPolicies(networkId)
		shell.Display(metadata, "grouppolicies", cmd.Flags())
	},
}

var grouppolicy = &cobra.Command{
	Use:   "grouppolicy",
	Short: "Display a group policy.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		groupPolicyId := args[0]
		metadata := configure.GetGroupPolicy(networkId, groupPolicyId)
		shell.Display(metadata, "grouppolicy", cmd.Flags())
	},
}

var merakiauthusers = &cobra.Command{
	Use:   "merakiauthusers",
	Short: "List the users configured under Meraki Authentication for a network (splash guest or RADIUS users for a wireless network, or client VPN users for a wired network).",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetMerakiAuthUsers(networkId)
		shell.Display(metadata, "merakiauthusers", cmd.Flags())
	},
}


var merakiauthuser = &cobra.Command{
	Use:   "merakiauthuser",
	Short: "Return the Meraki Auth splash guest, RADIUS, or client VPN user.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		merakiAuthUserId := args[0]
		metadata := configure.GetMerakiAuthUser(networkId, merakiAuthUserId)
		shell.Display(metadata, "merakiauthuser", cmd.Flags())
	},
}

var mqttbrokers = &cobra.Command{
	Use:   "mqttbrokers",
	Short: "List the MQTT brokers for this network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetMqttBrokers(networkId)
		shell.Display(metadata, "mqttbrokers", cmd.Flags())
	},
}


var mqttbroker = &cobra.Command{
	Use:   "mqttbroker",
	Short: "Return an MQTT broker.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		mqttBrokerId := args[0]
		metadata := configure.GetMqttBroker(networkId, mqttBrokerId)
		shell.Display(metadata, "mqttbroker", cmd.Flags())
	},
}

var netflow = &cobra.Command{
	Use:   "netflow",
	Short: "Return the NetFlow traffic reporting settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetNetFlow(networkId)
		shell.Display(metadata, "netflow", cmd.Flags())
	},
}

var channelutilization = &cobra.Command{
	Use:   "channelutilization",
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
		shell.Display(metadata, "channelutilization", cmd.Flags())
	},
}

var piirequests = &cobra.Command{
	Use:   "piirequests",
	Short: "List the PII requests for this network or organization.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetPiiRequests(networkId)
		shell.Display(metadata, "piirequests", cmd.Flags())
	},
}

var piirequest = &cobra.Command{
	Use:   "piirequest",
	Short: "Return a PII request.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		requestId := args[0]
		metadata := configure.GetPiiRequest(networkId, requestId)
		shell.Display(metadata, "piirequest", cmd.Flags())
	},
}

var smdevices = &cobra.Command{
	Use:   "smdevices",
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

		metadata := configure.GetSmDevicesForKey(networkId,
			username, email, mac, serial,imei, bluetoothMac)
		shell.Display(metadata, "smdevices", cmd.Flags())
	},
}

var smowners = &cobra.Command{
	Use:   "smowners",
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

		metadata := configure.GetSmOwnersForKey(networkId,
			username, email, mac, serial,imei, bluetoothMac)
		shell.Display(metadata, "smowners", cmd.Flags())
	},
}

var settings = &cobra.Command{
	Use:   "settings",
	Short: "Return the settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetSettings(networkId)
		shell.Display(metadata, "settings", cmd.Flags())
	},
}

var snmp = &cobra.Command{
	Use:   "snmp",
	Short: "Return the SNMP settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetSNMP(networkId)
		shell.Display(metadata, "snmp", cmd.Flags())
	},
}

var syslog = &cobra.Command{
	Use:   "syslog",
	Short: "List the syslog servers for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetSyslogServers(networkId)
		shell.Display(metadata, "syslog", cmd.Flags())
	},
}

var trafficanalysis = &cobra.Command{
	Use:   "trafficanalysis",
	Short: "Return the traffic analysis settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetTrafficAnalysis(networkId)
		shell.Display(metadata, "trafficanalysis", cmd.Flags())
	},
}

var trafficshaping = &cobra.Command{
	Use:   "trafficshaping",
	Short: "Returns the application categories for traffic shaping rules.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetTrafficAnalysis(networkId)
		shell.Display(metadata, "trafficshaping", cmd.Flags())
	},
}

var dscp = &cobra.Command{
	Use:   "dscp",
	Short: "Returns the available DSCP tagging options for your traffic shaping rules.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetDscpTaggingOptions(networkId)
		shell.Display(metadata, "dscp", cmd.Flags())
	},
}

var httpservers = &cobra.Command{
	Use:   "httpservers",
	Short: "List the HTTP servers for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetHTTPServers(networkId)
		shell.Display(metadata, "httpservers", cmd.Flags())
	},
}

var httpserver = &cobra.Command{
	Use:   "httpserver",
	Short: "Return an HTTP server for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		httpServerId := args[0]
		metadata := configure.GetHTTPServer(networkId, httpServerId)
		shell.Display(metadata, "httpserver", cmd.Flags())
	},
}

var webhooktest = &cobra.Command{
	Use:   "webhooktest",
	Short: "Return the status of a webhook test for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		weebhookTestId := args[0]
		metadata := configure.GetWebhookTests(networkId, weebhookTestId)
		shell.Display(metadata, "webhooktest", cmd.Flags())
	},
}

