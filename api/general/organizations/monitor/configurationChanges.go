package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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
	adminId, networkId  string ) (ConfigurationChanges, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/configurationChanges", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker

	var results = ConfigurationChanges{}
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("t1", t1)
	parameters.Add("timespan",timespan)
	parameters.Add("perPage", perPage)
	parameters.Add("startingAfter", startingAfter)
	parameters.Add("endingBefore", endingBefore)
	parameters.Add("networkId", networkId)
	parameters.Add("adminId", adminId)

	session.Request.URL.RawQuery = parameters.Encode()

	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)

	return results, traceback
}
