package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type DeviceClients []struct {
	Usage struct {
		Sent int `json:"sent"`
		Recv int `json:"recv"`
	} `json:"usage"`
	ID           string      `json:"id"`
	Description  string      `json:"description"`
	Mac          string      `json:"mac"`
	IP           string      `json:"ip"`
	User         string      `json:"user"`
	Vlan         int         `json:"vlan"`
	Switchport   interface{} `json:"switchport"`
	MdnsName     string      `json:"mdnsName"`
	DhcpHostname string      `json:"dhcpHostname"`
}

// GetSingleDevice - Return A Single Device
func GetDeviceClients(serial string) (DeviceClients, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/clients", api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var device = DeviceClients{}
	user_agent.UnMarshalJSON(session.Body, &device)
	traceback := user_agent.TraceBack(session)
	return device, traceback
}
