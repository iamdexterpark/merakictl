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

// GetDeviceClients - Return A Devices Clients
func GetDeviceClients(serial string) (DeviceClients, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/clients", api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var device = DeviceClients{}
	user_agent.UnMarshalJSON(session.Body, &device)
	traceback := user_agent.TraceBack(session)
	return device, traceback
}


type lldpCdp struct {
	SourceMac string `json:"sourceMac"`
	Ports     struct {
		Num8 struct {
			Cdp struct {
				DeviceID   string `json:"deviceId"`
				PortID     string `json:"portId"`
				Address    string `json:"address"`
				SourcePort string `json:"sourcePort"`
			} `json:"cdp"`
		} `json:"8"`
		Num12 struct {
			Cdp struct {
				DeviceID   string `json:"deviceId"`
				PortID     string `json:"portId"`
				Address    string `json:"address"`
				SourcePort string `json:"sourcePort"`
			} `json:"cdp"`
			Lldp struct {
				SystemName        string `json:"systemName"`
				PortID            string `json:"portId"`
				ManagementAddress string `json:"managementAddress"`
				SourcePort        string `json:"sourcePort"`
			} `json:"lldp"`
		} `json:"12"`
	} `json:"ports"`
}

// Get - Return
func GetlldpCdp(serial string) (lldpCdp, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/lldpCdp", api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var lldpCdp = lldpCdp{}
	user_agent.UnMarshalJSON(session.Body, &lldpCdp)
	traceback := user_agent.TraceBack(session)
	return lldpCdp, traceback
}