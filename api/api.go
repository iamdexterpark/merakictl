package api

import (
	"fmt"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

// Organization Data Model
type Organization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Organizations Data Model
type Organizations []struct {
	Organization
}

// GetOrganizations - List the organizations that the user has privileges on
func GetOrganizations() (Organizations, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations", BaseUrl())
	var payload io.ReadSeeker
	session := Session(baseurl, "GET", payload)

	var results = Organizations{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type Network struct {
	ID               string   `json:"id"`
	OrganizationID   string   `json:"organizationId"`
	Name             string   `json:"name"`
	TimeZone         string   `json:"timeZone"`
	Tags             []string `json:"tags"`
	ProductTypes     []string `json:"productTypes"`
	EnrollmentString string   `json:"enrollmentString"`
}

// List the networks that the user has privileges on in an organization
func GetOrganizationNetworks(organizationId, configTemplateId, tagsFilterType,
	startingAfter, endingBefore string, tags, perPage string ) (Network, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/networks", BaseUrl(), organizationId)
	var payload io.ReadSeeker

	var results = Network{}

	session := Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("configTemplateId", configTemplateId)
	parameters.Add("tags", tags)
	parameters.Add("tagsFilterType",tagsFilterType)
	parameters.Add("perPage",perPage)
	parameters.Add("startingAfter",startingAfter)
	parameters.Add("endingBefore", endingBefore)
	session.Request.URL.RawQuery = parameters.Encode()


	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


// Devices Data Model
type Devices []struct {
	Name      string `json:"name"`
	Serial    string `json:"serial"`
	Mac       string `json:"mac"`
	Status    string `json:"status"`
	LanIP     string `json:"lanIp"`
	PublicIP  string `json:"publicIp"`
	NetworkID string `json:"networkId"`
}

// List the devices in an organization
func GetOrganizationDevices(organizationId string) (Devices, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/devices", BaseUrl(), organizationId)
	var payload io.ReadSeeker
	session := Session(baseurl, "GET", payload)

	var results = Devices{}
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
// List the Clients in a Network
func GetNetworkClients(networkId, t0, t1, timespan,
	perPage, startingAfter, endingBefore string) (Clients, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/clients", BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("t1", t1)
	parameters.Add("timespan", timespan)
	parameters.Add("perPage",perPage)
	parameters.Add("startingAfter",startingAfter)
	parameters.Add("endingBefore", endingBefore)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = Clients{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
