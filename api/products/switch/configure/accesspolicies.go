package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type AccessPolicies []struct {
	AccessPolicy
}

type AccessPolicy struct {
	AccessPolicyNumber string `json:"accessPolicyNumber"`
	Name               string `json:"name"`
	RadiusServers      []struct {
		Host   string `json:"host"`
		Port   int    `json:"port"`
		Secret string `json:"secret"`
	} `json:"radiusServers"`
	RadiusTestingEnabled    bool `json:"radiusTestingEnabled"`
	RadiusCoaSupportEnabled bool `json:"radiusCoaSupportEnabled"`
	RadiusAccountingEnabled bool `json:"radiusAccountingEnabled"`
	RadiusAccountingServers []struct {
		Host   string `json:"host"`
		Port   int    `json:"port"`
		Secret string `json:"secret"`
	} `json:"radiusAccountingServers"`
	HostMode                       string `json:"hostMode"`
	AccessPolicyType               string `json:"accessPolicyType"`
	IncreaseAccessSpeed            bool   `json:"increaseAccessSpeed"`
	GuestVlanID                    int    `json:"guestVlanId"`
	VoiceVlanClients               bool   `json:"voiceVlanClients"`
	URLRedirectWalledGardenEnabled bool   `json:"urlRedirectWalledGardenEnabled"`
	URLRedirectWalledGardenRanges  string `json:"urlRedirectWalledGardenRanges"`
}

// List The Access Policies For A Switch Network
func GetAccessPolicies(networkId string) (AccessPolicies, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/accessPolicies",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = AccessPolicies{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// Return A Specific Access Policy For A Switch Network
func GetAccessPolicy(networkId, accessPolicyNumber string) (AccessPolicies, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/accessPolicies/%s",
		api.BaseUrl(), networkId, accessPolicyNumber)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = AccessPolicies{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
