package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
	"time"
)

type DeviceStatuses []struct {
	DeviceStatus
}
type DeviceStatus struct {
	Name           string    `json:"name"`
	Serial         string    `json:"serial"`
	Mac            string    `json:"mac"`
	PublicIP       string    `json:"publicIp"`
	NetworkID      string    `json:"networkId"`
	Status         string    `json:"status"`
	LastReportedAt time.Time `json:"lastReportedAt"`
	LanIP          string    `json:"lanIp"`
	Gateway        string    `json:"gateway"`
	IPType         string    `json:"ipType"`
	PrimaryDNS     string    `json:"primaryDns"`
	SecondaryDNS   string    `json:"secondaryDns"`
}

type LossLatencyList []struct {
}
type LossLatency []struct {
	NetworkID  string `json:"networkId"`
	Serial     string `json:"serial"`
	Uplink     string `json:"uplink"`
	IP         string `json:"ip"`
	TimeSeries []struct {
		Ts          time.Time `json:"ts"`
		LossPercent float64   `json:"lossPercent"`
		LatencyMs   float64   `json:"latencyMs"`
	} `json:"timeSeries"`
}

// List the status of every Meraki device in the organization
func GetDeviceStatus(organizationId, perPage, startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/devices/statuses", api.BaseUrl(), organizationId)
	var datamodel = DeviceStatuses{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"perPage":       perPage,
		"startingAfter": startingAfter,
		"endingBefore":  endingBefore}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return the uplink loss and latency for every MX in the organization from at latest 2 minutes ago
func GetLossLatency(organizationId, t0, t1, timespan, uplink, ip string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/devices/uplinksLossAndLatency", api.BaseUrl(), organizationId)
	var datamodel = APIRequests{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":       t0,
		"t1":       t1,
		"timespan": timespan,
		"uplink":   uplink,
		"ip":       ip}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
