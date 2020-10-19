package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type APNSCertificate struct {
	Certificate string `json:"certificate"`
}

// Get the organization's APNS certificate
func GetAPNSCertificate(organizationId string ) (APNSCertificate, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/sm/apnsCert", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker
	var results = APNSCertificate{}

	session := api.Session(baseurl, "GET", payload)

	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

