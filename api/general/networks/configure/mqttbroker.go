package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetMqttBrokers(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/mqttBrokers", api.BaseUrl(), networkId)
	var datamodel = MQTTBrokers{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return An MQTT Broker
func GetMqttBroker(networkId, mqttBrokerId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/mqttBrokers/%s", api.BaseUrl(), networkId, mqttBrokerId)
	var datamodel = MQTTBroker{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}