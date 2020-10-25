package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type STPSettings struct {
	RstpEnabled       bool `json:"rstpEnabled"`
	StpBridgePriority []struct {
		Switches    []string `json:"switches,omitempty"`
		StpPriority int      `json:"stpPriority"`
		Stacks      []string `json:"stacks,omitempty"`
	} `json:"stpBridgePriority"`
}

// Returns STP Settings
func GetSTPSettings(networkId string) (STPSettings , interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stp",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = STPSettings {}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
