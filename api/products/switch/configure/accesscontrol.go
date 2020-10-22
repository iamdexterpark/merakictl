package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type AccessControlLists struct {
	Rules []struct {
		Comment   string `json:"comment"`
		Policy    string `json:"policy"`
		IPVersion string `json:"ipVersion"`
		Protocol  string `json:"protocol"`
		SrcCidr   string `json:"srcCidr"`
		SrcPort   string `json:"srcPort"`
		DstCidr   string `json:"dstCidr"`
		DstPort   int    `json:"dstPort"`
		Vlan      int    `json:"vlan"`
	} `json:"rules"`
}

// Return The Access Control Lists For A MS Network
func GetAccessControlLists(networkId string) (AccessControlLists, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/accessControlLists",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = AccessControlLists{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
