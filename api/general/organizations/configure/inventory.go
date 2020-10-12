package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
	"time"
)

type DeviceInventory []struct {
	Inventory
}
type Inventory struct {
	Mac                   string    `json:"mac"`
	Serial                string    `json:"serial"`
	Name                  string    `json:"name"`
	Model                 string    `json:"model"`
	NetworkID             string    `json:"networkId"`
	OrderNumber           string    `json:"orderNumber"`
	ClaimedAt             time.Time `json:"claimedAt"`
	LicenseExpirationDate time.Time `json:"licenseExpirationDate"`
}

// Return The Device Inventory For An Organization
func GetDeviceInventory(organizationId string) (DeviceInventory, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/inventoryDevices", api.BaseUrl(),
		organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = DeviceInventory{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// Return A Single Device From The Inventory Of An Organization
func GetSingleDeviceInventory(organizationId string) (Inventory, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/inventoryDevices", api.BaseUrl(),
		organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = Inventory{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}