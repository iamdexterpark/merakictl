package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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
func GetConfigurationTemplates(organizationId string) (ConfigurationTemplates, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/configTemplates", api.BaseUrl(),
		organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = ConfigurationTemplates{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// Return a Configuration Template For This Organization
func GetConfigurationTemplate(organizationId, configTemplateId string) (ConfigurationTemplate, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/configTemplates/%s", api.BaseUrl(),
		organizationId, configTemplateId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = ConfigurationTemplate{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}