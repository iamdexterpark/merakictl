package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type SubnetPool struct {
	Cidr    string `json:"cidr"`
	Mask    string `json:"mask"`
	Subnets []struct {
		Serial      string `json:"serial"`
		Name        string `json:"name"`
		ApplianceIP string `json:"applianceIp"`
		Subnet      string `json:"subnet"`
	} `json:"subnets"`
}

// Return the subnet pool and mask configured for MGs in the network.
func GetSubnetPool(networkId string) (SubnetPool, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/cellularGateway/subnetPool",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = SubnetPool{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}