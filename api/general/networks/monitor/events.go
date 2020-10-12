package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
	"time"
)

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
	smDeviceMac, smDeviceName, perPage, startingAfter, endingBefore string) (Events, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/events", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("productType", productType)
	parameters.Add("includedEventTypes", includedEventTypes)
	parameters.Add("excludedEventTypes", excludedEventTypes)
	parameters.Add("deviceMac", deviceMac)
	parameters.Add("deviceSerial", deviceSerial)
	parameters.Add("deviceName", deviceName)
	parameters.Add("clientIp", clientIp)
	parameters.Add("clientMac", clientMac)
	parameters.Add("clientName", clientName)
	parameters.Add("smDeviceMac", smDeviceMac)
	parameters.Add("smDeviceName", smDeviceName)
	parameters.Add("perPage",perPage)
	parameters.Add("startingAfter",startingAfter)
	parameters.Add("endingBefore", endingBefore)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = Events{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}