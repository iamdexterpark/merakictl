package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type MonitoredMediaServers []struct {
	MonitoredMediaServer
}
type MonitoredMediaServer struct {
	ID                          string `json:"id"`
	Name                        string `json:"name"`
	Address                     string `json:"address"`
	BestEffortMonitoringEnabled bool   `json:"bestEffortMonitoringEnabled"`
}

// List The Monitored Media Servers For This Organization
func GetMonitoredMediaServers(organizationId string ) (MonitoredMediaServers, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/insight/monitoredMediaServers", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker
	var results = MonitoredMediaServers{}

	session := api.Session(baseurl, "GET", payload)

	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// Return A Monitored Media Server For This Organization
func GetMonitoredMediaServer(organizationId, monitoredMediaServerId string ) (MonitoredMediaServer, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/insight/monitoredMediaServers/%s", api.BaseUrl(),
		organizationId, monitoredMediaServerId)
	var payload io.ReadSeeker
	var results = MonitoredMediaServer{}

	session := api.Session(baseurl, "GET", payload)

	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}