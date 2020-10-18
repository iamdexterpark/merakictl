package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type WarmSpare struct {
	Enabled       bool   `json:"enabled"`
	PrimarySerial string `json:"primarySerial"`
	SpareSerial   string `json:"spareSerial"`
	UplinkMode    string `json:"uplinkMode"`
	Wan1          struct {
		IP     string `json:"ip"`
		Subnet string `json:"subnet"`
	} `json:"wan1"`
	Wan2 struct {
		IP     string `json:"ip"`
		Subnet string `json:"subnet"`
	} `json:"wan2"`
}

// Return MX warm spare settings
func GetWarmSpare(networkId string ) (FirewallRules, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/warmSpare", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = FirewallRules{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}