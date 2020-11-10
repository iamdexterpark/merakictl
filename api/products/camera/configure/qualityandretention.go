package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type QualityAndRetention struct {
	MotionBasedRetentionEnabled    bool        `json:"motionBasedRetentionEnabled"`
	AudioRecordingEnabled          bool        `json:"audioRecordingEnabled"`
	RestrictedBandwidthModeEnabled bool        `json:"restrictedBandwidthModeEnabled"`
	ProfileID                      interface{} `json:"profileId"`
	Quality                        string      `json:"quality"`
	MotionDetectorVersion          int         `json:"motionDetectorVersion"`
	Resolution                     int         `json:"resolution"`
}

// Returns quality and retention settings for the given camera
func GetQualityAndRetention(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/qualityAndRetention", api.BaseUrl(), serial)
	var datamodel = QualityAndRetention{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
