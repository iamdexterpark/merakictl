package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type OSPFRouting struct {
	Enabled             bool `json:"enabled"`
	HelloTimerInSeconds int  `json:"helloTimerInSeconds"`
	DeadTimerInSeconds  int  `json:"deadTimerInSeconds"`
	Areas               []struct {
		AreaID   string `json:"areaId"`
		AreaName string `json:"areaName"`
		AreaType string `json:"areaType"`
	} `json:"areas"`
	Md5AuthenticationEnabled bool `json:"md5AuthenticationEnabled"`
	Md5AuthenticationKey     struct {
		ID         string `json:"id"`
		Passphrase string `json:"passphrase"`
	} `json:"md5AuthenticationKey"`
}

// Return Layer 3 OSPF Routing Configuration
func GetOSPFRouting(networkId string) (OSPFRouting, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/routing/ospf",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = OSPFRouting{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
