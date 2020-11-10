package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetBypassActivationLockStatus(networkId, attemptId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/bypassActivationLockAttempts/%s",
		api.BaseUrl(), networkId, attemptId)

	var datamodel = BypassActivationLock{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
