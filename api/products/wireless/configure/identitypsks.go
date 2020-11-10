package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type IdentityPSKs []struct {
	IdentityPSK
}
type IdentityPSK struct {
	IdentityPskID string `json:"identityPskId"`
	Name          string `json:"name"`
	Passphrase    string `json:"passphrase"`
	GroupPolicyID string `json:"groupPolicyId"`
}

// List All Identity PSKs In A Wireless Network
func GetIdentityPSKs(networkId, number string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s/identityPsks",
		api.BaseUrl(), networkId, number)
	var datamodel IdentityPSKs
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return An Identity PSK
func GetIdentityPSK(networkId, number, identityPskId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s/identityPsks/%s",
		api.BaseUrl(), networkId, number, identityPskId)
	var datamodel IdentityPSK
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
