package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type LANConfiguration struct {
	Subnet      string `json:"subnet"`
	ApplianceIP string `json:"applianceIp"`
}

// Return single LAN configuration
func GetLANConfiguration(networkId string ) (LANConfiguration, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/singleLan", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = LANConfiguration{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}