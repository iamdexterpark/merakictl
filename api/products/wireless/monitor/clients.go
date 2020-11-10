package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type AggregatedConnectivityInfo struct {
	Mac             string `json:"mac"`
	ConnectionStats struct {
		Assoc   int `json:"assoc"`
		Auth    int `json:"auth"`
		Dhcp    int `json:"dhcp"`
		DNS     int `json:"dns"`
		Success int `json:"success"`
	} `json:"connectionStats"`
}

type AggregatedConnectivityClients struct {
	Assoc   int `json:"assoc"`
	Auth    int `json:"auth"`
	Dhcp    int `json:"dhcp"`
	DNS     int `json:"dns"`
	Success int `json:"success"`
}

type WirelessConnectivityEvents []struct {
	OccurredAt   int     `json:"occurredAt"`
	DeviceSerial string  `json:"deviceSerial"`
	Band         int     `json:"band"`
	SsidNumber   int     `json:"ssidNumber"`
	Type         string  `json:"type"`
	Subtype      string  `json:"subtype"`
	Severity     string  `json:"severity"`
	DurationMs   float64 `json:"durationMs"`
	Channel      int     `json:"channel"`
	Rssi         int     `json:"rssi"`
	EventData    struct {
		ClientIP string `json:"clientIp"`
	} `json:"eventData"`
}

type ClientLatencyTraffic struct {
	Zero5   int `json:"0.5"`
	One1    int `json:"1.1"`
	Two1    int `json:"2.1"`
	Four1   int `json:"4.1"`
	Eight1  int `json:"8.1"`
	One61   int `json:"16.1"`
	Three21 int `json:"32.1"`
	Six41   int `json:"64.1"`
	One281  int `json:"128.1"`
	Two561  int `json:"256.1"`
	Five121 int `json:"512.1"`
	One0241 int `json:"1024.1"`
	Two0481 int `json:"2048.1"`
}
type ClientLatencyHistory []struct {
	T0                    int `json:"t0"`
	T1                    int `json:"t1"`
	LatencyBinsByCategory struct {
		BackgroundTraffic struct {
			ClientLatencyTraffic
		} `json:"backgroundTraffic"`
		BestEffortTraffic struct {
			ClientLatencyTraffic
		} `json:"bestEffortTraffic"`
		VideoTraffic struct {
			ClientLatencyTraffic
		} `json:"videoTraffic"`
		VoiceTraffic struct {
			ClientLatencyTraffic
		} `json:"voiceTraffic"`
	} `json:"latencyBinsByCategory"`
}

type AggregatedLatencies struct {
	AggregatedLatency
}

type AggregatedLatency struct {
	Mac          string `json:"mac"`
	LatencyStats struct {
		BackgroundTraffic struct {
			RawDistribution struct {
				Num0    int `json:"0"`
				Num1    int `json:"1"`
				Num2    int `json:"2"`
				Num4    int `json:"4"`
				Num8    int `json:"8"`
				Num16   int `json:"16"`
				Num32   int `json:"32"`
				Num64   int `json:"64"`
				Num128  int `json:"128"`
				Num256  int `json:"256"`
				Num512  int `json:"512"`
				Num1024 int `json:"1024"`
				Num2048 int `json:"2048"`
			} `json:"rawDistribution"`
			Avg float64 `json:"avg"`
		} `json:"backgroundTraffic"`
		BestEffortTraffic string `json:"bestEffortTraffic"`
		VideoTraffic      string `json:"videoTraffic"`
		VoiceTraffic      string `json:"voiceTraffic"`
	} `json:"latencyStats"`
}

// Aggregated Connectivity Info For A Given Client On This Network
func GetAggregatedConnectivityInfo(devices, serial, t0, t1, timespan,
	band, ssid, vlan, apTag string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/wireless/clients/%s/connectionStats",
		api.BaseUrl(), devices, serial)
	var datamodel = AggregatedConnectivityInfo{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0": t0,
		"t1": t1,
		"timespan": timespan,
		"band": band,
		"ssid": ssid,
		"vlan": vlan,
		"apTag": apTag}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Aggregated Connectivity Info For This Network Grouped By Clients
func GetAggregatedConnectivityClients(networkId, clientId, t0, t1, timespan,
	band, ssid, vlan, apTag string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/clients/%s/connectionStats",
		api.BaseUrl(), networkId, clientId)
	var datamodel = AggregatedConnectivityClients{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0": t0,
		"t1": t1,
		"timespan": timespan,
		"band": band,
		"ssid": ssid,
		"vlan": vlan,
		"apTag": apTag}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List the wireless connectivity events for a client within a network in the timespan
func GetWirelessConnectivityEvents(networkId, clientId, perPage, startingAfter,
	endingBefore, t0, t1, timespan, types, includedSeverities, band, ssidNumber,
	deviceSerial string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/clients/%s/connectivityEvents",
		api.BaseUrl(), networkId, clientId)
	var datamodel = WirelessConnectivityEvents{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"perPage": perPage,
		"startingAfter": startingAfter,
		"endingBefore": endingBefore,
		"t0": t0,
		"t1": t1,
		"timespan": timespan,
		"types": types,
		"includedSeverities": includedSeverities,
		"band": band,
		"ssidNumber": ssidNumber,
		"deviceSerial": deviceSerial}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return the latency history for a client
func GetClientLatencyHistory(networkId, clientId, t0, t1, timespan,
	resolution string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/clients/%s/latencyHistory",
		api.BaseUrl(), networkId, clientId)
	var datamodel = ClientLatencyHistory{}


	// Parameters for Request URL
	var parameters = map[string]string{
		"t0": t0,
		"t1": t1,
		"timespan": timespan,
		"resolution": resolution}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Aggregated Latency Info For This Network Grouped By Clients
func GetAggregatedLatencies(networkId, t0, t1, timespan,
	band, ssid, vlan, apTag, fields string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/clients/latencyStats",
		api.BaseUrl(), networkId)
	var datamodel = AggregatedLatencies{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0": t0,
		"t1": t1,
		"timespan": timespan,
		"band": band,
		"ssid": ssid,
		"vlan": vlan,
		"apTag": apTag,
		"fields": fields}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Aggregated Latency Info For A Given Client On This Network
func GetAggregatedLatency(networkId, clientId, t0, t1, timespan,
	band, ssid, vlan, apTag, fields string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/clients/%s/latencyStats",
		api.BaseUrl(), networkId, clientId)
	var datamodel = AggregatedLatency{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0": t0,
		"t1": t1,
		"timespan": timespan,
		"band": band,
		"ssid": ssid,
		"vlan": vlan,
		"apTag": apTag,
		"fields": fields}


	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

