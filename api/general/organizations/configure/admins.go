package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
	"time"
)


// Data Model
type Admin []struct {
	ID                   string    `json:"id"`
	Name                 string    `json:"name"`
	Email                string    `json:"email"`
	OrgAccess            string    `json:"orgAccess"`
	AccountStatus        string    `json:"accountStatus"`
	TwoFactorAuthEnabled bool      `json:"twoFactorAuthEnabled"`
	HasAPIKey            bool      `json:"hasApiKey"`
	LastActive           time.Time `json:"lastActive"`
	Tags                 []struct {
		Tag    string `json:"tag"`
		Access string `json:"access"`
	} `json:"tags"`
	Networks []struct {
		ID     string `json:"id"`
		Access string `json:"access"`
	} `json:"networks"`
	AuthenticationMethod string `json:"authenticationMethod"`
}

// List The Dashboard Administrators In This Organization
func GetAdmins(organizationId string) (Admin, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/admins", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker

	session := api.Session(baseurl, "GET", payload)

	var admins = Admin{}
	user_agent.UnMarshalJSON(session.Body, &admins)
	traceback := user_agent.TraceBack(session)
	return admins, traceback
}