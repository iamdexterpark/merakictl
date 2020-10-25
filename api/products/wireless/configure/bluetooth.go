package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type BluetoothNetworkSettings struct {
	ScanningEnabled          bool   `json:"scanningEnabled"`
	AdvertisingEnabled       bool   `json:"advertisingEnabled"`
	UUID                     string `json:"uuid"`
	MajorMinorAssignmentMode string `json:"majorMinorAssignmentMode"`
	Major                    int    `json:"major"`
}

// Return The Bluetooth Settings For A Network A Href Https Documentation Meraki Com MR Bluetooth Bluetooth Low Energy BLE Bluetooth Settings A Must Be Enabled On The Network
func GetBluetoothNetworkSettings(networkId string) (BluetoothNetworkSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/bluetooth/settings",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = BluetoothNetworkSettings{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type BluetoothDeviceSettings struct {
	UUID  string `json:"uuid"`
	Major int    `json:"major"`
	Minor int    `json:"minor"`
}
// Return The Bluetooth Settings For A Wireless Device
func GetBluetoothDeviceSettings(serial string) (BluetoothDeviceSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/wireless/bluetooth/settings",
		api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = BluetoothDeviceSettings{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}