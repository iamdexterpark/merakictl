package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
	"time"
)

type SplashLogin []struct {
	Name             string    `json:"name"`
	Login            string    `json:"login"`
	Ssid             string    `json:"ssid"`
	LoginAt          time.Time `json:"loginAt"`
	GatewayDeviceMac string    `json:"gatewayDeviceMac"`
	ClientMac        string    `json:"clientMac"`
	ClientID         string    `json:"clientId"`
	Authorization    string    `json:"authorization"`
}

// List the splash login attempts for a network
func GetSplashLoginAttempts(networkId, splashLoginAttempts, ssidNumber, loginIdentifier, timespan string) (SplashLogin, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/splashLoginAttempts", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("splashLoginAttempts", splashLoginAttempts)
	parameters.Add("ssidNumber", ssidNumber)
	parameters.Add("loginIdentifier", loginIdentifier)
	parameters.Add("timespan", timespan)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = SplashLogin{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}