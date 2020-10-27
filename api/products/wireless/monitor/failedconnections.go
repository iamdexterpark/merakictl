package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
	"time"
)

type FailedClientConnections []struct {
	SsidNumber  int       `json:"ssidNumber"`
	Vlan        int       `json:"vlan"`
	ClientMac   string    `json:"clientMac"`
	Serial      string    `json:"serial,omitempty"`
	FailureStep string    `json:"failureStep"`
	Type        string    `json:"type"`
	Ts          time.Time `json:"ts"`
	NodeID      string    `json:"nodeId,omitempty"`
}

// List of all failed client connection events on this network in a given time range
func GetFailedClientConnections(networkId, t0, t1, timespan,
	band, ssid, vlan, apTag string) (FailedClientConnections, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/failedConnections",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("t1", t1)
	parameters.Add("timespan", timespan)
	parameters.Add("band", band)
	parameters.Add("ssid", ssid)
	parameters.Add("vlan", vlan)
	parameters.Add("apTag", apTag)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = FailedClientConnections{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}