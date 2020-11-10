package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type AggregatedConnectivityPerAP struct {
	Serial          string `json:"serial"`
	ConnectionStats struct {
		Assoc   int `json:"assoc"`
		Auth    int `json:"auth"`
		Dhcp    int `json:"dhcp"`
		DNS     int `json:"dns"`
		Success int `json:"success"`
	} `json:"connectionStats"`
}

type AggregatedConnectivityPerNetwork struct {
	Assoc   int `json:"assoc"`
	Auth    int `json:"auth"`
	Dhcp    int `json:"dhcp"`
	DNS     int `json:"dns"`
	Success int `json:"success"`
}

// Aggregated Connectivity Info For A Given AP On This Network
func GetAPAggregatedConnectivity(serial, t0, t1, timespan,
	band, ssid, vlan, apTag string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/wireless/connectionStats",
		api.BaseUrl(), serial)
	var datamodel = AggregatedConnectivityPerAP{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":       t0,
		"t1":       t1,
		"timespan": timespan,
		"band":     band,
		"ssid":     ssid,
		"vlan":     vlan,
		"apTag":    apTag}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Aggregated Connectivity Info For This Network
func GetAggregatedConnectivityPerNetwork(networkId, t0, t1, timespan,
	band, ssid, vlan, apTag string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/connectionStats",
		api.BaseUrl(), networkId)
	var datamodel = AggregatedConnectivityPerNetwork{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":       t0,
		"t1":       t1,
		"timespan": timespan,
		"band":     band,
		"ssid":     ssid,
		"vlan":     vlan,
		"apTag":    apTag}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
