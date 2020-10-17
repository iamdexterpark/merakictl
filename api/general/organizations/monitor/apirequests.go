package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
	"time"
)

type APIRequestsOverview struct {
	ResponseCodeCounts struct {
		Num200 int `json:"200"`
		Num201 int `json:"201"`
		Num204 int `json:"204"`
		Num400 int `json:"400"`
		Num404 int `json:"404"`
		Num429 int `json:"429"`
	} `json:"responseCodeCounts"`
}

// Return an aggregated overview of API requests data
func GetAPIRequestsOverview(organizationId, t0, t1, timespan string ) (APIRequestsOverview, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/apiRequests/overview", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker
	var results = APIRequestsOverview{}

	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("t1", t1)
	parameters.Add("timespan",timespan)
	session.Request.URL.RawQuery = parameters.Encode()


	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type APIRequests []struct {
	APIRequest
}
type APIRequest struct {
	AdminID      string    `json:"adminId"`
	Method       string    `json:"method"`
	Host         string    `json:"host"`
	Path         string    `json:"path"`
	QueryString  string    `json:"queryString"`
	UserAgent    string    `json:"userAgent"`
	Ts           time.Time `json:"ts"`
	ResponseCode int       `json:"responseCode"`
	SourceIP     string    `json:"sourceIp"`
}

// Return an aggregated overview of API requests data
func GetAPIRequests(organizationId, t0, t1, timespan, perPage, startingAfter, endingBefore,
	adminId, path, method, responseCode, sourceIp  string ) (APIRequests, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/apiRequests", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker
	var results = APIRequests{}

	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("t1", t1)
	parameters.Add("timespan",timespan)
	parameters.Add("perPage", perPage)
	parameters.Add("startingAfter", startingAfter)
	parameters.Add("endingBefore", endingBefore)
	parameters.Add("adminId", adminId)
	parameters.Add("path", path)
	parameters.Add("method", method)
	parameters.Add("responseCode", responseCode)
	parameters.Add("sourceIp", sourceIp)


	session.Request.URL.RawQuery = parameters.Encode()


	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}