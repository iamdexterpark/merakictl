package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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
	url  string ) (WebHookLogs, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/webhooks/logs", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker
	var results = WebHookLogs{}

	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("t1", t1)
	parameters.Add("timespan",timespan)
	parameters.Add("perPage", perPage)
	parameters.Add("startingAfter", startingAfter)
	parameters.Add("endingBefore", endingBefore)
	parameters.Add("url", url)

	session.Request.URL.RawQuery = parameters.Encode()

	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}