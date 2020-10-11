package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type PiiKeys struct {
	N1234 struct {
		Macs          []string `json:"macs"`
		Emails        []string `json:"emails"`
		Usernames     []string `json:"usernames"`
		Serials       []string `json:"serials"`
		Imeis         []string `json:"imeis"`
		BluetoothMacs []string `json:"bluetoothMacs"`
	} `json:"N_1234"`
}

// List the keys required to access Personally Identifiable Information (PII) for a given identifier
func GetPiiKeys(networkId string) (PiiKeys, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/pii/piiKeys", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = PiiKeys{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type PiiRequest struct {
	ID               string `json:"id"`
	OrganizationWide bool   `json:"organizationWide"`
	NetworkID        string `json:"networkId"`
	Type             string `json:"type"`
	Mac              string `json:"mac"`
	Datasets         string `json:"datasets"`
	Status           string `json:"status"`
	CreatedAt        int    `json:"createdAt"`
	CompletedAt      int    `json:"completedAt"`
}

type PiiRequests []struct {
	PiiRequest
}

// List The PII Requests For This Network Or Organization
func GetPiiRequests(networkId string) (PiiRequests, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/pii/requests", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = PiiRequests{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// Return A PII Request
func GetPiiRequest(networkId, requestId string) (PiiRequest, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/pii/requests/%s", api.BaseUrl(), networkId, requestId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = PiiRequest{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type SmDevicesForKey  struct {
	N1234 []string `json:"N_1234"`
}

// Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier
func GetSmDevicesForKey(networkId string) (SmDevicesForKey, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/pii/smDevicesForKey", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = SmDevicesForKey{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}