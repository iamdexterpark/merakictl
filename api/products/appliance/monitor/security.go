package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
	"time"
)

type SecurityEvents []struct {
	Ts              time.Time `json:"ts"`
	EventType       string    `json:"eventType"`
	ClientName      string    `json:"clientName,omitempty"`
	ClientMac       string    `json:"clientMac"`
	ClientIP        string    `json:"clientIp,omitempty"`
	SrcIP           string    `json:"srcIp"`
	DestIP          string    `json:"destIp"`
	Protocol        string    `json:"protocol"`
	URI             string    `json:"uri,omitempty"`
	CanonicalName   string    `json:"canonicalName,omitempty"`
	DestinationPort int       `json:"destinationPort,omitempty"`
	FileHash        string    `json:"fileHash,omitempty"`
	FileType        string    `json:"fileType,omitempty"`
	FileSizeBytes   int       `json:"fileSizeBytes,omitempty"`
	Disposition     string    `json:"disposition,omitempty"`
	Action          string    `json:"action,omitempty"`
	DeviceMac       string    `json:"deviceMac,omitempty"`
	Priority        string    `json:"priority,omitempty"`
	Classification  string    `json:"classification,omitempty"`
	Blocked         bool      `json:"blocked,omitempty"`
	Message         string    `json:"message,omitempty"`
	Signature       string    `json:"signature,omitempty"`
	SigSource       string    `json:"sigSource,omitempty"`
	RuleID          string    `json:"ruleId,omitempty"`
}

// List the security events for a network
func GetNetworkSecurityEvents(networkId, t0, t1, timespan, perPage,
	startingAfter, endingBefore, sortOrder string) (SecurityEvents, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/security/events", api.BaseUrl(), networkId)
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
	parameters.Add("sortOrder", sortOrder)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = SecurityEvents{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


// List the security events for an organization
func GetOrganizationSecurityEvents(organizationId, t0, t1, timespan, perPage,
	startingAfter, endingBefore, sortOrder string) (SecurityEvents, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/security/events", api.BaseUrl(), organizationId)
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
	parameters.Add("sortOrder", sortOrder)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = SecurityEvents{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}