package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type PacketCounters []struct {
	PortID  string `json:"portId"`
	Packets []struct {
		Desc       string `json:"desc"`
		Total      int    `json:"total"`
		Sent       int    `json:"sent"`
		Recv       int    `json:"recv"`
		RatePerSec struct {
			Total int `json:"total"`
			Sent  int `json:"sent"`
			Recv  int `json:"recv"`
		} `json:"ratePerSec"`
	} `json:"packets"`
}

// Return the packet counters for all the ports of a switch
func GetPacketCounters(serial, t0, timespan string) (PacketCounters, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/ports/statuses/packets",
		api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("timespan",timespan)

	session.Request.URL.RawQuery = parameters.Encode()
	var results = PacketCounters{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


type PortsStatus []struct {
	PortID    string   `json:"portId"`
	Enabled   bool     `json:"enabled"`
	Status    string   `json:"status"`
	Errors    []string `json:"errors"`
	Warnings  []string `json:"warnings"`
	Speed     string   `json:"speed"`
	Duplex    string   `json:"duplex"`
	UsageInKb struct {
		Total int `json:"total"`
		Sent  int `json:"sent"`
		Recv  int `json:"recv"`
	} `json:"usageInKb"`
	Cdp struct {
		SystemName          string `json:"systemName"`
		Platform            string `json:"platform"`
		DeviceID            string `json:"deviceId"`
		PortID              string `json:"portId"`
		NativeVlan          int    `json:"nativeVlan"`
		Address             string `json:"address"`
		ManagementAddress   string `json:"managementAddress"`
		Version             string `json:"version"`
		VtpManagementDomain string `json:"vtpManagementDomain"`
		Capabilities        string `json:"capabilities"`
	} `json:"cdp"`
	Lldp struct {
		SystemName         string `json:"systemName"`
		SystemDescription  string `json:"systemDescription"`
		PortID             string `json:"portId"`
		PortDescription    string `json:"portDescription"`
		ChassisID          string `json:"chassisId"`
		ManagementVlan     int    `json:"managementVlan"`
		PortVlan           int    `json:"portVlan"`
		ManagementAddress  string `json:"managementAddress"`
		SystemCapabilities string `json:"systemCapabilities"`
	} `json:"lldp"`
	ClientCount    int     `json:"clientCount"`
	PowerUsageInWh float64 `json:"powerUsageInWh"`
	TrafficInKbps  struct {
		Total float64 `json:"total"`
		Sent  float64 `json:"sent"`
		Recv  int     `json:"recv"`
	} `json:"trafficInKbps"`
}

// Return The Status For All The Ports Of A Switch
func GetPortsStatus(serial, t0, timespan string) (PortsStatus, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/ports/statuses",
		api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("timespan",timespan)

	session.Request.URL.RawQuery = parameters.Encode()
	var results = PortsStatus{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}