package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type AlertConfig struct {
	DefaultDestinations struct {
		Emails    []string `json:"emails"`
		AllAdmins bool     `json:"allAdmins"`
		Snmp      bool     `json:"snmp"`
	} `json:"defaultDestinations"`
	Alerts []struct {
		Type              string `json:"type"`
		Enabled           bool   `json:"enabled"`
		AlertDestinations struct {
			Emails    []string `json:"emails"`
			AllAdmins bool     `json:"allAdmins"`
			Snmp      bool     `json:"snmp"`
		} `json:"alertDestinations"`
		Filters struct {
			Timeout int `json:"timeout"`
		} `json:"filters"`
	} `json:"alerts"`
}

// GetNetworkAlertConfig - Return The Alert Configuration For This Network
func GetNetworkAlertConfig(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/alerts/settings", api.BaseUrl(), networkId)

	var datamodel = AlertConfig{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}