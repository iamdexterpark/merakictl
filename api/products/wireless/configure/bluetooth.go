package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type BluetoothNetworkSettings struct {
	ScanningEnabled          bool   `json:"scanningEnabled"`
	AdvertisingEnabled       bool   `json:"advertisingEnabled"`
	UUID                     string `json:"uuid"`
	MajorMinorAssignmentMode string `json:"majorMinorAssignmentMode"`
	Major                    int    `json:"major"`
}

type BluetoothDeviceSettings struct {
	UUID  string `json:"uuid"`
	Major int    `json:"major"`
	Minor int    `json:"minor"`
}

// Return The Bluetooth Settings For A Network A Href Https Documentation Meraki Com MR Bluetooth Bluetooth Low Energy BLE Bluetooth Settings A Must Be Enabled On The Network
func GetBluetoothNetworkSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/bluetooth/settings",
		api.BaseUrl(), networkId)
	var datamodel = BluetoothNetworkSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return The Bluetooth Settings For A Wireless Device
func GetBluetoothDeviceSettings(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/wireless/bluetooth/settings",
		api.BaseUrl(), serial)
	var datamodel = BluetoothDeviceSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
