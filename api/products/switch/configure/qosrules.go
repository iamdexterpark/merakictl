package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type QoSRuleIds struct {
	RuleIds []string `json:"ruleIds"`
}

// Return the quality of service rule IDs by order in which they will be processed by the switch
func GetQoSRuleIds(networkId string) (QoSRuleIds, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/qosRules/order",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = QoSRuleIds{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
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


// List Quality Of Service Rules
func GetQoSRules(networkId string) (QoSRules, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/qosRules",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = QoSRules{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


// Return a Quality Of Service Rule
func GetQoSRule(networkId, qosRuleId string) (QoSRule, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/qosRules/%s",
		api.BaseUrl(), networkId, qosRuleId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = QoSRule{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}