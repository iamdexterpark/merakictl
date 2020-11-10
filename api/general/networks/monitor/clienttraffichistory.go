package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
	"time"
)

// ClientNetworkTraffic - Return the client's network traffic data over time
type ClientNetworkTraffic []struct {
	Ts            time.Time `json:"ts"`
	Application   string    `json:"application"`
	Destination   string    `json:"destination"`
	Protocol      string    `json:"protocol"`
	Port          int       `json:"port"`
	Recv          int       `json:"recv"`
	Sent          int       `json:"sent"`
	NumFlows      int       `json:"numFlows"`
	ActiveSeconds int       `json:"activeSeconds"`
}

// ClientUsageHistory - Return the client's daily usage history
type ClientUsageHistory []struct {
	Sent     int       `json:"sent"`
	Received int       `json:"received"`
	Ts       time.Time `json:"ts"`
}

// Clients - List The Clients That Have Used This Network In The Timespan
type Clients []struct {
	Usage struct {
		Sent int `json:"sent"`
		Recv int `json:"recv"`
	} `json:"usage"`
	ID                 string      `json:"id"`
	Description        string      `json:"description"`
	Mac                string      `json:"mac"`
	IP                 string      `json:"ip"`
	User               string      `json:"user"`
	Vlan               int         `json:"vlan"`
	Switchport         interface{} `json:"switchport"`
	IP6                string      `json:"ip6"`
	FirstSeen          int         `json:"firstSeen"`
	LastSeen           int         `json:"lastSeen"`
	Manufacturer       string      `json:"manufacturer"`
	Os                 string      `json:"os"`
	RecentDeviceSerial string      `json:"recentDeviceSerial"`
	RecentDeviceName   string      `json:"recentDeviceName"`
	RecentDeviceMac    string      `json:"recentDeviceMac"`
	Ssid               string      `json:"ssid"`
	Status             string      `json:"status"`
	Notes              string      `json:"notes"`
	IP6Local           string      `json:"ip6Local"`
	SmInstalled        bool        `json:"smInstalled"`
	GroupPolicy8021X   string      `json:"groupPolicy8021x"`
}


// AssociatedClient - Return The Client Associated With The Given Identifier
type AssociatedClient struct {
	ID                   string      `json:"id"`
	Description          string      `json:"description"`
	Mac                  string      `json:"mac"`
	IP                   string      `json:"ip"`
	User                 string      `json:"user"`
	Vlan                 string      `json:"vlan"`
	Switchport           interface{} `json:"switchport"`
	IP6                  string      `json:"ip6"`
	FirstSeen            int         `json:"firstSeen"`
	LastSeen             int         `json:"lastSeen"`
	Manufacturer         string      `json:"manufacturer"`
	Os                   string      `json:"os"`
	Ssid                 string      `json:"ssid"`
	WirelessCapabilities string      `json:"wirelessCapabilities"`
	SmInstalled          bool        `json:"smInstalled"`
	RecentDeviceMac      string      `json:"recentDeviceMac"`
	ClientVpnConnections []struct {
		RemoteIP       string `json:"remoteIp"`
		ConnectedAt    int    `json:"connectedAt"`
		DisconnectedAt int    `json:"disconnectedAt"`
	} `json:"clientVpnConnections"`
	Lldp   [][]string  `json:"lldp"`
	Cdp    interface{} `json:"cdp"`
	Status string      `json:"status"`
}


// Return the client's network traffic data over time
func GetClientNetworkTraffic(networkId, clientId, trafficHistory, perPage,
	startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/clients/%s/trafficHistory", api.BaseUrl(), networkId, clientId)
	var datamodel = ClientNetworkTraffic{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"trafficHistory": trafficHistory,
		"perPage": perPage,
		"startingAfter": startingAfter,
		"endingBefore": endingBefore}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}


// Return the client's daily usage history
func GetClientUsageHistory(networkId, clientId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/clients/%s/usageHistory", api.BaseUrl(), networkId, clientId)
	var datamodel = ClientUsageHistory{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}



// List The Clients That Have Used This Network In The Timespan
func GetClients(networkId, t0, timespan, perPage, startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/clients", api.BaseUrl(), networkId)
	var datamodel = Clients{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0": t0,
		"timespan": timespan,
		"perPage": perPage,
		"startingAfter": startingAfter,
		"endingBefore": endingBefore}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}


// Return The Client Associated With The Given Identifier
func GetAssociatedClient(networkId, clientId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/clients/%s", api.BaseUrl(), networkId, clientId)
	var datamodel = AssociatedClient{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}