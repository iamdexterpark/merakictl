package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type PacketCounters []struct {
	PortID  string `json:"portId"`
	Packets []struct {
		Desc       string `json:"desc"`
		Total      int    `json:"total"`
		Sent       int    `json:"sent"`
		Recv       int    `json:"recv"`
		RatePerSec struct {
			Total int `json:"total"`
			Sent  int `json:"sent"`
			Recv  int `json:"recv"`
		} `json:"ratePerSec"`
	} `json:"packets"`
}

// Return the packet counters for all the ports of a switch
func GetPacketCounters(serial, t0, timespan string) (PacketCounters, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/ports/statuses/packets",
		api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("timespan",timespan)

	session.Request.URL.RawQuery = parameters.Encode()
	var results = PacketCounters{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
// Return The Status For All The Ports Of A Switch
