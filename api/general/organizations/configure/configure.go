package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"github.com/ddexterpark/merakictl/user-agent"
	"io"
)

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

// Claim A List Of Devices Licenses And Or Orders Into An Organization

// Create A New Organization By Cloning The Addressed Organization

