package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type BluetoothClients []struct {
	BluetoothClient
}
type BluetoothClient struct {
	ID              string   `json:"id"`
	Mac             string   `json:"mac"`
	NetworkID       string   `json:"networkId"`
	Name            string   `json:"name"`
	DeviceName      string   `json:"deviceName"`
	Manufacturer    string   `json:"manufacturer"`
	LastSeen        int      `json:"lastSeen"`
	SeenByDeviceMac string   `json:"seenByDeviceMac"`
	InSightAlert    bool     `json:"inSightAlert"`
	OutOfSightAlert bool     `json:"outOfSightAlert"`
	Tags            []string `json:"tags"`
}

// List The Bluetooth Clients Seen By APs In This Network
func GetBluetoothClients(networkId string) (BluetoothClients, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/bluetoothClients", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = BluetoothClients{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

// List A Bluetooth Client Seen By APs In This Network
func GetBluetoothClient (networkId, bluetoothClientId string) (BluetoothClient, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/bluetoothClients/%s", api.BaseUrl(), networkId, bluetoothClientId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = BluetoothClient{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

