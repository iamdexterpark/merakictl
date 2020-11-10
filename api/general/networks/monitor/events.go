package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
	"time"
)

// Events - List the events for the network
type Events struct {
	Message     interface{} `json:"message"`
	PageStartAt time.Time   `json:"pageStartAt"`
	PageEndAt   time.Time   `json:"pageEndAt"`
	Events      []struct {
		OccurredAt        time.Time `json:"occurredAt"`
		NetworkID         string    `json:"networkId"`
		Type              string    `json:"type"`
		Description       string    `json:"description"`
		ClientID          string    `json:"clientId"`
		ClientDescription string    `json:"clientDescription"`
		DeviceSerial      string    `json:"deviceSerial"`
		DeviceName        string    `json:"deviceName"`
		SsidNumber        int       `json:"ssidNumber"`
		SsidName          string    `json:"ssidName"`
		EventData         struct {
			Radio     string `json:"radio"`
			Vap       string `json:"vap"`
			ClientMac string `json:"client_mac"`
			ClientIP  string `json:"client_ip"`
			Channel   string `json:"channel"`
			Rssi      string `json:"rssi"`
			Aid       string `json:"aid"`
		} `json:"eventData"`
	} `json:"events"`
}

// List the events for the network
func GetEvents(networkId, productType, includedEventTypes, excludedEventTypes,
	deviceMac, deviceSerial, deviceName, clientIp, clientMac, clientName,
	smDeviceMac, smDeviceName, perPage, startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/events", api.BaseUrl(), networkId)
	var datamodel = Events{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"productType": productType,
		"includedEventTypes": includedEventTypes,
		"excludedEventTypes": excludedEventTypes,
		"deviceMac": deviceMac,
		"deviceSerial": deviceSerial,
		"deviceName": deviceName,
		"clientIp": clientIp,
		"clientMac": clientMac,
		"clientName": clientName,
		"smDeviceMac": smDeviceMac,
		"smDeviceName": smDeviceName,
		"perPage": perPage,
		"startingAfter": startingAfter,
		"endingBefore": endingBefore}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}