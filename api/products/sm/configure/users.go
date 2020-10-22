package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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

// Get the profiles associated with a user
func GetUserProfiles(networkId, userId string) (UserProfiles, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/users/%s/deviceProfiles",
		api.BaseUrl(), networkId, userId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = UserProfiles{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


type UserAssociatedSoftware[]struct {
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

// Get a list of softwares associated with a user
func GetUserAssociatedSoftware(networkId, userId string) (UserAssociatedSoftware, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/users/%s/softwares",
		api.BaseUrl(), networkId, userId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = UserAssociatedSoftware{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
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

//  List The Owners In An SM Network With Various Specified Fields And Filters
func GetSMNetworkOwners(networkId, ids, usernames, emails, scope string) (SMNetworkOwners, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/users",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("ids", ids)
	parameters.Add("usernames", usernames)
	parameters.Add("emails", emails)
	parameters.Add("scope", scope)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = SMNetworkOwners{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}