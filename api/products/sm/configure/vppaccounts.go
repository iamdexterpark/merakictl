package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type VPPAccounts []struct {
	VPPAccount
}

type VPPAccount struct {
	ID              string `json:"id"`
	VppServiceToken string `json:"vppServiceToken"`
}

// Get A Hash Containing The Unparsed Token Of The VPP Account With The Given ID
func GetVPPAccount(organizationId, vppAccountId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/sm/vppAccounts/%s",
		api.BaseUrl(), organizationId, vppAccountId)
	var datamodel = VPPAccount{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List The VPP Accounts In The Organization
func GetVPPAccounts(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/sm/vppAccounts",
		api.BaseUrl(), organizationId)
	var datamodel = VPPAccounts{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
