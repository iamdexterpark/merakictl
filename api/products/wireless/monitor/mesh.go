package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type MeshStatuses struct {
	Serial                string   `json:"serial"`
	MeshRoute             []string `json:"meshRoute"`
	LatestMeshPerformance struct {
		Mbps            int    `json:"mbps"`
		Metric          int    `json:"metric"`
		UsagePercentage string `json:"usagePercentage"`
	} `json:"latestMeshPerformance"`
}

// List wireless mesh statuses for repeaters
func GetMeshStatuses(networkId, perPage, startingAfter, endingBefore string) (MeshStatuses, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/meshStatuses",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("perPage", perPage)
	parameters.Add("startingAfter", startingAfter)
	parameters.Add("endingBefore", endingBefore)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = MeshStatuses{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}