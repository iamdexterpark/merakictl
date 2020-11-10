package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type WarmSpareConfiguration struct {
	Enabled       bool   `json:"enabled"`
	PrimarySerial string `json:"primarySerial"`
	SpareSerial   string `json:"spareSerial"`
}

// Return Warm Spare Configuration For A Switch
func GetWarmSpare(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/warmSpare",
		api.BaseUrl(), serial)
	var datamodel = WarmSpareConfiguration{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
