package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
	"time"
)

type DeviceCerts []struct {
	Name           string    `json:"name"`
	NotValidAfter  time.Time `json:"notValidAfter"`
	NotValidBefore time.Time `json:"notValidBefore"`
	CertPem        string    `json:"certPem"`
	DeviceID       string    `json:"deviceId"`
	Issuer         string    `json:"issuer"`
	Subject        string    `json:"subject"`
	ID             string    `json:"id"`
}

type DeviceProfiles []struct {
	DeviceID    string `json:"deviceId"`
	ID          string `json:"id"`
	IsEncrypted bool   `json:"isEncrypted"`
	IsManaged   bool   `json:"isManaged"`
	ProfileData struct {
	} `json:"profileData"`
	ProfileIdentifier string `json:"profileIdentifier"`
	Name              string `json:"name"`
	Version           string `json:"version"`
}

type NetworkAdapters []struct {
	DhcpServer string `json:"dhcpServer"`
	DNSServer  string `json:"dnsServer"`
	Gateway    string `json:"gateway"`
	ID         string `json:"id"`
	IP         string `json:"ip"`
	Mac        string `json:"mac"`
	Name       string `json:"name"`
	Subnet     string `json:"subnet"`
}

type DeviceRestrictions []struct {
	Profile      string `json:"profile"`
	Restrictions struct {
		MyRestriction struct {
			Value bool `json:"value"`
		} `json:"myRestriction"`
	} `json:"restrictions"`
}

type DeviceSecurityCenters []struct {
	IsRooted             bool   `json:"isRooted"`
	HasAntiVirus         bool   `json:"hasAntiVirus"`
	AntiVirusName        string `json:"antiVirusName"`
	IsFireWallEnabled    bool   `json:"isFireWallEnabled"`
	HasFireWallInstalled bool   `json:"hasFireWallInstalled"`
	FireWallName         string `json:"fireWallName"`
	IsDiskEncrypted      bool   `json:"isDiskEncrypted"`
	IsAutoLoginDisabled  bool   `json:"isAutoLoginDisabled"`
	ID                   string `json:"id"`
	RunningProcs         string `json:"runningProcs"`
}

type DeviceAssociatedSoftware []struct {
	AppID             string      `json:"appId"`
	BundleSize        interface{} `json:"bundleSize"`
	CreatedAt         time.Time   `json:"createdAt"`
	DeviceID          string      `json:"deviceId"`
	DynamicSize       interface{} `json:"dynamicSize"`
	ID                string      `json:"id"`
	Identifier        string      `json:"identifier"`
	InstalledAt       time.Time   `json:"installedAt"`
	ToInstall         interface{} `json:"toInstall"`
	IosRedemptionCode interface{} `json:"iosRedemptionCode"`
	IsManaged         bool        `json:"isManaged"`
	ItunesID          interface{} `json:"itunesId"`
	LicenseKey        interface{} `json:"licenseKey"`
	Name              string      `json:"name"`
	Path              string      `json:"path"`
	RedemptionCode    interface{} `json:"redemptionCode"`
	ShortVersion      interface{} `json:"shortVersion"`
	Status            interface{} `json:"status"`
	ToUninstall       bool        `json:"toUninstall"`
	UninstalledAt     interface{} `json:"uninstalledAt"`
	UpdatedAt         time.Time   `json:"updatedAt"`
	Vendor            string      `json:"vendor"`
	Version           string      `json:"version"`
}

type SMEnrolledDevices struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Tags         []string `json:"tags"`
	Ssid         string   `json:"ssid"`
	WifiMac      string   `json:"wifiMac"`
	OsName       string   `json:"osName"`
	SystemModel  string   `json:"systemModel"`
	UUID         string   `json:"uuid"`
	SerialNumber string   `json:"serialNumber"`
	Serial       string   `json:"serial"`
	IP           string   `json:"ip"`
	Notes        string   `json:"notes"`
}

// List the certs on a device
func GetDeviceCerts(networkId, deviceId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/certs",
		api.BaseUrl(), networkId, deviceId)
	var datamodel = DeviceCerts{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Get the profiles associated with a device
func GetDeviceProfiles(networkId, deviceId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/deviceProfiles",
		api.BaseUrl(), networkId, deviceId)
	var datamodel = DeviceProfiles{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List the network adapters of a device
func GetNetworkAdapters(networkId, deviceId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/networkAdapters",
		api.BaseUrl(), networkId, deviceId)
	var datamodel = NetworkAdapters{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List the restrictions on a device
func GetDeviceRestrictions(networkId, deviceId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/restrictions",
		api.BaseUrl(), networkId, deviceId)
	var datamodel = DeviceRestrictions{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List the security centers on a device
func GetDeviceSecurityCenters(networkId, deviceId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/securityCenters",
		api.BaseUrl(), networkId, deviceId)
	var datamodel = DeviceSecurityCenters{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Get a list of softwares associated with a device
func GetDeviceAssociatedSoftware(networkId, deviceId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/softwares",
		api.BaseUrl(), networkId, deviceId)

	var datamodel = DeviceAssociatedSoftware{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

type DeviceSSIDNames []struct {
	CreatedAt time.Time `json:"createdAt"`
	ID        string    `json:"id"`
	XML       string    `json:"xml"`
}

// List the saved SSID names on a device
func GetDeviceSSIDNames(networkId, deviceId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/wlanLists",
		api.BaseUrl(), networkId, deviceId)

	var datamodel = DeviceSSIDNames{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List The Devices Enrolled In An SM Network With Various Specified Fields And Filters
func GetSMEnrolledDevices(networkId, fields, wifiMacs, serials, ids, scope, perPage,
	startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices", api.BaseUrl(), networkId)
	var datamodel = SMEnrolledDevices{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"fields":        fields,
		"wifiMacs":      wifiMacs,
		"serials":       serials,
		"ids":           ids,
		"scope":         scope,
		"perPage":       perPage,
		"startingAfter": startingAfter,
		"endingBefore":  endingBefore}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
