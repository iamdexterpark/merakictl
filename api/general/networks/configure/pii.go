package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type PiiKeys struct {
	N1234 struct {
		Macs          []string `json:"macs"`
		Emails        []string `json:"emails"`
		Usernames     []string `json:"usernames"`
		Serials       []string `json:"serials"`
		Imeis         []string `json:"imeis"`
		BluetoothMacs []string `json:"bluetoothMacs"`
	} `json:"N_1234"`
}

// List the keys required to access Personally Identifiable Information (PII) for a given identifier
func GetPiiKeys(networkId string) (PiiKeys, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/pii/piiKeys", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = PiiKeys{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

