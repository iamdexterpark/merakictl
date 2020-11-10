package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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

type SmDevicesForKey struct {
	N1234 []string `json:"N_1234"`
}

type SmOwnersForKey struct {
	N1234 []string `json:"N_1234"`
}

// List the keys required to access Personally Identifiable Information (PII) for a given identifier
func GetPiiKeys(networkId, username, email, mac, serial, imei, bluetoothMac string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/pii/piiKeys", api.BaseUrl(), networkId)
	var datamodel = PiiKeys{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"username":     username,
		"email":        email,
		"mac":          mac,
		"serial":       serial,
		"imei":         imei,
		"bluetoothMac": bluetoothMac}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List The PII Requests For This Network Or Organization
func GetPiiRequests(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/pii/requests", api.BaseUrl(), networkId)
	var datamodel = PiiRequests{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return A PII Request
func GetPiiRequest(networkId, requestId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/pii/requests/%s", api.BaseUrl(), networkId, requestId)
	var datamodel = PiiRequest{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier
func GetSmDevicesForKey(networkId, username, email, mac, serial, imei, bluetoothMac string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/pii/smDevicesForKey", api.BaseUrl(), networkId)
	var datamodel = SmDevicesForKey{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"username":     username,
		"email":        email,
		"mac":          mac,
		"serial":       serial,
		"imei":         imei,
		"bluetoothMac": bluetoothMac}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Given a piece of Personally Identifiable Information (PII), return the Systems Manager owner ID(s) associated with that identifier
func GetSmOwnersForKey(networkId, username, email, mac, serial, imei, bluetoothMac string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/pii/smOwnersForKey", api.BaseUrl(), networkId)
	var datamodel = SmOwnersForKey{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"username":     username,
		"email":        email,
		"mac":          mac,
		"serial":       serial,
		"imei":         imei,
		"bluetoothMac": bluetoothMac}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
