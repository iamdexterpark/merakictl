package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

// SNMP - Return The SNMP Settings For A Network
type SNMP struct {
	Access string `json:"access"`
	Users  []struct {
		Username   string `json:"username"`
		Passphrase string `json:"passphrase"`
	} `json:"users"`
}

// GetSNMP - Return The SNMP Settings For A Network
func GetSNMP(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/snmp", api.BaseUrl(), networkId)
	var datamodel = SNMP{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
