package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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
func GetIdentityPSKs(networkId, number string) (IdentityPSKs, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s/identityPsks",
		api.BaseUrl(), networkId, number)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var ssid IdentityPSKs
	user_agent.UnMarshalJSON(session.Body, &ssid)
	traceback := user_agent.TraceBack(session)
	return ssid, traceback
}


// Return An Identity PSK
func GetIdentityPSK(networkId, number, identityPskId string) (IdentityPSK, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s/identityPsks/%s",
		api.BaseUrl(), networkId, number, identityPskId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var ssid IdentityPSK
	user_agent.UnMarshalJSON(session.Body, &ssid)
	traceback := user_agent.TraceBack(session)
	return ssid, traceback
}