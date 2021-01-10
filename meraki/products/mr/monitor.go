package mr

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/wireless/monitor"
	shell "github.com/ddexterpark/merakictl/utilities"
	"github.com/spf13/cobra"
)

var GetAirMarshal = &cobra.Command{
	Use:   "AirMarshal",
	Short: "List Air Marshal scan results from a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		t0, _ := cmd.Flags().GetString("t0")
		timespan, _ := cmd.Flags().GetString("timespan")
		metadata := monitor.GetAirMarshal(networkId, t0, timespan)
			shell.Display(metadata, "AirMarshal", cmd.Flags())
	},
}



var GetChannelUtilizationHistory = &cobra.Command{
	Use:   "ChannelUtilizationHistory",
	Short: "Return AP channel utilization over time for a device or network client.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		t0, _ := cmd.Flags().GetString("t0")
		t1, _ := cmd.Flags().GetString("t1")
		timespan, _ := cmd.Flags().GetString("timespan")
		resolution, _ := cmd.Flags().GetString("resolution")
		autoResolution, _ := cmd.Flags().GetString("autoResolution")
		clientId, _ := cmd.Flags().GetString("clientId")
		deviceSerial,_ := cmd.Flags().GetString("deviceSerial")
		apTag,_ := cmd.Flags().GetString("apTag")
		band,_ := cmd.Flags().GetString("band")

		metadata := monitor.GetChannelUtilizationHistory(networkId, t0, t1, timespan, resolution, autoResolution, clientId, deviceSerial, apTag, band)
		shell.Display(metadata, "ChannelUtilizationHistory", cmd.Flags())
	},
}


var GetClientCountHistory = &cobra.Command{
Use:   "ClientCountHistory",
Short: "Return mr client counts over time for a network, device, or network client.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

t0, _ := cmd.Flags().GetString("t0")
t1, _ := cmd.Flags().GetString("t1")
timespan, _ := cmd.Flags().GetString("timespan")
resolution, _ := cmd.Flags().GetString("resolution")
autoResolution, _ := cmd.Flags().GetString("autoResolution")
clientId, _ := cmd.Flags().GetString("clientId")
deviceSerial,_ := cmd.Flags().GetString("deviceSerial")
apTag,_ := cmd.Flags().GetString("apTag")
band,_ := cmd.Flags().GetString("band")
ssid,_ := cmd.Flags().GetString("ssid")
metadata := monitor.GetClientCountHistory(networkId, t0, t1, timespan, resolution, autoResolution, clientId, deviceSerial, apTag, band, ssid)
shell.Display(metadata, "ClientCountHistory", cmd.Flags())
},
}


var GetConnectionStat = &cobra.Command{
Use:   "ConnectionStat",
Short: "Aggregated connectivity info for a given client on this network. Clients are identified by their MAC.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

clientId, _ := cmd.Flags().GetString("clientId")
t0, _ := cmd.Flags().GetString("t0")
t1, _ := cmd.Flags().GetString("t1")
timespan, _ := cmd.Flags().GetString("timespan")
vlan,_ := cmd.Flags().GetString("vlan")
apTag,_ := cmd.Flags().GetString("apTag")
band,_ := cmd.Flags().GetString("band")
ssid,_ := cmd.Flags().GetString("ssid")
metadata := monitor.GetConnectionStat(networkId, clientId, t0, t1, timespan, band, ssid, vlan, apTag)
shell.Display(metadata, "ConnectionStat", cmd.Flags())
},
}


var GetConnectivityEvents = &cobra.Command{
Use:   "ConnectivityEvents",
Short: "List the mr connectivity events for a client within a network in the timespan.",
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
types, _ := cmd.Flags().GetString("types")
includedSeverities, _ := cmd.Flags().GetString("includedSeverities")
clientId, _ := cmd.Flags().GetString("clientId")
deviceSerial,_ := cmd.Flags().GetString("deviceSerial")
ssidNumber,_ := cmd.Flags().GetString("ssidNumber")
band,_ := cmd.Flags().GetString("band")

metadata := monitor.GetConnectivityEvents(networkId,
	clientId, perPage, startingAfter, endingBefore, t0, t1,
	timespan, types, includedSeverities, band, ssidNumber, deviceSerial)
shell.Display(metadata, "ConnectivityEvents", cmd.Flags())
},
}




var GetLatencyHistory = &cobra.Command{
Use:   "LatencyHistory",
Short: "Return the latency history for a client. ",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}


t0, _ := cmd.Flags().GetString("t0")
t1, _ := cmd.Flags().GetString("t1")
timespan, _ := cmd.Flags().GetString("timespan")
accessCategory, _ := cmd.Flags().GetString("accessCategory")
resolution, _ := cmd.Flags().GetString("resolution")
autoResolution, _ := cmd.Flags().GetString("autoResolution")
clientId, _ := cmd.Flags().GetString("clientId")
deviceSerial,_ := cmd.Flags().GetString("deviceSerial")
apTag,_ := cmd.Flags().GetString("apTag")
band,_ := cmd.Flags().GetString("band")
ssid,_ := cmd.Flags().GetString("ssid")
metadata := monitor.GetLatencyHistory(networkId, t0, t1, timespan,
	resolution, autoResolution, clientId, deviceSerial, apTag, band, ssid, accessCategory)
shell.Display(metadata, "LatencyHistory", cmd.Flags())
},
}




var GetLatencyStat = &cobra.Command{
Use:   "LatencyStat",
Short: "Aggregated latency info for a given client on this network. Clients are identified by their MAC.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

t0, _ := cmd.Flags().GetString("t0")
t1, _ := cmd.Flags().GetString("t1")
timespan, _ := cmd.Flags().GetString("timespan")
resolution, _ := cmd.Flags().GetString("resolution")
clientId, _ := cmd.Flags().GetString("clientId")
metadata := monitor.GetClientLatencyHistory(networkId, clientId, t0, t1, timespan, resolution)
shell.Display(metadata, "LatencyStat", cmd.Flags())
},
}




var GetLatencyStats = &cobra.Command{
Use:   "LatencyStats",
Short: "Aggregated latency info for this network, grouped by clients.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}


t0, _ := cmd.Flags().GetString("t0")
t1, _ := cmd.Flags().GetString("t1")
timespan, _ := cmd.Flags().GetString("timespan")
fields, _ := cmd.Flags().GetString("fields")
vlan, _ := cmd.Flags().GetString("vlan")
apTag,_ := cmd.Flags().GetString("apTag")
band,_ := cmd.Flags().GetString("band")
ssid,_ := cmd.Flags().GetString("ssid")
metadata := monitor.GetAggregatedLatencies(networkId, t0, t1, timespan, band, ssid, vlan, apTag, fields)
shell.Display(metadata, "LatencyStats", cmd.Flags())
},
}



var GetDeviceConnectionStats = &cobra.Command{
Use:   "DeviceConnectionStats",
Short: "Aggregated latency info for a given AP on this network.",
Run: func(cmd *cobra.Command, args []string) {
_, _, serial := shell.ResolveFlags(cmd.Flags())
if serial == "" {
	serial = args[0]
}

t0, _ := cmd.Flags().GetString("t0")
t1, _ := cmd.Flags().GetString("t1")
timespan, _ := cmd.Flags().GetString("timespan")
apTag,_ := cmd.Flags().GetString("apTag")
band,_ := cmd.Flags().GetString("band")
ssid,_ := cmd.Flags().GetString("ssid")
vlan, _ := cmd.Flags().GetString("vlan")

metadata := monitor.GetDeviceConnectionStats(serial, t0, t1, timespan, band, ssid, vlan, apTag)
shell.Display(metadata, "DeviceConnectionStats", cmd.Flags())
},
}


var GetNetworkConnectionStats = &cobra.Command{
Use:   "NetworkConnectionStats",
Short: "Aggregated connectivity info for this network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

t0, _ := cmd.Flags().GetString("t0")
t1, _ := cmd.Flags().GetString("t1")
timespan, _ := cmd.Flags().GetString("timespan")
apTag,_ := cmd.Flags().GetString("apTag")
band,_ := cmd.Flags().GetString("band")
ssid,_ := cmd.Flags().GetString("ssid")
vlan, _ := cmd.Flags().GetString("vlan")
fields, _ := cmd.Flags().GetString("fields")

metadata := monitor.GetAggregatedLatencies(networkId,
	t0, t1, timespan, band, ssid, vlan, apTag, fields)
shell.Display(metadata, "NetworkConnectionStats", cmd.Flags())
},
}

var GetDataRateHistory = &cobra.Command{
Use:   "DataRateHistory",
Short: "Return PHY data rates over time for a network, device, or network client.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

t0, _ := cmd.Flags().GetString("t0")
t1, _ := cmd.Flags().GetString("t1")
timespan, _ := cmd.Flags().GetString("timespan")
resolution, _ := cmd.Flags().GetString("resolution")
autoResolution, _ := cmd.Flags().GetString("autoResolution")
clientId, _ := cmd.Flags().GetString("clientId")
deviceSerial,_ := cmd.Flags().GetString("deviceSerial")
apTag,_ := cmd.Flags().GetString("apTag")
band,_ := cmd.Flags().GetString("band")
ssid,_ := cmd.Flags().GetString("ssid")

metadata := monitor.GetDataRateHistory(networkId, t0, t1,
	timespan, resolution, autoResolution, clientId, deviceSerial, apTag, band, ssid)
shell.Display(metadata, "DataRateHistory", cmd.Flags())
},
}



var GetConnectionStats = &cobra.Command{
Use:   "ConnectionStats",
Short: "Aggregated connectivity info for this network, grouped by node.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

t0, _ := cmd.Flags().GetString("t0")
t1, _ := cmd.Flags().GetString("t1")
timespan, _ := cmd.Flags().GetString("timespan")
clientId, _ := cmd.Flags().GetString("clientId")
apTag,_ := cmd.Flags().GetString("apTag")
band,_ := cmd.Flags().GetString("band")
ssid,_ := cmd.Flags().GetString("ssid")
vlan, _ := cmd.Flags().GetString("vlan")
metadata := monitor.GetConnectionStats(networkId,
	clientId, t0, t1, timespan, band, ssid, vlan, apTag)
shell.Display(metadata, "ConnectionStats", cmd.Flags())
},
}

var GetDeviceLatencyStats = &cobra.Command{
Use:   "DeviceLatencyStats",
Short: "Aggregated latency info for this network, grouped by node.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

t0, _ := cmd.Flags().GetString("t0")
t1, _ := cmd.Flags().GetString("t1")
timespan, _ := cmd.Flags().GetString("timespan")
apTag,_ := cmd.Flags().GetString("apTag")
band,_ := cmd.Flags().GetString("band")
ssid,_ := cmd.Flags().GetString("ssid")
fields, _ := cmd.Flags().GetString("fields")
vlan, _ := cmd.Flags().GetString("vlan")
metadata := monitor.GetAggregatedLatencies(networkId, t0, t1, timespan, band, ssid, vlan, apTag, fields)
shell.Display(metadata, "DeviceLatencyStats", cmd.Flags())
},
}

var GetFailedConnections = &cobra.Command{
Use:   "FailedConnections",
Short: "List of all failed client connection events on this network in a given time range.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

t0, _ := cmd.Flags().GetString("t0")
t1, _ := cmd.Flags().GetString("t1")
timespan, _ := cmd.Flags().GetString("timespan")
apTag,_ := cmd.Flags().GetString("apTag")
band,_ := cmd.Flags().GetString("band")
ssid,_ := cmd.Flags().GetString("ssid")
vlan, _ := cmd.Flags().GetString("vlan")
metadata := monitor.GetFailedConnections(networkId, t0, t1, timespan, band, ssid, vlan, apTag)
shell.Display(metadata, "FailedConnections", cmd.Flags())
},
}

var GetLatencyHistoryAverage = &cobra.Command{
Use:   "LatencyHistoryAverage",
Short: "Return average mr latency over time for a network, device, or network client.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

t0, _ := cmd.Flags().GetString("t0")
t1, _ := cmd.Flags().GetString("t1")
timespan, _ := cmd.Flags().GetString("timespan")

resolution, _ := cmd.Flags().GetString("resolution")
autoResolution, _ := cmd.Flags().GetString("autoResolution")
clientId, _ := cmd.Flags().GetString("clientId")
deviceSerial,_ := cmd.Flags().GetString("deviceSerial")
apTag,_ := cmd.Flags().GetString("apTag")
band,_ := cmd.Flags().GetString("band")
ssid,_ := cmd.Flags().GetString("ssid")
accessCategory, _ := cmd.Flags().GetString("accessCategory")

metadata := monitor.GetLatencyHistory(networkId,
	t0, t1, timespan, resolution, autoResolution,
	clientId, deviceSerial, apTag, band, ssid, accessCategory)
shell.Display(metadata, "LatencyHistoryAverage", cmd.Flags())
},
}

var GetAggregatedLatencies = &cobra.Command{
Use:   "AggregatedLatencies",
Short: "Aggregated latency info for a given AP on this network.",
Run: func(cmd *cobra.Command, args []string) {
_, _, serial := shell.ResolveFlags(cmd.Flags())
if serial == "" {
	serial = args[0]
}

t0, _ := cmd.Flags().GetString("t0")
t1, _ := cmd.Flags().GetString("t1")
timespan, _ := cmd.Flags().GetString("timespan")
apTag,_ := cmd.Flags().GetString("apTag")
band,_ := cmd.Flags().GetString("band")
ssid,_ := cmd.Flags().GetString("ssid")
fields, _ := cmd.Flags().GetString("fields")
vlan, _ := cmd.Flags().GetString("vlan")
metadata := monitor.GetAggregatedLatencies(serial, t0,
	t1, timespan, band, ssid, vlan, apTag, fields)
shell.Display(metadata, "AggregatedLatencies", cmd.Flags())
},
}

var GetNetworkLatencyStats = &cobra.Command{
Use:   "NetworkLatencyStats",
Short: "Aggregated latency info for this network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

t0, _ := cmd.Flags().GetString("t0")
t1, _ := cmd.Flags().GetString("t1")
timespan, _ := cmd.Flags().GetString("timespan")
apTag,_ := cmd.Flags().GetString("apTag")
band,_ := cmd.Flags().GetString("band")
ssid,_ := cmd.Flags().GetString("ssid")
fields, _ := cmd.Flags().GetString("fields")
vlan, _ := cmd.Flags().GetString("vlan")
metadata := monitor.GetAggregatedLatencies(networkId,
	t0, t1, timespan, band, ssid, vlan, apTag, fields)
shell.Display(metadata, "NetworkLatencyStats", cmd.Flags())
},
}

var GetMeshStatuses = &cobra.Command{
Use:   "MeshStatuses",
Short: "List mr mesh statuses for repeaters.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

perPage, _ := cmd.Flags().GetString("perPage")
startingAfter, _ := cmd.Flags().GetString("startingAfter")
endingBefore, _ := cmd.Flags().GetString("endingBefore")

metadata := monitor.GetMeshStatuses(networkId, perPage, startingAfter, endingBefore)
shell.Display(metadata, "MeshStatuses", cmd.Flags())
},
}


var GetSignalQualityHistory = &cobra.Command{
Use:   "SignalQualityHistory",
Short: "Return signal quality (SNR/RSSI) over time for a device or network client.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

t0, _ := cmd.Flags().GetString("t0")
t1, _ := cmd.Flags().GetString("t1")
timespan, _ := cmd.Flags().GetString("timespan")
resolution, _ := cmd.Flags().GetString("resolution")
autoResolution, _ := cmd.Flags().GetString("autoResolution")
clientId, _ := cmd.Flags().GetString("clientId")
deviceSerial,_ := cmd.Flags().GetString("deviceSerial")
apTag,_ := cmd.Flags().GetString("apTag")
band,_ := cmd.Flags().GetString("band")
ssid,_ := cmd.Flags().GetString("ssid")

metadata := monitor.GetSignalQualityHistory(networkId, t0, t1,
	timespan, resolution, autoResolution, clientId,
	deviceSerial, apTag, band, ssid)
shell.Display(metadata, "SignalQualityHistory", cmd.Flags())
},
}



var GetStatus = &cobra.Command{
Use:   "Status",
Short: "Return the SSID statuses of an access point.",
Run: func(cmd *cobra.Command, args []string) {
_, _, serial := shell.ResolveFlags(cmd.Flags())
if serial == "" {
	serial = args[0]
}

metadata := monitor.GetStatus(serial)
shell.Display(metadata, "Status", cmd.Flags())
},
}


var GetUsageHistory = &cobra.Command{
Use:   "UsageHistory",
Short: "Return AP usage over time for a device or network client.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

t0, _ := cmd.Flags().GetString("t0")
t1, _ := cmd.Flags().GetString("t1")
timespan, _ := cmd.Flags().GetString("timespan")
resolution, _ := cmd.Flags().GetString("resolution")
autoResolution, _ := cmd.Flags().GetString("autoResolution")
clientId, _ := cmd.Flags().GetString("clientId")
deviceSerial,_ := cmd.Flags().GetString("deviceSerial")
apTag,_ := cmd.Flags().GetString("apTag")
band,_ := cmd.Flags().GetString("band")
ssid,_ := cmd.Flags().GetString("ssid")
metadata := monitor.GetUsageHistory(networkId, t0,
	t1, timespan, resolution, autoResolution, clientId,
	deviceSerial, apTag, band, ssid)
shell.Display(metadata, "UsageHistory", cmd.Flags())
},
}