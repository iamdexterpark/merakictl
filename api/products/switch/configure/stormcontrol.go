package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type StormControl struct {
	BroadcastThreshold      int `json:"broadcastThreshold"`
	MulticastThreshold      int `json:"multicastThreshold"`
	UnknownUnicastThreshold int `json:"unknownUnicastThreshold"`
}

// Return The Storm Control Configuration For A Switch Network
func GetStormControl(networkId string) (StormControl, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stormControl",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = StormControl{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
