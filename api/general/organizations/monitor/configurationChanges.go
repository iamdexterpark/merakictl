package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
	"time"
)

type ConfigurationChanges []struct {
	ConfigurationChange
}

type ConfigurationChange struct {
	Ts         time.Time `json:"ts"`
	AdminName  string    `json:"adminName"`
	AdminEmail string    `json:"adminEmail"`
	AdminID    string    `json:"adminId"`
	Page       string    `json:"page"`
	Label      string    `json:"label"`
	OldValue   string    `json:"oldValue"`
	NewValue   string    `json:"newValue"`
}

// View the Change Log for your organization
func GetConfigurationChanges(organizationId, t0, t1, timespan, perPage, startingAfter, endingBefore,
	adminId, networkId  string ) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/configurationChanges", api.BaseUrl(), organizationId)
	var datamodel = ConfigurationChanges{}


	// Parameters for Request URL
	var parameters = map[string]string{
		"t0": t0,
		"t1": t1,
		"timespan": timespan,
		"perPage": perPage,
		"startingAfter": startingAfter,
		"endingBefore": endingBefore,
		"networkId": networkId,
		"adminId": adminId}


	sessions, err := api.Sessions(baseurl, "GET", nil,  parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
