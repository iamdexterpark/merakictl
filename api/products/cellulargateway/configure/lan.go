package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type LANSettings struct {
	DeviceName         string `json:"deviceName"`
	DeviceLanIP        string `json:"deviceLanIp"`
	DeviceSubnet       string `json:"deviceSubnet"`
	FixedIPAssignments []struct {
		Mac  string `json:"mac"`
		Name string `json:"name"`
		IP   string `json:"ip"`
	} `json:"fixedIpAssignments"`
	ReservedIPRanges []struct {
		Start   string `json:"start"`
		End     string `json:"end"`
		Comment string `json:"comment"`
	} `json:"reservedIpRanges"`
}

// Show the LAN Settings of a MG
func GetLANSettings(serial string) (LANSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/cellularGateway/lan",
		api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = LANSettings{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}