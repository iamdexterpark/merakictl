package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
	"time"
)

type Licenses []struct {
	License
}
type License struct {
	ID                        string      `json:"id"`
	LicenseType               string      `json:"licenseType"`
	LicenseKey                string      `json:"licenseKey"`
	OrderNumber               string      `json:"orderNumber"`
	DeviceSerial              string      `json:"deviceSerial"`
	NetworkID                 string      `json:"networkId"`
	State                     string      `json:"state"`
	SeatCount                 interface{} `json:"seatCount"`
	TotalDurationInDays       int         `json:"totalDurationInDays"`
	DurationInDays            int         `json:"durationInDays"`
	PermanentlyQueuedLicenses []struct {
		ID             string `json:"id"`
		LicenseType    string `json:"licenseType"`
		LicenseKey     string `json:"licenseKey"`
		OrderNumber    string `json:"orderNumber"`
		DurationInDays int    `json:"durationInDays"`
	} `json:"permanentlyQueuedLicenses"`
	ClaimDate      time.Time `json:"claimDate"`
	ActivationDate time.Time `json:"activationDate"`
	ExpirationDate time.Time `json:"expirationDate"`
}

// List The Licenses For An Organization
func GetLicenses(organizationId string) (Licenses, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/licenses", api.BaseUrl(),
		organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = Licenses{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// List A Single License For An Organization
func GetLicense(organizationId, licenseId string) (License, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/licenses/%s", api.BaseUrl(),
		organizationId, licenseId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = License{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}