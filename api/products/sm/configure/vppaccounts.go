package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type VPPAccounts []struct {
	VPPAccount
}

type VPPAccount struct {
	ID              string `json:"id"`
	VppServiceToken string `json:"vppServiceToken"`
}

// Get A Hash Containing The Unparsed Token Of The VPP Account With The Given ID
func GetVPPAccount(organizationId, vppAccountId string) (VPPAccount, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/sm/vppAccounts/%s",
		api.BaseUrl(), organizationId, vppAccountId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = VPPAccount{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


// List The VPP Accounts In The Organization
func GetVPPAccounts(organizationId string) (VPPAccounts, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/sm/vppAccounts",
		api.BaseUrl(), organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = VPPAccounts{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}