package api

import (
	"fmt"
	"log"
)

// Organization - List an organization that the user has privileges on
type Organization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Organizations - List the organizations that the user has privileges on
type Organizations []struct {
	Organization
}

// Network - List the networks that the user has privileges on in an organization
type Network struct {
	ID               string   `json:"id"`
	OrganizationID   string   `json:"organizationId"`
	Name             string   `json:"name"`
	TimeZone         string   `json:"timeZone"`
	Tags             []string `json:"tags"`
	ProductTypes     []string `json:"productTypes"`
	EnrollmentString string   `json:"enrollmentString"`
}

// Devices - List the devices in an organization
type Devices []struct {
	Name      string `json:"name"`
	Serial    string `json:"serial"`
	Mac       string `json:"mac"`
	Status    string `json:"status"`
	LanIP     string `json:"lanIp"`
	PublicIP  string `json:"publicIp"`
	NetworkID string `json:"networkId"`
}

// Clients - List the clients that have used this network in the timespan
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

// GetOrganizations - List the organizations that the user has privileges on
func GetOrganizations() []Results {
	baseurl := fmt.Sprintf("%s/organizations", BaseUrl())
	var datamodel = Organizations{}

	sessions, err := Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// GetOrganizationNetworks - List the networks that the user has privileges on in an organization
func GetOrganizationNetworks(organizationId, configTemplateId, tagsFilterType, startingAfter,
	endingBefore, tags, perPage  string) []Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/networks", BaseUrl(), organizationId)
	var datamodel = Network{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"configTemplateId": configTemplateId,
		"tags":             tags,
		"tagsFilterType":   tagsFilterType,
		"perPage":          perPage,
		"startingAfter":    startingAfter,
		"endingBefore":     endingBefore}

	sessions, err := Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}

// GetOrganizationDevices - List the devices in an organization
func GetOrganizationDevices(organizationId, perPage, startingAfter,
	configurationUpdatedAfter string) []Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/devices", BaseUrl(), organizationId)
	var datamodel = Devices{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"configurationUpdatedAfter": configurationUpdatedAfter,
		"perPage":                   perPage,
		"startingAfter":             startingAfter}

	sessions, err := Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}

// GetNetworkClients - List the Clients in a Network
func GetNetworkClients(networkId, t0, t1, timespan,
	perPage, startingAfter, endingBefore string) []Results {
	baseurl := fmt.Sprintf("%s/networks/%s/clients", BaseUrl(), networkId)
	var datamodel = Clients{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":            t0,
		"t1":            t1,
		"timespan":      timespan,
		"perPage":       perPage,
		"startingAfter": startingAfter,
		"endingBefore":  endingBefore}

	sessions, err := Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}
