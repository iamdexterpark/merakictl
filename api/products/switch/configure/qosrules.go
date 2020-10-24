package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type QoSRules struct {
	RuleIds []string `json:"ruleIds"`
}

// Return the quality of service rule IDs by order in which they will be processed by the switch
func GetQoSRules(networkId string) (QoSRules, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/qosRules/order",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = QoSRules{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}