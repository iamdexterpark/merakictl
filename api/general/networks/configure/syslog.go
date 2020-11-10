package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

// SyslogServers - List The Syslog Servers For A Network
type SyslogServers struct {
	Servers []struct {
		Host  string   `json:"host"`
		Port  int      `json:"port"`
		Roles []string `json:"roles"`
	} `json:"servers"`
}

// GetSyslogServers - List The Syslog Servers For A Network
func GetSyslogServers(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/syslogServers", api.BaseUrl(), networkId)
	var datamodel = SyslogServers{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
