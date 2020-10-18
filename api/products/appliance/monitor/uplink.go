package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
	"time"
)

type UplinkStatus []struct {
	NetworkID      string    `json:"networkId"`
	Serial         string    `json:"serial"`
	LastReportedAt time.Time `json:"lastReportedAt"`
	Uplinks        []struct {
		Interface    string `json:"interface"`
		Status       string `json:"status"`
		IP           string `json:"ip"`
		Gateway      string `json:"gateway,omitempty"`
		PublicIP     string `json:"publicIp"`
		PrimaryDNS   string `json:"primaryDns,omitempty"`
		SecondaryDNS string `json:"secondaryDns,omitempty"`
		IPAssignedBy string `json:"ipAssignedBy,omitempty"`
		Provider     string `json:"provider,omitempty"`
		Model        string `json:"model,omitempty"`
		SignalStat   struct {
			Rsrp string `json:"rsrp"`
			Rsrq string `json:"rsrq"`
		} `json:"signalStat,omitempty"`
		ConnectionType string `json:"connectionType,omitempty"`
		Apn            string `json:"apn,omitempty"`
	} `json:"uplinks"`
}

// List the uplink status of every Meraki MX and Z series appliances in the organization
func GetUplinkStatus(organizationId, perPage, startingAfter, endingBefore string) (UplinkStatus, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/appliance/uplink/statuses", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("perPage",perPage)
	parameters.Add("startingAfter",startingAfter)
	parameters.Add("endingBefore", endingBefore)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = UplinkStatus{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}