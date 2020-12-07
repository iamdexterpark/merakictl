package network

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/networks/monitor"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

var bluetoothclients = &cobra.Command{
	Use:   "bluetoothclients",
	Short: "List the Bluetooth clients seen by APs in this network.",
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
		includeConnectivityHistory, _ := cmd.Flags().GetString("includeConnectivityHistory")


		metadata := monitor.GetBluetoothClients(networkId,t0,t1,timespan,
			perPage,startingAfter,endingBefore,includeConnectivityHistory)
		shell.Display(metadata, "bluetoothclients", cmd.Flags())
	},
}

var bluetoothclient = &cobra.Command{
	Use:   "bluetoothclient",
	Short: "Return a Bluetooth client. Bluetooth clients can be identified by their ID or their MAC.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		bluetoothClientId := args[0]
		includeConnectivityHistory, _ := cmd.Flags().GetString("includeConnectivityHistory")
		connectivityHistoryTimespan, _ := cmd.Flags().GetString("connectivityHistoryTimespan")

		metadata := monitor.GetBluetoothClient(networkId, bluetoothClientId,
			includeConnectivityHistory, connectivityHistoryTimespan)
		shell.Display(metadata, "bluetoothclient", cmd.Flags())
	},
}


var clienttraffichistory = &cobra.Command{
	Use:   "clienttraffichistory",
	Short: "Return the client's network traffic data over time. Usage data is in kilobytes.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		clientId := args[0]
		trafficHistory, _ := cmd.Flags().GetString("trafficHistory")
		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")

		metadata := monitor.GetClientNetworkTraffic(networkId, clientId,
			trafficHistory,perPage,startingAfter,endingBefore)
		shell.Display(metadata, "clienttraffichistory", cmd.Flags())
	},
}

var clientusagehistory = &cobra.Command{
	Use:   "clientusagehistory",
	Short: "Return the client's daily usage history.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		clientId := args[0]
		metadata := monitor.GetClientUsageHistory(networkId, clientId)
		shell.Display(metadata, "clientusagehistory", cmd.Flags())
	},
}

var networkclients = &cobra.Command{
	Use:   "networkclients",
	Short: "List the clients that have used this network in the timespan.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		t0, _ := cmd.Flags().GetString("t0")
		timespan, _ := cmd.Flags().GetString("timespan")
		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")
		metadata := monitor.GetClients(networkId,t0,timespan,perPage, startingAfter,endingBefore)
		shell.Display(metadata, "networkclients", cmd.Flags())
	},
}

var clientidentifier = &cobra.Command{
	Use:   "clientidentifier",
	Short: "Return the client associated with the given identifier.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		clientId := args[0]
		metadata := monitor.GetAssociatedClient(networkId, clientId)
		shell.Display(metadata, "clientidentifier", cmd.Flags())
	},
}

var environmentalevents = &cobra.Command{
	Use:   "environmentalevents",
	Short: "List the environmental events for the network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		includedEventTypes, _ := cmd.Flags().GetString("includedEventTypes")
		excludedEventTypes, _ := cmd.Flags().GetString("excludedEventTypes")
		sensorSerial, _ := cmd.Flags().GetString("sensorSerial")
		gatewaySerial, _ := cmd.Flags().GetString("gatewaySerials")
		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")
		metadata := monitor.GetEnvironmentalEvents(networkId,includedEventTypes,
			excludedEventTypes,sensorSerial, gatewaySerial, perPage, startingAfter, endingBefore)
		shell.Display(metadata, "networkclients", cmd.Flags())
	},
}

var events = &cobra.Command{
	Use:   "events",
	Short: "List the events for the network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		productType, _ := cmd.Flags().GetString("productType")
		includedEventTypes, _ := cmd.Flags().GetString("includedEventTypes")
		excludedEventTypes, _ := cmd.Flags().GetString("excludedEventTypes")
		deviceMac, _ := cmd.Flags().GetString("deviceMac")
		deviceSerial, _ := cmd.Flags().GetString("deviceSerial")
		deviceName, _ := cmd.Flags().GetString("deviceName")
		clientIp, _ := cmd.Flags().GetString("clientIp")
		clientMac, _ := cmd.Flags().GetString("clientMac")
		clientName, _ := cmd.Flags().GetString("clientName")
		smDeviceMac, _ := cmd.Flags().GetString("smDeviceMac")
		smDeviceName, _ := cmd.Flags().GetString("smDeviceName")
		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")

		metadata := monitor.GetEvents(networkId, productType, includedEventTypes,
			excludedEventTypes,deviceMac, deviceSerial,deviceName, clientIp,
			clientMac, clientName,smDeviceMac,smDeviceName,
			perPage, startingAfter, endingBefore)
		shell.Display(metadata, "events", cmd.Flags())
	},
}

var splashloginattempts = &cobra.Command{
	Use:   "splashloginattempts",
	Short: "List the splash login attempts for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		splashLoginAttempts, _ := cmd.Flags().GetString("splashLoginAttempts")
		ssidNumber, _ := cmd.Flags().GetString("ssidNumber")
		loginIdentifier, _ := cmd.Flags().GetString("loginIdentifier")
		timespan, _ := cmd.Flags().GetString("timespan")

		metadata := monitor.GetSplashLoginAttempts(networkId, splashLoginAttempts,
			ssidNumber,loginIdentifier, timespan)
		shell.Display(metadata, "splashloginattempts", cmd.Flags())
	},
}


var traffic = &cobra.Command{
	Use:   "traffic",
	Short: "Return the traffic analysis data for this network. Traffic analysis with hostname visibility must be enabled on the network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		t0, _ := cmd.Flags().GetString("t0")
		timespan, _ := cmd.Flags().GetString("timespan")
		deviceType, _ := cmd.Flags().GetString("deviceType")

		metadata := monitor.GetTrafficAnalysis(networkId,t0,timespan,deviceType)
		shell.Display(metadata, "traffic", cmd.Flags())
	},
}