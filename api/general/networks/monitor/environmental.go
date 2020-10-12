package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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
func GetEnvironmentalEvents(networkId string) (EnvironmentalEvents, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/environmental/events", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = EnvironmentalEvents{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
