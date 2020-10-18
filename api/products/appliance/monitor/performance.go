package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type PerformanceScore struct {
	PerfScore int `json:"perfScore"`
}

// Return the performance score for a single MX. Only primary MX devices supported.
// If no data is available, a 204 error code is returned.
func GetPerformanceScore(serial string ) (PerformanceScore, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/appliance/performance", api.BaseUrl(), serial)
	var payload io.ReadSeeker
	var results = PerformanceScore{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}