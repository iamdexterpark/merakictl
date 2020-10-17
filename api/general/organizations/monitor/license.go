package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type License struct {
	Status               string `json:"status"`
	ExpirationDate       string `json:"expirationDate"`
	LicensedDeviceCounts struct {
		MS int `json:"MS"`
	} `json:"licensedDeviceCounts"`
}

// Return an overview of the license state for an organization
func GetLicenseOverview(organizationId string ) (License, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/licenses/overview", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker
	var results = License{}

	session := api.Session(baseurl, "GET", payload)

	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
