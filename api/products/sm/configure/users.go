package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
	"time"
)

type UserProfiles []struct {
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

type UserAssociatedSoftware []struct {
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

type SMNetworkOwners []struct {
	ID                     string        `json:"id"`
	Email                  string        `json:"email"`
	FullName               string        `json:"fullName"`
	Username               string        `json:"username"`
	HasPassword            bool          `json:"hasPassword"`
	Tags                   []string      `json:"tags"`
	AdGroups               []interface{} `json:"adGroups"`
	AsmGroups              []interface{} `json:"asmGroups"`
	IsExternal             bool          `json:"isExternal"`
	DisplayName            string        `json:"displayName"`
	HasIdentityCertificate bool          `json:"hasIdentityCertificate"`
	UserThumbnail          string        `json:"userThumbnail"`
}

// Get the profiles associated with a user
func GetUserProfiles(networkId, userId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/users/%s/deviceProfiles",
		api.BaseUrl(), networkId, userId)
	var datamodel = UserProfiles{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Get a list of softwares associated with a user
func GetUserAssociatedSoftware(networkId, userId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/users/%s/softwares",
		api.BaseUrl(), networkId, userId)
	var datamodel = UserAssociatedSoftware{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

//  List The Owners In An SM Network With Various Specified Fields And Filters
func GetSMNetworkOwners(networkId, ids, usernames, emails, scope string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/users",
		api.BaseUrl(), networkId)
	var datamodel = SMNetworkOwners{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"ids":       ids,
		"usernames": usernames,
		"emails":    emails,
		"scope":     scope}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
