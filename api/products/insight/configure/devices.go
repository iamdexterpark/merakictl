package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
	"time"
)

type DeviceCerts []struct {
	Name           string    `json:"name"`
	NotValidAfter  time.Time `json:"notValidAfter"`
	NotValidBefore time.Time `json:"notValidBefore"`
	CertPem        string    `json:"certPem"`
	DeviceID       string    `json:"deviceId"`
	Issuer         string    `json:"issuer"`
	Subject        string    `json:"subject"`
	ID             string    `json:"id"`
}

// List the certs on a device
func GetDeviceCerts(networkId, deviceId string) (DeviceCerts, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/certs",
		api.BaseUrl(), networkId, deviceId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = DeviceCerts{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

