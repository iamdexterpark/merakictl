package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type Device []struct {
	Name      string   `json:"name"`
	Lat       float64  `json:"lat"`
	Lng       float64  `json:"lng"`
	Address   string   `json:"address"`
	Notes     string   `json:"notes"`
	Tags      []string `json:"tags"`
	NetworkID string   `json:"networkId"`
	Serial    string   `json:"serial"`
	Model     string   `json:"model"`
	Mac       string   `json:"mac"`
	LanIP     string   `json:"lanIp"`
	Firmware  string   `json:"firmware"`
}

// List the devices in an organization
func GetDeviceList(organizationId string) (Device, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/devices", api.BaseUrl(),
		organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = Device{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}