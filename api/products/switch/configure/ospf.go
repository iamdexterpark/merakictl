package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetOSPFRouting(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/routing/ospf",
		api.BaseUrl(), networkId)
	var datamodel = OSPFRouting{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
