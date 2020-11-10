package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type GroupPolicies []struct {
	GroupPolicy
}

type GroupPolicy struct {
	Name          string `json:"name"`
	GroupPolicyID string `json:"groupPolicyId"`
	Scheduling    struct {
		Enabled bool `json:"enabled"`
		Monday  struct {
			Active bool   `json:"active"`
			From   string `json:"from"`
			To     string `json:"to"`
		} `json:"monday"`
		Tuesday struct {
			Active bool   `json:"active"`
			From   string `json:"from"`
			To     string `json:"to"`
		} `json:"tuesday"`
		Wednesday struct {
			Active bool   `json:"active"`
			From   string `json:"from"`
			To     string `json:"to"`
		} `json:"wednesday"`
		Thursday struct {
			Active bool   `json:"active"`
			From   string `json:"from"`
			To     string `json:"to"`
		} `json:"thursday"`
		Friday struct {
			Active bool   `json:"active"`
			From   string `json:"from"`
			To     string `json:"to"`
		} `json:"friday"`
		Saturday struct {
			Active bool   `json:"active"`
			From   string `json:"from"`
			To     string `json:"to"`
		} `json:"saturday"`
		Sunday struct {
			Active bool   `json:"active"`
			From   string `json:"from"`
			To     string `json:"to"`
		} `json:"sunday"`
	} `json:"scheduling"`
	Bandwidth struct {
		Settings        string `json:"settings"`
		BandwidthLimits struct {
			LimitUp   int `json:"limitUp"`
			LimitDown int `json:"limitDown"`
		} `json:"bandwidthLimits"`
	} `json:"bandwidth"`
	FirewallAndTrafficShaping struct {
		Settings            string `json:"settings"`
		TrafficShapingRules []struct {
			Definitions []struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"definitions"`
			PerClientBandwidthLimits struct {
				Settings        string `json:"settings"`
				BandwidthLimits struct {
					LimitUp   int `json:"limitUp"`
					LimitDown int `json:"limitDown"`
				} `json:"bandwidthLimits"`
			} `json:"perClientBandwidthLimits"`
			DscpTagValue interface{} `json:"dscpTagValue"`
			PcpTagValue  interface{} `json:"pcpTagValue"`
		} `json:"trafficShapingRules"`
		L3FirewallRules []struct {
			Comment  string `json:"comment"`
			Policy   string `json:"policy"`
			Protocol string `json:"protocol"`
			DestPort int    `json:"destPort"`
			DestCidr string `json:"destCidr"`
		} `json:"l3FirewallRules"`
		L7FirewallRules []struct {
			Policy string `json:"policy"`
			Type   string `json:"type"`
			Value  struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"value"`
		} `json:"l7FirewallRules"`
	} `json:"firewallAndTrafficShaping"`
	ContentFiltering struct {
		AllowedURLPatterns struct {
			Settings string        `json:"settings"`
			Patterns []interface{} `json:"patterns"`
		} `json:"allowedUrlPatterns"`
		BlockedURLPatterns struct {
			Settings string   `json:"settings"`
			Patterns []string `json:"patterns"`
		} `json:"blockedUrlPatterns"`
		BlockedURLCategories struct {
			Settings   string   `json:"settings"`
			Categories []string `json:"categories"`
		} `json:"blockedUrlCategories"`
	} `json:"contentFiltering"`
	SplashAuthSettings string `json:"splashAuthSettings"`
	VlanTagging        struct {
		Settings string `json:"settings"`
		VlanID   string `json:"vlanId"`
	} `json:"vlanTagging"`
	BonjourForwarding struct {
		Settings string `json:"settings"`
		Rules    []struct {
			Description string   `json:"description"`
			VlanID      string   `json:"vlanId"`
			Services    []string `json:"services"`
		} `json:"rules"`
	} `json:"bonjourForwarding"`
}

// List The Group Policies In A Network
func GetGroupPolicies(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/groupPolicies", api.BaseUrl(), networkId)
	var datamodel = GroupPolicies{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Display A Group Policy
func GetGroupPolicy(networkId, groupPolicyId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/groupPolicies/%s", api.BaseUrl(), networkId, groupPolicyId)
	var datamodel = GroupPolicy{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
