package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type StaticRoutes []struct {
	StaticRoute
}
type StaticRoute struct {
	ID                 string `json:"id"`
	NetworkID          string `json:"networkId"`
	Enabled            bool   `json:"enabled"`
	Name               string `json:"name"`
	Subnet             string `json:"subnet"`
	GatewayIP          string `json:"gatewayIp"`
	FixedIPAssignments struct {
		Two23344556677 struct {
			IP   string `json:"ip"`
			Name string `json:"name"`
		} `json:"22:33:44:55:66:77"`
	} `json:"fixedIpAssignments"`
	ReservedIPRanges []struct {
		Start   string `json:"start"`
		End     string `json:"end"`
		Comment string `json:"comment"`
	} `json:"reservedIpRanges"`
}

// List the static routes for an MX or teleworker network
func GetStaticRoutes(networkId string ) (StaticRoutes, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/staticRoutes", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = StaticRoutes{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


// Return a static route for an MX or teleworker network
func GetStaticRoute(networkId, staticRouteId string ) (StaticRoutes, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/staticRoutes/%s", api.BaseUrl(), networkId, staticRouteId)
	var payload io.ReadSeeker
	var results = StaticRoutes{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}