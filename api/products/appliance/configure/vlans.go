package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type VLANStatus struct {
	VlansEnabled bool `json:"vlansEnabled"`
}

// Returns the enabled status of VLANs for the network
func GetVLANStatus(networkId string ) (VLANStatus, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/vlans/settings", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = VLANStatus{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type VLANs []struct {
	VLAN
}
type VLAN struct {
	ID                 string `json:"id"`
	NetworkID          string `json:"networkId"`
	Name               string `json:"name"`
	ApplianceIP        string `json:"applianceIp"`
	Subnet             string `json:"subnet"`
	GroupPolicyID      string `json:"groupPolicyId"`
	FixedIPAssignments struct {
		Two23344556677 struct {
			IP   string `json:"ip"`
			Name string `json:"name"`
		} `json:"22:33:44:55:66:77"`
	} `json:"fixedIpAssignments"`
	ReservedIPRanges []struct {
		Start   string `json:"start"`
		End     string `json:"end"`
		Comment string `json:"comment"`
	} `json:"reservedIpRanges"`
	DNSNameservers         string      `json:"dnsNameservers"`
	DhcpHandling           string      `json:"dhcpHandling"`
	DhcpLeaseTime          string      `json:"dhcpLeaseTime"`
	DhcpBootOptionsEnabled bool        `json:"dhcpBootOptionsEnabled"`
	DhcpBootNextServer     interface{} `json:"dhcpBootNextServer"`
	DhcpBootFilename       interface{} `json:"dhcpBootFilename"`
	DhcpOptions            []struct {
		Code  int    `json:"code"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"dhcpOptions"`
}

// List the VLANs for an MX network
func GetVLANs(networkId string ) (VLANStatus, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/vlans/", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = VLANStatus{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// List a VLAN for an MX network
func GetVLAN(networkId, vlanId string ) (VLANStatus, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/vlans/%s", api.BaseUrl(), networkId, vlanId)
	var payload io.ReadSeeker
	var results = VLANStatus{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}