package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type MaintenanceWindow struct {
	UpgradeWindow struct {
		HourOfDay string `json:"hourOfDay"`
		DayOfWeek string `json:"dayOfWeek"`
	} `json:"upgradeWindow"`
}

// Get Current Maintenance Window For A Network
func GetMaintenanceWindow(networkId string) (MaintenanceWindow, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/firmwareUpgrades", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var maint = MaintenanceWindow{}
	user_agent.UnMarshalJSON(session.Body, &maint)
	traceback := user_agent.TraceBack(session)
	return maint, traceback
}