package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
	"time"
)

type MerakiAuthUsers []struct {
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
func GetMerakiAuthUsers(networkId string) (MerakiAuthUsers, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/merakiAuthUsers", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = MerakiAuthUsers{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}