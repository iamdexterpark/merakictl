package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type ContentFiltering struct {
	Categories []interface{} `json:"categories"`
}

// List all available content filtering categories for an MX network
func GetContentFiltering(networkId string ) (ContentFiltering, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/contentFiltering/categories", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = ContentFiltering{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

