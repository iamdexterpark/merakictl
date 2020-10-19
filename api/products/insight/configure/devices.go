package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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

// List the certs on a device
func GetDeviceCerts(networkId, deviceId string) (DeviceCerts, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/certs",
		api.BaseUrl(), networkId, deviceId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = DeviceCerts{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
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

// Get the profiles associated with a device
func GetDeviceProfiles(networkId, deviceId string) (DeviceProfiles, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/deviceProfiles",
		api.BaseUrl(), networkId, deviceId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = DeviceProfiles{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
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

// List the network adapters of a device
func GetNetworkAdapters(networkId, deviceId string) (NetworkAdapters, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/networkAdapters",
		api.BaseUrl(), networkId, deviceId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = NetworkAdapters{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


type DeviceRestrictions []struct {
	Profile      string `json:"profile"`
	Restrictions struct {
		MyRestriction struct {
			Value bool `json:"value"`
		} `json:"myRestriction"`
	} `json:"restrictions"`
}

// List the restrictions on a device
func GetDeviceRestrictions(networkId, deviceId string) (DeviceRestrictions, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/restrictions",
		api.BaseUrl(), networkId, deviceId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = DeviceRestrictions{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
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

// List the security centers on a device
func GetDeviceSecurityCenters(networkId, deviceId string) (DeviceSecurityCenters, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/securityCenters",
		api.BaseUrl(), networkId, deviceId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = DeviceSecurityCenters{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
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


// Get a list of softwares associated with a device
func GetDeviceAssociatedSoftware(networkId, deviceId string) (DeviceAssociatedSoftware, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/softwares",
		api.BaseUrl(), networkId, deviceId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = DeviceAssociatedSoftware{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type DeviceSSIDNames []struct {
	CreatedAt time.Time `json:"createdAt"`
	ID        string    `json:"id"`
	XML       string    `json:"xml"`
}

// List the saved SSID names on a device
func GetDeviceSSIDNames(networkId, deviceId string) (DeviceSSIDNames, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/wlanLists",
		api.BaseUrl(), networkId, deviceId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = DeviceSSIDNames{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
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


// List The Devices Enrolled In An SM Network With Various Specified Fields And Filters
func GetSMEnrolledDevices(networkId, fields, wifiMacs, serials, ids, scope, perPage,
	startingAfter, endingBefore string) (SMEnrolledDevices, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("fields", fields)
	parameters.Add("wifiMacs", wifiMacs)
	parameters.Add("serials", serials)
	parameters.Add("ids", ids)
	parameters.Add("scope", scope)
	parameters.Add("perPage", perPage)
	parameters.Add("startingAfter", startingAfter)
	parameters.Add("endingBefore", endingBefore)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = SMEnrolledDevices{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
