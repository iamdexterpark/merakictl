package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type UserProfiles []struct {
	DeviceID    string `json:"deviceId"`
	ID          string `json:"id"`
	IsEncrypted bool   `json:"isEncrypted"`
	IsManaged   bool   `json:"isManaged"`
	ProfileData struct {
	} `json:"profileData"`
	ProfileIdentifier string `json:"profileIdentifier"`
	Name              string `json:"name"`
	Version           string `json:"version"`
}

// Get the profiles associated with a user
func GetUserProfiles(networkId, userId string) (UserProfiles, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/users/%s/deviceProfiles",
		api.BaseUrl(), networkId, userId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = UserProfiles{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
