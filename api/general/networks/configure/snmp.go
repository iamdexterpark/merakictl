package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type SNMP struct {
	Access string `json:"access"`
	Users  []struct {
		Username   string `json:"username"`
		Passphrase string `json:"passphrase"`
	} `json:"users"`
}


// Return The SNMP Settings For A Network
func GetSNMP(networkId string) (SNMP, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/snmp", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = SNMP{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}