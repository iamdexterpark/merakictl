package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type ManagementInterface struct {
	DdnsHostnames struct {
		ActiveDdnsHostname string `json:"activeDdnsHostname"`
		DdnsHostnameWan1   string `json:"ddnsHostnameWan1"`
		DdnsHostnameWan2   string `json:"ddnsHostnameWan2"`
	} `json:"ddnsHostnames"`
	Wan1 struct {
		WanEnabled       string   `json:"wanEnabled"`
		UsingStaticIP    bool     `json:"usingStaticIp"`
		StaticIP         string   `json:"staticIp"`
		StaticSubnetMask string   `json:"staticSubnetMask"`
		StaticGatewayIP  string   `json:"staticGatewayIp"`
		StaticDNS        []string `json:"staticDns"`
		Vlan             int      `json:"vlan"`
	} `json:"wan1"`
	Wan2 struct {
		WanEnabled    string `json:"wanEnabled"`
		UsingStaticIP bool   `json:"usingStaticIp"`
		Vlan          int    `json:"vlan"`
	} `json:"wan2"`
}

// GetManagementInterface - Return The Management Interface Settings For A Device
func GetManagementInterface(serial string) (ManagementInterface, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/managementInterface", api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var clients = ManagementInterface{}
	user_agent.UnMarshalJSON(session.Body, &clients)
	traceback := user_agent.TraceBack(session)
	return clients, traceback
}

type Device struct {
	Name           string   `json:"name"`
	Lat            float64  `json:"lat"`
	Lng            float64  `json:"lng"`
	Serial         string   `json:"serial"`
	Mac            string   `json:"mac"`
	Model          string   `json:"model"`
	Address        string   `json:"address"`
	Notes          string   `json:"notes"`
	LanIP          string   `json:"lanIp"`
	Tags           []string `json:"tags"`
	NetworkID      string   `json:"networkId"`
	BeaconIDParams struct {
		UUID  string `json:"uuid"`
		Major int    `json:"major"`
		Minor int    `json:"minor"`
	} `json:"beaconIdParams"`
	Firmware    string `json:"firmware"`
	FloorPlanID string `json:"floorPlanId"`
}

// GetSingleDevice - Return A Single Device
func GetSingleDevice(serial string) (Device, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s", api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var device = Device{}
	user_agent.UnMarshalJSON(session.Body, &device)
	traceback := user_agent.TraceBack(session)
	return device, traceback
}