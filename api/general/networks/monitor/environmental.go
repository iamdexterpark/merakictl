package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
	"time"
)

type EnvironmentalEvents struct {
	PageStartAt time.Time `json:"pageStartAt"`
	PageEndAt   time.Time `json:"pageEndAt"`
	Events      []struct {
		OccurredAt    time.Time `json:"occurredAt"`
		NetworkID     string    `json:"networkId"`
		Type          string    `json:"type"`
		Description   string    `json:"description"`
		SensorSerial  string    `json:"sensorSerial"`
		GatewaySerial string    `json:"gatewaySerial"`
		EventData     struct {
			Version string `json:"version"`
		} `json:"eventData"`
	} `json:"events"`
}

// List the environmental events for the network
func GetEnvironmentalEvents(networkId, includedEventTypes, excludedEventTypes, sensorSerial,
	gatewaySerial, perPage, startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/environmental/events", api.BaseUrl(), networkId)
	var datamodel = EnvironmentalEvents{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"includedEventTypes": includedEventTypes,
		"excludedEventTypes": excludedEventTypes,
		"sensorSerial":       sensorSerial,
		"gatewaySerial":      gatewaySerial,
		"perPage":            perPage,
		"startingAfter":      startingAfter,
		"endingBefore":       endingBefore}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}
