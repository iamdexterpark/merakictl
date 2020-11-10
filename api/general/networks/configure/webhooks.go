package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type HTTPServers []struct {
	HTTPServer
}
type HTTPServer struct {
	ID           string `json:"id"`
	NetworkID    string `json:"networkId"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	SharedSecret string `json:"sharedSecret"`
}

type Webhook struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Status string `json:"status"`
}

// List The HTTP Servers For A Network
func GetHTTPServers(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/webhooks/httpServers", api.BaseUrl(), networkId)
	var datamodel = HTTPServers{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}

// Return An HTTP Server For A Network
func GetHTTPServer(networkId, httpServerId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/webhooks/httpServers/%s", api.BaseUrl(), networkId, httpServerId)
	var datamodel HTTPServer
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}

// Return The Status Of A Webhook Test For A Network
func GetWebhookStatus(networkId, webhookTestId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/webhooks/webhookTests/%s", api.BaseUrl(), networkId, webhookTestId)
	var datamodel = Webhook{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}