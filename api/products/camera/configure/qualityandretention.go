package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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
func GetQualityAndRetention(serial string) (QualityAndRetention, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/qualityAndRetention", api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = QualityAndRetention{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
