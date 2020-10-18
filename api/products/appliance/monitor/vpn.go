package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type VPNHistory []struct {
	NetworkID      string `json:"networkId"`
	NetworkName    string `json:"networkName"`
	MerakiVpnPeers []struct {
		NetworkID    string `json:"networkId"`
		NetworkName  string `json:"networkName"`
		UsageSummary struct {
			ReceivedInKilobytes int `json:"receivedInKilobytes"`
			SentInKilobytes     int `json:"sentInKilobytes"`
		} `json:"usageSummary"`
		LatencySummaries []struct {
			SenderUplink   string `json:"senderUplink"`
			ReceiverUplink string `json:"receiverUplink"`
			AvgLatencyMs   int    `json:"avgLatencyMs"`
			MinLatencyMs   int    `json:"minLatencyMs"`
			MaxLatencyMs   int    `json:"maxLatencyMs"`
		} `json:"latencySummaries"`
		LossPercentageSummaries []struct {
			SenderUplink      string  `json:"senderUplink"`
			ReceiverUplink    string  `json:"receiverUplink"`
			AvgLossPercentage int     `json:"avgLossPercentage"`
			MinLossPercentage int     `json:"minLossPercentage"`
			MaxLossPercentage float64 `json:"maxLossPercentage"`
		} `json:"lossPercentageSummaries"`
		JitterSummaries []struct {
			SenderUplink   string  `json:"senderUplink"`
			ReceiverUplink string  `json:"receiverUplink"`
			AvgJitter      float64 `json:"avgJitter"`
			MinJitter      int     `json:"minJitter"`
			MaxJitter      float64 `json:"maxJitter"`
		} `json:"jitterSummaries"`
		MosSummaries []struct {
			SenderUplink   string  `json:"senderUplink"`
			ReceiverUplink string  `json:"receiverUplink"`
			AvgMos         float64 `json:"avgMos"`
			MinMos         int     `json:"minMos"`
			MaxMos         float64 `json:"maxMos"`
		} `json:"mosSummaries"`
	} `json:"merakiVpnPeers"`
}

// Show VPN history stat for networks in an organization
func GetVPNHistory(organizationId, t0, t1, timespan, perPage, startingAfter, endingBefore,
	networkIds string) (VPNHistory, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/appliance/vpn/stats", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("t1", t1)
	parameters.Add("timespan", timespan)
	parameters.Add("perPage",perPage)
	parameters.Add("startingAfter",startingAfter)
	parameters.Add("endingBefore", endingBefore)
	parameters.Add("networkIds", networkIds)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = VPNHistory{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}




// Show VPN status for networks in an organization