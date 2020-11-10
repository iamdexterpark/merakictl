package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetUplinkStatus(organizationId, perPage, startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/appliance/uplink/statuses", api.BaseUrl(), organizationId)

	var datamodel = UplinkStatus{}
	// Parameters for Request URL
	var parameters = map[string]string{
		"perPage": perPage,
		"startingAfter": startingAfter,
		"endingBefore": endingBefore}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}