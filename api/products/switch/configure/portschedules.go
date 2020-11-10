package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type SwitchPortSchedules []struct {
	ID           string `json:"id"`
	NetworkID    string `json:"networkId"`
	Name         string `json:"name"`
	PortSchedule struct {
		Monday struct {
			Active bool   `json:"active"`
			From   string `json:"from"`
			To     string `json:"to"`
		} `json:"monday"`
		Tuesday struct {
			Active bool   `json:"active"`
			From   string `json:"from"`
			To     string `json:"to"`
		} `json:"tuesday"`
		Wednesday struct {
			Active bool   `json:"active"`
			From   string `json:"from"`
			To     string `json:"to"`
		} `json:"wednesday"`
		Thursday struct {
			Active bool   `json:"active"`
			From   string `json:"from"`
			To     string `json:"to"`
		} `json:"thursday"`
		Friday struct {
			Active bool   `json:"active"`
			From   string `json:"from"`
			To     string `json:"to"`
		} `json:"friday"`
		Saturday struct {
			Active bool   `json:"active"`
			From   string `json:"from"`
			To     string `json:"to"`
		} `json:"saturday"`
		Sunday struct {
			Active bool   `json:"active"`
			From   string `json:"from"`
			To     string `json:"to"`
		} `json:"sunday"`
	} `json:"portSchedule"`
}

// List Switch Port Schedules
func GetSwitchPortSchedules(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/portSchedules",
		api.BaseUrl(), networkId)
	var datamodel = SwitchPortSchedules{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
