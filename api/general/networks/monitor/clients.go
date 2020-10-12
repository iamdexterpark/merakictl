package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
	"time"
)

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

// Return the client's network traffic data over time
func GetClientNetworkTraffic(networkId, clientId string) (ClientNetworkTraffic, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/clients/%s/trafficHistory", api.BaseUrl(), networkId, clientId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = ClientNetworkTraffic{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type ClientUsageHistory []struct {
	Sent     int       `json:"sent"`
	Received int       `json:"received"`
	Ts       time.Time `json:"ts"`
}

// Return the client's daily usage history
func GetClientUsageHistory(networkId, clientId string) (ClientUsageHistory, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/clients/%s/usageHistory", api.BaseUrl(), networkId, clientId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = ClientUsageHistory{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

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

// List The Clients That Have Used This Network In The Timespan
func GetClients(networkId string) (Clients, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/clients", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = Clients{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


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

// Return The Client Associated With The Given Identifier
func GetAssociatedClient(networkId, clientId string) (AssociatedClient, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/clients/%s", api.BaseUrl(), networkId, clientId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = AssociatedClient{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}