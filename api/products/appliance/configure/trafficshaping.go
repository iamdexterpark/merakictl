package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type PerformanceClasses []struct {
	PerformanceClass
}
type PerformanceClass struct {
	CustomPerformanceClassID string `json:"customPerformanceClassId"`
	Name                     string `json:"name"`
	MaxLatency               int    `json:"maxLatency"`
	MaxJitter                int    `json:"maxJitter"`
	MaxLossPercentage        int    `json:"maxLossPercentage"`
}

// List all custom performance classes for an MX network
func GetPerformanceClasses(networkId string ) (PerformanceClasses, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/trafficShaping/customPerformanceClasses", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = PerformanceClasses{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// Return a custom performance class for an MX network
func GetPerformanceClass(networkId, customPerformanceClassId string ) (PerformanceClass, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/trafficShaping/customPerformanceClasses/%s", api.BaseUrl(), networkId, customPerformanceClassId)
	var payload io.ReadSeeker
	var results = PerformanceClass{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}