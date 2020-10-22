package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
	"time"
)

type ClientCellularData []struct {
	Received int       `json:"received"`
	Sent     int       `json:"sent"`
	Ts       time.Time `json:"ts"`
}

// Return the client's daily cellular data usage history
func GetClientCellularData(networkId, deviceId string) (ClientCellularData, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/cellularUsageHistory",
		api.BaseUrl(), networkId, deviceId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = ClientCellularData{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


type HistoricalConnectivityData []struct {
	FirstSeenAt time.Time `json:"firstSeenAt"`
	LastSeenAt  time.Time `json:"lastSeenAt"`
}

// Returns historical connectivity data (whether a device is regularly checking in to Dashboard).
func GetHistoricalConnectivityData(networkId, deviceId,
	perPage, startingAfter, endingBefore string) (HistoricalConnectivityData, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm//devices/%s/connectivity",
		api.BaseUrl(), networkId, deviceId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("perPage", perPage)
	parameters.Add("startingAfter", startingAfter)
	parameters.Add("endingBefore", endingBefore )
	session.Request.URL.RawQuery = parameters.Encode()

	var results = HistoricalConnectivityData{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type DesktopDevicesHistoricalRecords []struct {
	MeasuredAt    time.Time `json:"measuredAt"`
	User          string    `json:"user"`
	NetworkDevice string    `json:"networkDevice"`
	NetworkDriver string    `json:"networkDriver"`
	WifiChannel   string    `json:"wifiChannel"`
	WifiAuth      string    `json:"wifiAuth"`
	WifiBssid     string    `json:"wifiBssid"`
	WifiSsid      string    `json:"wifiSsid"`
	WifiRssi      string    `json:"wifiRssi"`
	WifiNoise     string    `json:"wifiNoise"`
	DhcpServer    string    `json:"dhcpServer"`
	IP            string    `json:"ip"`
	NetworkMTU    string    `json:"networkMTU"`
	Subnet        string    `json:"subnet"`
	Gateway       string    `json:"gateway"`
	PublicIP      string    `json:"publicIP"`
	DNSServer     string    `json:"dnsServer"`
	Ts            time.Time `json:"ts"`
}

// Return historical records of various Systems Manager network connection details for desktop devices.
func GetDesktopDevicesHistoricalRecords(networkId, deviceId,
	perPage, startingAfter, endingBefore string) (DesktopDevicesHistoricalRecords, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm//devices/%s/desktopLogs",
		api.BaseUrl(), networkId, deviceId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("perPage", perPage)
	parameters.Add("startingAfter", startingAfter)
	parameters.Add("endingBefore", endingBefore )
	session.Request.URL.RawQuery = parameters.Encode()

	var results = DesktopDevicesHistoricalRecords{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


// Return historical records of commands sent to Systems Manager devices

// Return historical records of various Systems Manager client metrics for desktop devices.


