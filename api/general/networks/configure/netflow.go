package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type NetFlow struct {
	ReportingEnabled bool   `json:"reportingEnabled"`
	CollectorIP      string `json:"collectorIp"`
	CollectorPort    int    `json:"collectorPort"`
}

// Return The Net Flow Traffic Reporting Settings For A Network
func GetNetFlow(networkId string) (NetFlow, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/netflow", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = NetFlow{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

