package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type ContentFilteringCategories struct {
	Categories []interface{} `json:"categories"`
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

// List all available content filtering categories for an MX network
func GetContentFilteringCategories(networkId string ) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/contentFiltering/categories", api.BaseUrl(), networkId)
	var datamodel = ContentFilteringCategories{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return the content filtering settings for an MX network
func GetContentFilteringSettings(networkId string ) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/contentFiltering", api.BaseUrl(), networkId)
	var datamodel = ContentFilteringSettings{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

