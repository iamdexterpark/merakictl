package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetAPIRequestsOverview(organizationId, t0, t1, timespan string ) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/apiRequests/overview", api.BaseUrl(), organizationId)
	var datamodel = APIRequestsOverview{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0": t0,
		"t1": t1,
		"timespan": timespan}


	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return an aggregated overview of API requests data
func GetAPIRequests(organizationId, t0, t1, timespan, perPage, startingAfter, endingBefore,
	adminId, path, method, responseCode, sourceIp  string )[]api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/apiRequests", api.BaseUrl(), organizationId)
	var datamodel = APIRequests{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0": t0,
		"t1": t1,
		"timespan": timespan,
		"perPage": perPage,
		"startingAfter": startingAfter,
		"endingBefore": endingBefore,
		"adminId": adminId,
		"path": path,
		"method": method,
		"responseCode": responseCode,
		"sourceIp": sourceIp,
	}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}