package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type Layer3DHCPInterface struct {
	DhcpMode             string   `json:"dhcpMode"`
	DhcpLeaseTime        string   `json:"dhcpLeaseTime"`
	DNSNameserversOption string   `json:"dnsNameserversOption"`
	DNSCustomNameservers []string `json:"dnsCustomNameservers"`
	BootOptionsEnabled   bool     `json:"bootOptionsEnabled"`
	BootNextServer       string   `json:"bootNextServer"`
	BootFileName         string   `json:"bootFileName"`
	DhcpOptions          []struct {
		Code  string `json:"code"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"dhcpOptions"`
	ReservedIPRanges []struct {
		Start   string `json:"start"`
		End     string `json:"end"`
		Comment string `json:"comment"`
	} `json:"reservedIpRanges"`
	FixedIPAssignments []struct {
		Mac  string `json:"mac"`
		Name string `json:"name"`
		IP   string `json:"ip"`
	} `json:"fixedIpAssignments"`
}

// Return A Layer 3 Interface DHCP Configuration For A Switch
func GetLayer3DHCPInterface(serial, interfaceId string) (Layer3DHCPInterface, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/routing/interfaces/%s/dhcp",
		api.BaseUrl(), serial, interfaceId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = Layer3DHCPInterface{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}