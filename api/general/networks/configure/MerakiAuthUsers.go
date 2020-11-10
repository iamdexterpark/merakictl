package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
	"time"
)

type MerakiAuthUsers []struct {
	MerakiAuthUser
}
type MerakiAuthUser struct {
	MerakiAuthUserID string    `json:"merakiAuthUserId"`
	Email            string    `json:"email"`
	Name             string    `json:"name"`
	CreatedAt        time.Time `json:"createdAt"`
	AccountType      string    `json:"accountType"`
	Authorizations   []struct {
		SsidNumber        int       `json:"ssidNumber"`
		AuthorizedZone    string    `json:"authorizedZone"`
		ExpiresAt         time.Time `json:"expiresAt"`
		AuthorizedByName  string    `json:"authorizedByName"`
		AuthorizedByEmail string    `json:"authorizedByEmail"`
	} `json:"authorizations"`
}

// List The Users Configured Under Meraki Authentication For A Network Splash Guest Or RADIUS Users For A Wireless Network Or Client VPN Users For A Wired Network
func GetMerakiAuthUsers(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/merakiAuthUsers", api.BaseUrl(), networkId)
	var datamodel = MerakiAuthUsers{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return The Meraki Auth Splash Guest RADIUS Or Client VPN User
func GetMerakiAuthUser(networkId, merakiAuthUserId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/merakiAuthUsers/%s", api.BaseUrl(), networkId, merakiAuthUserId)
	var datamodel = MerakiAuthUser{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
