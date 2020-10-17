package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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

// List the status of every Meraki device in the organization
func GetDeviceStatus(organizationId, perPage, startingAfter, endingBefore string ) (DeviceStatuses, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/devices/statuses", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker
	var results = DeviceStatuses{}

	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("perPage", perPage)
	parameters.Add("startingAfter", startingAfter)
	parameters.Add("endingBefore", endingBefore)

	session.Request.URL.RawQuery = parameters.Encode()

	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
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

// Return the uplink loss and latency for every MX in the organization from at latest 2 minutes ago
func GetLossLatency(organizationId, t0, t1, timespan, uplink, ip  string ) (APIRequests, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/devices/uplinksLossAndLatency", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker
	var results = APIRequests{}

	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("t1", t1)
	parameters.Add("timespan",timespan)
	parameters.Add("uplink", uplink)
	parameters.Add("ip", ip)

	session.Request.URL.RawQuery = parameters.Encode()

	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}