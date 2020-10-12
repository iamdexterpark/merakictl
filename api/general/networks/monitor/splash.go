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
func GetSplashLoginAttempts(networkId string) (SplashLogin, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/splashLoginAttempts", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = SplashLogin{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}