package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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
func GetNetworkAlertConfig(networkId string) (AlertConfig, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/alerts/settings", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var alertconfig = AlertConfig{}
	user_agent.UnMarshalJSON(session.Body, &alertconfig)
	traceback := user_agent.TraceBack(session)
	return alertconfig, traceback
}