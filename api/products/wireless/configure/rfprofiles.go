package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type RFProfiles []struct {
	RFProfile
}
type RFProfile struct {
	ID                     string `json:"id"`
	NetworkID              string `json:"networkId"`
	Name                   string `json:"name"`
	ClientBalancingEnabled bool   `json:"clientBalancingEnabled"`
	MinBitrateType         string `json:"minBitrateType"`
	BandSelectionType      string `json:"bandSelectionType"`
	ApSelectionSettings    struct {
		BandOperationMode   string `json:"bandOperationMode"`
		BandSteeringEnabled bool   `json:"bandSteeringEnabled"`
	} `json:"apSelectionSettings"`
	TwoFourGhzSettings struct {
		MaxPower          int         `json:"maxPower"`
		MinPower          int         `json:"minPower"`
		MinBitrate        int         `json:"minBitrate"`
		Rxsop             interface{} `json:"rxsop"`
		ValidAutoChannels []int       `json:"validAutoChannels"`
		AxEnabled         bool        `json:"axEnabled"`
	} `json:"twoFourGhzSettings"`
	FiveGhzSettings struct {
		MaxPower          int         `json:"maxPower"`
		MinPower          int         `json:"minPower"`
		MinBitrate        int         `json:"minBitrate"`
		Rxsop             interface{} `json:"rxsop"`
		ValidAutoChannels []int       `json:"validAutoChannels"`
		ChannelWidth      string      `json:"channelWidth"`
	} `json:"fiveGhzSettings"`
}

// List The Non Basic RF Profiles For This Network
func GetRFProfiles(networkId, includeTemplateProfiles string) (RFProfiles, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/rfProfiles",
		api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("includeTemplateProfiles", includeTemplateProfiles)

	var results = RFProfiles{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


// Return A RF Profile
func GetRFProfile(networkId, rfProfileId string) (RFProfile, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/rfProfiles/%s",
		api.BaseUrl(), networkId, rfProfileId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = RFProfile{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}