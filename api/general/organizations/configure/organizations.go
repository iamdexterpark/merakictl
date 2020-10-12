package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"github.com/ddexterpark/merakictl/user-agent"
	"io"
)

/*
MERAKI API CALLS

GET
List the Organizations
List the Networks in an Organization
List the Devices in an Organization

*/

// Organization Data Model
type Organization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Organizations Data Model
type Organizations []struct {
	Organization
}

// GetOrganizations - List the organizations that the user has privileges on
func GetOrganizations() (Organizations, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations", api.BaseUrl())
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = Organizations{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// GetOrganization - Return a specific organization
func GetOrganization(organizationId string) (Organization, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = Organization{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// CreateOrganization - Create a new organization
func CreateOrganization( name string) (Organization, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations", api.BaseUrl())
	data := Organization{
		Name: name,
	}
	payload := user_agent.MarshalJSON(data)
	session := api.Session(baseurl, "POST", payload)

	var results = Organization{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// UpdateOrganization - Update an organization
func UpdateOrganization(organizationId, name string) (Organization, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s", api.BaseUrl(), organizationId)
	data := Organization{
		Name: name,
	}

	payload := user_agent.MarshalJSON(data)
	session := api.Session(baseurl, "PUT", payload)

	var results = Organization{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


// DeleteOrganization - Delete an organization
func DeleteOrganization(organizationId string) (Organization, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s", api.BaseUrl(), organizationId)

	var payload io.ReadSeeker
	session := api.Session(baseurl, "DELETE", payload)

	var results = Organization{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// GetOrganizationNetworks - List the networks that the user has privileges on in an organization
func GetOrganizationNetworks(organizationId, configTemplateId, tagsFilterType,
	startingAfter, endingBefore string, tags, perPage string ) (Organizations, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/networks", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker

	var results = Organizations{}

	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("configTemplateId", configTemplateId)
	parameters.Add("tags", tags)
	parameters.Add("tagsFilterType",tagsFilterType)
	parameters.Add("perPage",perPage)
	parameters.Add("startingAfter",startingAfter)
	parameters.Add("endingBefore", endingBefore)
	session.Request.URL.RawQuery = parameters.Encode()


	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// Devices Data Model
type Devices []struct {
	Name      string `json:"name"`
	Serial    string `json:"serial"`
	Mac       string `json:"mac"`
	Status    string `json:"status"`
	LanIP     string `json:"lanIp"`
	PublicIP  string `json:"publicIp"`
	NetworkID string `json:"networkId"`
}

// GetOrganizationDevices -
func GetOrganizationDevices(organizationId string) (Devices, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/devices", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = Devices{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}