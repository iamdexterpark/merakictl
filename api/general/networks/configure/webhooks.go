package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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

// List The HTTP Servers For A Network
func GetHTTPServers(networkId string) (HTTPServers, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/webhooks/httpServers", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = HTTPServers{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// Return An HTTP Server For A Network
func GetHTTPServer(networkId, httpServerId string) (HTTPServer, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/webhooks/httpServers/%s", api.BaseUrl(), networkId, httpServerId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = HTTPServer{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type Webhook struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Status string `json:"status"`
}

// Return The Status Of A Webhook Test For A Network
func GetWebhookStatus(networkId, webhookTestId string) (Webhook, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/webhooks/webhookTests/%s", api.BaseUrl(), networkId, webhookTestId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = Webhook{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}