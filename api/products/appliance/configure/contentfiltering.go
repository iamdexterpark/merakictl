package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type ContentFilteringCategories struct {
	Categories []interface{} `json:"categories"`
}

// List all available content filtering categories for an MX network
func GetContentFilteringCategories(networkId string ) (ContentFilteringCategories, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/contentFiltering/categories", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = ContentFilteringCategories{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type ContentFilteringSettings struct {
	AllowedURLPatterns   []string `json:"allowedUrlPatterns"`
	BlockedURLPatterns   []string `json:"blockedUrlPatterns"`
	BlockedURLCategories []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"blockedUrlCategories"`
	URLCategoryListSize string `json:"urlCategoryListSize"`
}

// Return the content filtering settings for an MX network
func GetContentFilteringSettings(networkId string ) (ContentFilteringSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/contentFiltering", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	var results = ContentFilteringSettings{}

	session := api.Session(baseurl, "GET", payload)
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

