package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetDeviceInventory(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/inventoryDevices", api.BaseUrl(),
		organizationId)

	var datamodel = DeviceInventory{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return A Single Device From The Inventory Of An Organization
func GetSingleDeviceInventory(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/inventoryDevices", api.BaseUrl(),
		organizationId)

	var datamodel = Inventory{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
