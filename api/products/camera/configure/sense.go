package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type ObjectDetectionModel []struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type SenseSettings struct {
	SenseEnabled bool     `json:"senseEnabled"`
	MqttBrokerID string   `json:"mqttBrokerId"`
	MqttTopics   []string `json:"mqttTopics"`
}

// Returns the MV Sense object detection model list for the given camera
func GetObjectDetectionModel(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/sense/objectDetectionModels", api.BaseUrl(), serial)
	var datamodel = ObjectDetectionModel{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Returns sense settings for a given camera
func GetSenseSettings(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/sense", api.BaseUrl(), serial)
	var datamodel = SenseSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
