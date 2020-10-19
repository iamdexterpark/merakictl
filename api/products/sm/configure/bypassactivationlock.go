package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type BypassActivationLock struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Data   struct {
		Num2090938209 struct {
			Success bool     `json:"success"`
			Errors  []string `json:"errors"`
		} `json:"2090938209"`
		Num38290139892 struct {
			Success bool `json:"success"`
		} `json:"38290139892"`
	} `json:"data"`
}

// Bypass Activation Lock Attempt Status
func GetBypassActivationLockStatus(networkId, attemptId string) (BypassActivationLock, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/bypassActivationLockAttempts/%s",
		api.BaseUrl(), networkId, attemptId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = BypassActivationLock{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}