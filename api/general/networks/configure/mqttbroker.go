package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type MQTTBrokers []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}


// Return The Meraki Auth Splash Guest RADIUS Or Client VPN User
func GetMqttBrokers(networkId string) (MerakiAuthUsers, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/mqttBrokers", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = MerakiAuthUsers{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}