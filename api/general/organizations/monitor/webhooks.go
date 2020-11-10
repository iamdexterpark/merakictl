package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
	"time"
)

type WebHookLogs []struct {
	Log
}
type Log struct {
	OrganizationID   string    `json:"organizationId"`
	NetworkID        string    `json:"networkId"`
	AlertType        string    `json:"alertType"`
	URL              string    `json:"url"`
	SentAt           time.Time `json:"sentAt"`
	LoggedAt         time.Time `json:"loggedAt"`
	ResponseCode     int       `json:"responseCode"`
	ResponseDuration int       `json:"responseDuration"`
}

// Return the log of webhook POSTs sent
func GetWebHookLogs(organizationId, t0, t1, timespan, perPage, startingAfter, endingBefore,
	url  string ) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/webhooks/logs", api.BaseUrl(), organizationId)
	var datamodel = WebHookLogs{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0": t0,
		"t1": t1,
		"timespan": timespan,
		"perPage": perPage,
		"startingAfter": startingAfter,
		"endingBefore": endingBefore,
		"url": url}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}