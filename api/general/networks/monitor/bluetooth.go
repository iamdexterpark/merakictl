package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

// BluetoothClients - List The Bluetooth Clients Seen By APs In This Network
type BluetoothClients []struct {
	BluetoothClient
}

// BluetoothClient - List A Bluetooth Client Seen By APs In This Network
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

// GetBluetoothClients - List The Bluetooth Clients Seen By APs In This Network
func GetBluetoothClients(networkId, t0, t1, timespan, perPage, startingAfter, endingBefore, includeConnectivityHistory string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/bluetoothClients", api.BaseUrl(), networkId)
	var datamodel = BluetoothClients{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":                         t0,
		"t1":                         t1,
		"timespan":                   timespan,
		"perPage":                    perPage,
		"startingAfter":              startingAfter,
		"endingBefore":               endingBefore,
		"includeConnectivityHistory": includeConnectivityHistory,
	}
	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// GetBluetoothClient - List A Bluetooth Client Seen By APs In This Network
func GetBluetoothClient(networkId, bluetoothClientId, includeConnectivityHistory, connectivityHistoryTimespan string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/bluetoothClients/%s", api.BaseUrl(), networkId, bluetoothClientId)
	var datamodel = BluetoothClient{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"includeConnectivityHistory":  includeConnectivityHistory,
		"connectivityHistoryTimespan": connectivityHistoryTimespan}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}
