package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

//SSID
type SSID struct {
	Number              int    `json:"number"`
	Name                string `json:"name"`
	Enabled             bool   `json:"enabled"`
	SplashPage          string `json:"splashPage"`
	SsidAdminAccessible bool   `json:"ssidAdminAccessible"`
	AuthMode            string `json:"authMode"`
	EncryptionMode      string `json:"encryptionMode"`
	WpaEncryptionMode   string `json:"wpaEncryptionMode"`
	RadiusServers       []struct {
		Host string `json:"host"`
		Port int    `json:"port"`
		Secret string `json:"secret"`
	} `json:"radiusServers"`
	RadiusAccountingEnabled         bool     `json:"radiusAccountingEnabled"`
	RadiusEnabled                   bool     `json:"radiusEnabled"`
	RadiusAttributeForGroupPolicies string   `json:"radiusAttributeForGroupPolicies"`
	RadiusFailoverPolicy            string   `json:"radiusFailoverPolicy"`
	RadiusLoadBalancingPolicy       string   `json:"radiusLoadBalancingPolicy"`
	IPAssignmentMode                string   `json:"ipAssignmentMode"`
	AdminSplashURL                  string   `json:"adminSplashUrl"`
	SplashTimeout                   string   `json:"splashTimeout"`
	WalledGardenEnabled             bool     `json:"walledGardenEnabled"`
	WalledGardenRanges              []string `json:"walledGardenRanges"`
	MinBitrate                      int      `json:"minBitrate"`
	BandSelection                   string   `json:"bandSelection"`
	PerClientBandwidthLimitUp       int      `json:"perClientBandwidthLimitUp"`
	PerClientBandwidthLimitDown     int      `json:"perClientBandwidthLimitDown"`
	Visible                         bool     `json:"visible"`
	AvailableOnAllAps               bool     `json:"availableOnAllAps"`
	AvailabilityTags                []string `json:"availabilityTags"`
	PerSsidBandwidthLimitUp         int      `json:"perSsidBandwidthLimitUp"`
	PerSsidBandwidthLimitDown       int      `json:"perSsidBandwidthLimitDown"`
	MandatoryDhcpEnabled            bool     `json:"mandatoryDhcpEnabled"`
}

// SSIDS Data Model
type SSIDS []struct {
	SSID
}

// List The MR SSIDs In A Network
func GetSSIDS(networkId string) (SSIDS, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var ssids SSIDS
	user_agent.UnMarshalJSON(session.Body, &ssids)
	traceback := user_agent.TraceBack(session)
	return ssids, traceback
}


// Return A Single MR SSID
func GetSSID(networkId, ssidNumber string) (SSID, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s", api.BaseUrl(), networkId,ssidNumber)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var ssid SSID
	user_agent.UnMarshalJSON(session.Body, &ssid)
	traceback := user_agent.TraceBack(session)
	return ssid, traceback
}

// Update A Single MR SSID
func UpdateSSID(networkId, ssidNumber string, data interface{}) (SSID, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s", api.BaseUrl(), networkId,ssidNumber)

	payload := user_agent.MarshalJSON(data)
	session := api.Session(baseurl, "PUT", payload)

	var ssid SSID
	user_agent.UnMarshalJSON(session.Body, &ssid)
	traceback := user_agent.TraceBack(session)

	return ssid, traceback
}
