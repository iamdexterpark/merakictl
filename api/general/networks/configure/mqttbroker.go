package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type MQTTBrokers []struct {
	MQTTBroker
}

type MQTTBroker struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

// Return A List of MQTT Broker
func GetMqttBrokers(networkId string) (MQTTBrokers, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/mqttBrokers", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = MQTTBrokers{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// Return An MQTT Broker
func GetMqttBroker(networkId, mqttBrokerId string) (MQTTBroker, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/mqttBrokers/%s", api.BaseUrl(), networkId, mqttBrokerId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = MQTTBroker{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}