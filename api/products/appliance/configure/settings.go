package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type ApplianceSettings struct {
	ClientTrackingMethod string `json:"clientTrackingMethod"`
}

// Return the appliance settings for a network
func GetApplianceSettings(networkId string ) (ApplianceSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/settings", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = ApplianceSettings{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}