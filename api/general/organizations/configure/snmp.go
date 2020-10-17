package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type SNMP struct {
	V2CEnabled bool     `json:"v2cEnabled"`
	V3Enabled  bool     `json:"v3Enabled"`
	V3AuthMode string   `json:"v3AuthMode"`
	V3PrivMode string   `json:"v3PrivMode"`
	PeerIps    []string `json:"peerIps"`
	Hostname   string   `json:"hostname"`
	Port       int      `json:"port"`
}

// Return the SNMP settings for an organization
func GetSNMP(organizationId string) (SNMP, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/snmp", api.BaseUrl(),
		organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = SNMP{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}