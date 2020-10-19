package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type UplinkSettings struct {
	BandwidthLimits struct {
		LimitUp   int `json:"limitUp"`
		LimitDown int `json:"limitDown"`
	} `json:"bandwidthLimits"`
}

// Returns The Uplink Settings For Your MG Network
func GetUplinkSettings(networkId string) (UplinkSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/cellularGateway/uplink",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = UplinkSettings{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}