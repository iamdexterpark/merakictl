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
	StaticRouteID               string `json:"staticRouteId"`
	Name                        string `json:"name"`
	Subnet                      string `json:"subnet"`
	NextHopIP                   string `json:"nextHopIp"`
	AdvertiseViaOspfEnabled     bool   `json:"advertiseViaOspfEnabled"`
	PreferOverOspfRoutesEnabled bool   `json:"preferOverOspfRoutesEnabled"`
}

// List Layer 3 Static Routes For A Switch
func GetStaticRoutes(serial string) (StaticRoutes, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/routing/staticRoutes",
		api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = StaticRoutes{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback

}

// Return A Layer 3 Static Route For A Switch
func GetStaticRoute(serial, staticRouteId string) (StaticRoute, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/routing/staticRoutes/%s",
		api.BaseUrl(), serial, staticRouteId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = StaticRoute{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback

}