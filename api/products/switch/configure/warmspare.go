package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type WarmSpareConfiguration struct {
	Enabled       bool   `json:"enabled"`
	PrimarySerial string `json:"primarySerial"`
	SpareSerial   string `json:"spareSerial"`
}

// Return Warm Spare Configuration For A Switch
func GetWarmSpare(serial string) (WarmSpareConfiguration, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/warmSpare",
		api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = WarmSpareConfiguration{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
