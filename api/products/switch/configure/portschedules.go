package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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
func GetSwitchPortSchedules(networkId string) (SwitchPortSchedules, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/portSchedules",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = SwitchPortSchedules{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
