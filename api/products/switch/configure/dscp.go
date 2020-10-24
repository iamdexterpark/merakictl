package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type DSCPCOSMapping struct {
	Mappings []struct {
		Dscp  int    `json:"dscp"`
		Cos   int    `json:"cos"`
		Title string `json:"title"`
	} `json:"mappings"`
}

// Return The DSCP To CoS Mappings
func GetDSCPCOSMapping(networkId string) (DSCPCOSMapping, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/dscpToCosMappings",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = DSCPCOSMapping{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
