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