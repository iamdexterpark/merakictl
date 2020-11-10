package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetSplashLoginAttempts(networkId, splashLoginAttempts, ssidNumber, loginIdentifier, timespan string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/splashLoginAttempts", api.BaseUrl(), networkId)
	datamodel := SplashLogin{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"splashLoginAttempts": splashLoginAttempts,
		"ssidNumber": ssidNumber,
		"loginIdentifier": loginIdentifier,
		"timespan": timespan}


	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}