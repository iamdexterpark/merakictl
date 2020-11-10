package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type PerformanceScore struct {
	PerfScore int `json:"perfScore"`
}

// Return the performance score for a single MX. Only primary MX devices supported.
// If no data is available, a 204 error code is returned.
func GetPerformanceScore(serial string ) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/appliance/performance", api.BaseUrl(), serial)
	var datamodel = PerformanceScore{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}