package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type QoSRuleIds struct {
	RuleIds []string `json:"ruleIds"`
}

type QoSRules []struct {
	QoSRule
}

type QoSRule struct {
	ID           string      `json:"id"`
	Vlan         int         `json:"vlan"`
	Protocol     string      `json:"protocol"`
	SrcPort      int         `json:"srcPort"`
	SrcPortRange interface{} `json:"srcPortRange"`
	DstPort      interface{} `json:"dstPort"`
	DstPortRange string      `json:"dstPortRange"`
	Dscp         int         `json:"dscp"`
}

// Return the quality of service rule IDs by order in which they will be processed by the switch
func GetQoSRuleIds(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/qosRules/order",
		api.BaseUrl(), networkId)
	var datamodel = QoSRuleIds{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List Quality Of Service Rules
func GetQoSRules(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/qosRules",
		api.BaseUrl(), networkId)
	var datamodel = QoSRules{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return a Quality Of Service Rule
func GetQoSRule(networkId, qosRuleId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/qosRules/%s",
		api.BaseUrl(), networkId, qosRuleId)
	var datamodel = QoSRule{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
