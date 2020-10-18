package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type DHCP []struct {
	Subnet    string `json:"subnet"`
	VlanID    int    `json:"vlanId"`
	UsedCount int    `json:"usedCount"`
	FreeCount int    `json:"freeCount"`
}

// Return the DHCP subnet information for an appliance
func GetDHCP(serial string ) (DHCP, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/appliance/dhcp/subnets", api.BaseUrl(), serial)
	var payload io.ReadSeeker
	var results = DHCP{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}