package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
	"time"
)

// DeviceClients - Return A Devices Clients
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

// lldpCdp - List LLDP and CDP information for a device
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

// UplinkLoss -  Get the uplink loss percentage and latency in milliseconds for a wired network device.
type UplinkLoss []struct {
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	LossPercent int       `json:"lossPercent"`
	LatencyMs   int       `json:"latencyMs"`
}

// GetDeviceClients - Return A Devices Clients
func GetDeviceClients(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/clients", api.BaseUrl(), serial)
	var datamodel = DeviceClients{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}

// GetlldpCdp - List LLDP and CDP information for a device
func GetlldpCdp(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/lldpCdp", api.BaseUrl(), serial)
	var datamodel = lldpCdp{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}

// GetlldpCdp - Get the uplink loss percentage and latency in milliseconds for a wired network device.
func GetUplinkLoss(serial, t0, t1, timespan, resolution, uplink, ip string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/lossAndLatencyHistory", api.BaseUrl(), serial)
	var datamodel = UplinkLoss{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":         t0,
		"t1":         t1,
		"timespan":   timespan,
		"resolution": resolution,
		"uplink":     uplink,
		"ip":         ip,
	}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions

}
