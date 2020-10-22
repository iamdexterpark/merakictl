package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
	"time"
)

type ClientCellularData []struct {
	Received int       `json:"received"`
	Sent     int       `json:"sent"`
	Ts       time.Time `json:"ts"`
}

// Return the client's daily cellular data usage history
func GetClientCellularData(networkId, deviceId string) (ClientCellularData, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/cellularUsageHistory",
		api.BaseUrl(), networkId, deviceId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = ClientCellularData{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


// Returns historical connectivity data (whether a device is regularly checking in to Dashboard).

// Return historical records of various Systems Manager network connection details for desktop devices.

// Return historical records of commands sent to Systems Manager devices

// Return historical records of various Systems Manager client metrics for desktop devices.


