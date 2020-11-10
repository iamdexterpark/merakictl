package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"log"
)


type L3FirewallRules struct {
	Rules []struct {
		Comment  string `json:"comment"`
		Policy   string `json:"policy"`
		Protocol string `json:"protocol"`
		DestPort int    `json:"destPort"`
		DestCidr string `json:"destCidr"`
	} `json:"rules"`
}

type L7FirewallRules struct {
	Rules []struct {
		Policy string `json:"policy"`
		Type   string `json:"type"`
		Value  struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"value"`
	} `json:"rules"`
}

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

type SSIDS []struct {
	SSID
}

// Return The L3 Firewall Rules For An SSID On An MR Network
func GetL3FirewallRules(networkId, number string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s/firewall/l3FirewallRules",
		api.BaseUrl(), networkId, number)
	var datamodel L3FirewallRules
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return The L7 Firewall Rules For An SSID On An MR Network
func GetL7FirewallRules(networkId, number string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s/firewall/l7FirewallRules",
		api.BaseUrl(), networkId, number)
	var datamodel L7FirewallRules
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List The MR SSIDs In A Network
func GetSSIDS(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids", api.BaseUrl(), networkId)
	var datamodel SSIDS
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


// Return A Single MR SSID
func GetSSID(networkId, ssidNumber string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s", api.BaseUrl(), networkId,ssidNumber)
	var datamodel SSID
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update A Single MR SSID
func UpdateSSID(networkId, ssidNumber string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s", api.BaseUrl(), networkId,ssidNumber)

	payload := user_agent.MarshalJSON(data)
	var datamodel SSID
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
