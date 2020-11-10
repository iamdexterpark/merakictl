package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type ClientPolicy struct {
	Mac           string `json:"mac"`
	DevicePolicy  string `json:"devicePolicy"`
	GroupPolicyID string `json:"groupPolicyId"`
}

type SplashAuthorization struct {
	Ssids struct {
		Num0 struct {
			IsAuthorized bool   `json:"isAuthorized"`
			AuthorizedAt string `json:"authorizedAt"`
			ExpiresAt    string `json:"expiresAt"`
		} `json:"0"`
		Num2 struct {
			IsAuthorized bool `json:"isAuthorized"`
		} `json:"2"`
	} `json:"ssids"`
}

// GetClientPolicy -  Return The Policy Assigned To A Client On The Network
func GetClientPolicy(networkId, clientId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/clients/%s/policy", api.BaseUrl(), networkId, clientId)
	var datamodel = ClientPolicy{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return The Splash Authorization For A Client For Each SSID They've Associated With Through Splash
func GetSplashAuthorization(networkId, clientId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/clients/%s/splashAuthorizationStatus", api.BaseUrl(), networkId, clientId)
	var datamodel = SplashAuthorization{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
