package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetWarmSpare(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/warmSpare", api.BaseUrl(), networkId)
	var datamodel = FirewallRules{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
