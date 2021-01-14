package network

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/networks/monitor"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)

var GetBluetoothClients = &cobra.Command{
	Use:   "bluetoothClients",
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


		metadata := monitor.GetBluetoothClients(networkId, t0, t1, timespan,
			perPage, startingAfter, endingBefore, includeConnectivityHistory)
		shell.Display(metadata, "BluetoothClients", cmd.Flags())
	},
}

var GetBluetoothClient = &cobra.Command{
	Use:   "bluetoothClient",
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
		shell.Display(metadata, "BluetoothClient", cmd.Flags())
	},
}


var GetClientTrafficHistory = &cobra.Command{
	Use:   "clientTrafficHistory",
	Short: "Return the client's network traffic data over time. Usage data is in kilobytes.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		clientId := args[0]
		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")

		metadata := monitor.GetClientTrafficHistory(networkId, clientId, perPage, startingAfter, endingBefore)
		shell.Display(metadata, "ClientTrafficHistory", cmd.Flags())
	},
}

var GetClientUsageHistory = &cobra.Command{
	Use:   "clientUsageHistory",
	Short: "Return the client's daily usage history.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		clientId := args[0]
		metadata := monitor.GetClientUsageHistory(networkId, clientId)
		shell.Display(metadata, "ClientUsageHistory", cmd.Flags())
	},
}

var GetNetworkClients = &cobra.Command{
	Use:   "networkClients",
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
		metadata := monitor.GetClients(networkId, t0, timespan, perPage, startingAfter, endingBefore)
		shell.Display(metadata, "NetworkClients", cmd.Flags())
	},
}



var GetNetworkClient = &cobra.Command{
	Use:   "clientIdentifier",
	Short: "Return the client associated with the given identifier.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		clientId := args[0]
		metadata := monitor.GetClient(networkId, clientId)
		shell.Display(metadata, "ClientIdentifier", cmd.Flags())
	},
}

var GetEnvironmentalEvents = &cobra.Command{
	Use:   "environmentalEvents",
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
		metadata := monitor.GetEnvironmentalEvents(networkId, includedEventTypes,
			excludedEventTypes, sensorSerial, gatewaySerial, perPage, startingAfter, endingBefore)
		shell.Display(metadata, "EnvironmentalEvents", cmd.Flags())
	},
}

var GetEvents = &cobra.Command{
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
			excludedEventTypes, deviceMac, deviceSerial, deviceName, clientIp,
			clientMac, clientName, smDeviceMac, smDeviceName,
			perPage, startingAfter, endingBefore)
		shell.Display(metadata, "Events", cmd.Flags())
	},
}

var GetSplashLoginAttempts = &cobra.Command{
	Use:   "splashLoginAttempts",
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
			ssidNumber, loginIdentifier, timespan)
		shell.Display(metadata, "SplashLoginAttempts", cmd.Flags())
	},
}


var GetTraffic = &cobra.Command{
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

		metadata := monitor.GetTraffic(networkId, t0, timespan, deviceType)
		shell.Display(metadata, "Traffic", cmd.Flags())
	},
}