package wireless

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/wireless/monitor"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

var airmarshal = &cobra.Command{
	Use:   "airmarshal",
	Short: "List Air Marshal scan results from a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		t0, _ := cmd.Flags().GetString("t0")
		timespan, _ := cmd.Flags().GetString("timespan")
		metadata := monitor.GetAirMarshalScanResults(networkId, t0, timespan)
			shell.Display(metadata, "airmarshal", cmd.Flags())
	},
}



var channelutilizationhistory = &cobra.Command{
	Use:   "channelutilizationhistory",
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

		metadata := monitor.GetAPChannelUtilization(networkId, t0, t1, timespan, resolution, autoResolution, clientId, deviceSerial,apTag,band)
		shell.Display(metadata, "channelutilizationhistory", cmd.Flags())
	},
}


var clientcounthistory = &cobra.Command{
Use:   "clientcounthistory",
Short: "Return wireless client counts over time for a network, device, or network client.",
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
metadata := monitor.GetWirelessClientCount(networkId,t0,t1, timespan, resolution, autoResolution, clientId, deviceSerial, apTag, band, ssid)
shell.Display(metadata, "clientcounthistory", cmd.Flags())
},
}


var connectionstat = &cobra.Command{
Use:   "connectionstat",
Short: "Aggregated connectivity info for a given client on this network. Clients are identified by their MAC.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}


t0, _ := cmd.Flags().GetString("t0")
t1, _ := cmd.Flags().GetString("t1")
timespan, _ := cmd.Flags().GetString("timespan")
vlan,_ := cmd.Flags().GetString("vlan")
apTag,_ := cmd.Flags().GetString("apTag")
band,_ := cmd.Flags().GetString("band")
ssid,_ := cmd.Flags().GetString("ssid")
metadata := monitor.GetAggregatedConnectivity(networkId, t0, t1, timespan, band, ssid, vlan, apTag)
shell.Display(metadata, "connectionstat", cmd.Flags())
},
}




var connectionstats = &cobra.Command{
Use:   "connectionstats",
Short: "Aggregated connectivity info for this network, grouped by clients.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

t0, _ := cmd.Flags().GetString("t0")
t1, _ := cmd.Flags().GetString("t1")
timespan, _ := cmd.Flags().GetString("timespan")
vlan, _ := cmd.Flags().GetString("vlan")
apTag,_ := cmd.Flags().GetString("apTag")
band,_ := cmd.Flags().GetString("band")
ssid,_ := cmd.Flags().GetString("ssid")
metadata := monitor.GetAggregatedConnectivityPerNetwork(networkId,
	t0, t1, timespan, band, ssid, vlan, apTag)
shell.Display(metadata, "connectionstats", cmd.Flags())
},
}




var connectivityevents = &cobra.Command{
Use:   "connectivityevents",
Short: "List the wireless connectivity events for a client within a network in the timespan.",
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

metadata := monitor.GetWirelessConnectivityEvents(networkId,
	clientId, perPage, startingAfter, endingBefore, t0, t1,
	timespan,types, includedSeverities, band, ssidNumber, deviceSerial)
shell.Display(metadata, "connectivityevents", cmd.Flags())
},
}




var latencyhistory = &cobra.Command{
Use:   "latencyhistory",
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
shell.Display(metadata, "latencyhistory", cmd.Flags())
},
}




var latencystat = &cobra.Command{
Use:   "latencystat",
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
metadata := monitor.GetClientLatencyHistory(networkId,clientId,t0, t1,timespan,resolution)
shell.Display(metadata, "latencystat", cmd.Flags())
},
}




var latencystats = &cobra.Command{
Use:   "latencystats",
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
shell.Display(metadata, "latencystats", cmd.Flags())
},
}




var deviceconnectionstats = &cobra.Command{
Use:   "deviceconnectionstats",
Short: "",
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

metadata := monitor.GetAPAggregatedConnectivity(serial, t0,t1,timespan,band,ssid,vlan,apTag)
shell.Display(metadata, "deviceconnectionstats", cmd.Flags())
},
}


var networkconnectionstats = &cobra.Command{
Use:   "networkconnectionstats",
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
metadata := monitor.GetAggregatedConnectivityPerNetwork(networkId,
	t0, t1,timespan,band, ssid, vlan, apTag)
shell.Display(metadata, "networkconnectionstats", cmd.Flags())
},
}

var dataratehistory = &cobra.Command{
Use:   "dataratehistory",
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
	timespan, resolution, autoResolution, clientId, deviceSerial, apTag, band,ssid)
shell.Display(metadata, "dataratehistory", cmd.Flags())
},
}



var connectionstatsnode = &cobra.Command{
Use:   "connectionstatsnode",
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
metadata := monitor.GetAggregatedConnectivityClients(networkId,
	clientId, t0, t1, timespan, band, ssid,vlan, apTag)
shell.Display(metadata, "connectionstatsnode", cmd.Flags())
},
}

var latencystatsnode = &cobra.Command{
Use:   "latencystatsnode",
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
shell.Display(metadata, "latencystatsnode", cmd.Flags())
},
}

var failedconnections = &cobra.Command{
Use:   "failedconnections",
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
metadata := monitor.GetFailedClientConnections(networkId, t0, t1, timespan, band, ssid, vlan, apTag)
shell.Display(metadata, "failedconnections", cmd.Flags())
},
}

var latencyhistoryaverage = &cobra.Command{
Use:   "latencyhistory",
Short: "Return average wireless latency over time for a network, device, or network client.",
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
shell.Display(metadata, "latencyhistory", cmd.Flags())
},
}

var devicelatencystats = &cobra.Command{
Use:   "devicelatencystats",
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
metadata := monitor.GetAggregatedLatencyPerAP(serial, t0,
	t1,timespan,band,ssid,vlan,apTag,fields)
shell.Display(metadata, "devicelatencystats", cmd.Flags())
},
}

var networklatencystats = &cobra.Command{
Use:   "networklatencystats",
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
metadata := monitor.GetAggregatedLatencyPerNetwork(networkId,
	t0,t1,timespan,band,ssid,vlan,apTag, fields)
shell.Display(metadata, "networklatencystats", cmd.Flags())
},
}

var meshstatuses = &cobra.Command{
Use:   "meshstatuses",
Short: "List wireless mesh statuses for repeaters.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

perPage, _ := cmd.Flags().GetString("perPage")
startingAfter, _ := cmd.Flags().GetString("startingAfter")
endingBefore, _ := cmd.Flags().GetString("endingBefore")

metadata := monitor.GetMeshStatuses(networkId, perPage, startingAfter, endingBefore)
shell.Display(metadata, "meshstatuses", cmd.Flags())
},
}


var signalqualityhistory = &cobra.Command{
Use:   "signalqualityhistory",
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

metadata := monitor.GetSignalQuality(networkId,t0,t1,
	timespan,resolution,autoResolution,clientId,
	deviceSerial,apTag,band, ssid)
shell.Display(metadata, "signalqualityhistory", cmd.Flags())
},
}



var status = &cobra.Command{
Use:   "status",
Short: "Return the SSID statuses of an access point.",
Run: func(cmd *cobra.Command, args []string) {
_, _, serial := shell.ResolveFlags(cmd.Flags())
if serial == "" {
	serial = args[0]
}

metadata := monitor.GetSSIDStatuses(serial)
shell.Display(metadata, "status", cmd.Flags())
},
}


var usagehistory = &cobra.Command{
Use:   "usagehistory",
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
metadata := monitor.GetAPUsageHistory(networkId,t0,
	t1,timespan,resolution,autoResolution,clientId,
	deviceSerial,apTag,band,ssid)
shell.Display(metadata, "usagehistory", cmd.Flags())
},
}