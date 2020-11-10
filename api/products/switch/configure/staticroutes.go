package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetStaticRoutes(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/routing/staticRoutes",
		api.BaseUrl(), serial)
	var datamodel = StaticRoutes{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return A Layer 3 Static Route For A Switch
func GetStaticRoute(serial, staticRouteId string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/routing/staticRoutes/%s",
		api.BaseUrl(), serial, staticRouteId)
	var datamodel = StaticRoute{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions

}
