package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type QualityRetentionProfiles []struct {
	QualityRetentionProfile
}
type QualityRetentionProfile struct {
	ID                             string      `json:"id"`
	NetworkID                      string      `json:"networkId"`
	Name                           string      `json:"name"`
	RestrictedBandwidthModeEnabled bool        `json:"restrictedBandwidthModeEnabled"`
	MotionBasedRetentionEnabled    bool        `json:"motionBasedRetentionEnabled"`
	AudioRecordingEnabled          bool        `json:"audioRecordingEnabled"`
	CloudArchiveEnabled            bool        `json:"cloudArchiveEnabled"`
	MaxRetentionDays               int         `json:"maxRetentionDays"`
	ScheduleID                     interface{} `json:"scheduleId"`
	MotionDetectorVersion          int         `json:"motionDetectorVersion"`
	VideoSettings                  struct {
		MV32 struct {
			Quality    string `json:"quality"`
			Resolution string `json:"resolution"`
		} `json:"MV32"`
		MV21MV71 struct {
			Quality    string `json:"quality"`
			Resolution string `json:"resolution"`
		} `json:"MV21/MV71"`
		MV12MV22MV72 struct {
			Quality    string `json:"quality"`
			Resolution string `json:"resolution"`
		} `json:"MV12/MV22/MV72"`
		MV12WE struct {
			Quality    string `json:"quality"`
			Resolution string `json:"resolution"`
		} `json:"MV12WE"`
	} `json:"videoSettings"`
}

// List the quality retention profiles for this network
func GetQualityRetentionProfiles(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/camera/qualityRetentionProfiles", api.BaseUrl(), networkId)
	var datamodel = QualityRetentionProfiles{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Retrieve a single quality retention profile
func GetQualityRetentionProfile(networkId, qualityRetentionProfileId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/camera/qualityRetentionProfiles/%s", api.BaseUrl(),
		networkId, qualityRetentionProfileId)
	var datamodel = QualityRetentionProfile{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
