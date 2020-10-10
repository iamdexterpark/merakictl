package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type ClientPolicy struct {
	Mac           string `json:"mac"`
	DevicePolicy  string `json:"devicePolicy"`
	GroupPolicyID string `json:"groupPolicyId"`
}

// GetClientPolicy -  Return The Policy Assigned To A Client On The Network
func GetClientPolicy(networkId, clientId string) (ClientPolicy, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/clients/%s/policy", api.BaseUrl(), networkId, clientId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var clientpolicy = ClientPolicy{}
	user_agent.UnMarshalJSON(session.Body, &clientpolicy)
	traceback := user_agent.TraceBack(session)
	return clientpolicy, traceback
}
