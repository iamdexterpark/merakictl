package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type MaintenanceWindow struct {
	UpgradeWindow struct {
		HourOfDay string `json:"hourOfDay"`
		DayOfWeek string `json:"dayOfWeek"`
	} `json:"upgradeWindow"`
}

// Get Current Maintenance Window For A Network
func GetMaintenanceWindow(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/firmwareUpgrades", api.BaseUrl(), networkId)
	var datamodel = MaintenanceWindow{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
