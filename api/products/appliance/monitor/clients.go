package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
	"time"
)

type ClientSecurityEvents []struct {
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

// List the security events for a client.
// Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
func GetSecurityEvents(networkId, clientId, t0, t1, timespan, perPage, startingAfter, endingBefore,
	sortOrder string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/clients/%s/security/events", api.BaseUrl(), networkId, clientId)
	var datamodel = ClientSecurityEvents{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0": t0,
		"t1": t1,
		"timespan": timespan,
		"perPage": perPage,
		"startingAfter": startingAfter,
		"endingBefore": endingBefore,
		"sortOrder": sortOrder}


	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}