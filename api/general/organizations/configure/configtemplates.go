package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type ConfigurationTemplates []struct {
	ConfigurationTemplate
}
type ConfigurationTemplate struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	ProductTypes []string `json:"productTypes"`
	TimeZone     string   `json:"timeZone"`
}

// List The Configuration Templates For This Organization
func GetConfigurationTemplates(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/configTemplates", api.BaseUrl(),
		organizationId)

	var datamodel = ConfigurationTemplates{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return a Configuration Template For This Organization
func GetConfigurationTemplate(organizationId, configTemplateId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/configTemplates/%s", api.BaseUrl(),
		organizationId, configTemplateId)

	var datamodel = ConfigurationTemplate{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}