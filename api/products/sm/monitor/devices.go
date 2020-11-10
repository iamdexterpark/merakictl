package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
	"time"
)

type ClientCellularData []struct {
	Received int       `json:"received"`
	Sent     int       `json:"sent"`
	Ts       time.Time `json:"ts"`
}

type HistoricalConnectivityData []struct {
	FirstSeenAt time.Time `json:"firstSeenAt"`
	LastSeenAt  time.Time `json:"lastSeenAt"`
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

type CommandHistoricalRecords []struct {
	Action        string    `json:"action"`
	Name          string    `json:"name"`
	Details       string    `json:"details"`
	DashboardUser string    `json:"dashboardUser"`
	Ts            time.Time `json:"ts"`
}

type ClientMetricsHistoricalRecords []struct {
	CPUPercentUsed  float64 `json:"cpuPercentUsed"`
	MemFree         int     `json:"memFree"`
	MemWired        int     `json:"memWired"`
	MemActive       int     `json:"memActive"`
	MemInactive     int     `json:"memInactive"`
	NetworkSent     int     `json:"networkSent"`
	NetworkReceived int     `json:"networkReceived"`
	SwapUsed        int     `json:"swapUsed"`
	DiskUsage       struct {
		C struct {
			Used  int `json:"used"`
			Space int `json:"space"`
		} `json:"c"`
	} `json:"diskUsage"`
	Ts time.Time `json:"ts"`
}

// Return the client's daily cellular data usage history
func GetClientCellularData(networkId, deviceId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/cellularUsageHistory",
		api.BaseUrl(), networkId, deviceId)
	var datamodel = ClientCellularData{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Returns historical connectivity data (whether a device is regularly checking in to Dashboard).
func GetHistoricalConnectivityData(networkId, deviceId,
	perPage, startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm//devices/%s/connectivity",
		api.BaseUrl(), networkId, deviceId)
	var datamodel = HistoricalConnectivityData{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"perPage":       perPage,
		"startingAfter": startingAfter,
		"endingBefore":  endingBefore}
	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return historical records of various Systems Manager network connection details for desktop devices.
func GetDesktopDevicesHistoricalRecords(networkId, deviceId,
	perPage, startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm//devices/%s/desktopLogs",
		api.BaseUrl(), networkId, deviceId)
	var datamodel = DesktopDevicesHistoricalRecords{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"perPage":       perPage,
		"startingAfter": startingAfter,
		"endingBefore":  endingBefore}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return historical records of commands sent to Systems Manager devices
func GetCommandHistoricalRecords(networkId, deviceId,
	perPage, startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm//devices/%s/deviceCommandLogs",
		api.BaseUrl(), networkId, deviceId)
	var datamodel = CommandHistoricalRecords{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"perPage":       perPage,
		"startingAfter": startingAfter,
		"endingBefore":  endingBefore}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return historical records of various Systems Manager client metrics for desktop devices.
func GetClientMetricsHistoricalRecords(networkId, deviceId,
	perPage, startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm//devices/%s/performanceHistory",
		api.BaseUrl(), networkId, deviceId)
	var datamodel = ClientMetricsHistoricalRecords{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"perPage":       perPage,
		"startingAfter": startingAfter,
		"endingBefore":  endingBefore}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
