package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"github.com/ddexterpark/merakictl/user-agent"
	"log"
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
func GetOrganizations() []api.Results {
	baseurl := fmt.Sprintf("%s/organizations", api.BaseUrl())

	var datamodel = Organizations{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


// CreateOrganization - Create a new organization
func CreateOrganization( name string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations", api.BaseUrl())
	data := Organization{
		Name: name,
	}
	payload := user_agent.MarshalJSON(data)

	var datamodel = Organization{}
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


// GetOrganization - Return a specific organization
func GetOrganization(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s", api.BaseUrl(), organizationId)

	var datamodel = Organization{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


// UpdateOrganization - Update an organization
func UpdateOrganization(organizationId, name string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s", api.BaseUrl(), organizationId)
	data := Organization{
		Name: name,
	}

	payload := user_agent.MarshalJSON(data)

	var datamodel = Organization{}
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


// DeleteOrganization - Delete an organization
func DeleteOrganization(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s", api.BaseUrl(), organizationId)
	var datamodel = Organization{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Claim A List Of Devices Licenses And Or Orders Into An Organization

// Create A New Organization By Cloning The Addressed Organization

