package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetAccessPolicies(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/accessPolicies",
		api.BaseUrl(), networkId)
	var datamodel = AccessPolicies{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return A Specific Access Policy For A Switch Network
func GetAccessPolicy(networkId, accessPolicyNumber string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/accessPolicies/%s",
		api.BaseUrl(), networkId, accessPolicyNumber)
	var datamodel = AccessPolicies{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
