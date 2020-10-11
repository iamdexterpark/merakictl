package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type SyslogServers struct {
	Servers []struct {
		Host  string   `json:"host"`
		Port  int      `json:"port"`
		Roles []string `json:"roles"`
	} `json:"servers"`
}


// List The Syslog Servers For A Network
func GetSyslogServers(networkId string) (SyslogServers, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/syslogServers", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = SyslogServers{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
