package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type PortForwardingRules struct {
	Rules []struct {
		LanIP      string   `json:"lanIp"`
		Name       string   `json:"name"`
		Access     string   `json:"access"`
		PublicPort string   `json:"publicPort"`
		LocalPort  string   `json:"localPort"`
		Uplink     string   `json:"uplink"`
		Protocol   string   `json:"protocol"`
		AllowedIps []string `json:"allowedIps,omitempty"`
	} `json:"rules"`
}

// Returns the port forwarding rules for a single MG
func GetPortForwardingRules(serial string) (PortForwardingRules, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/cellularGateway/portForwardingRules",
		api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = PortForwardingRules{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}