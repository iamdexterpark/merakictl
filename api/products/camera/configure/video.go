package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type VideoSettings struct {
	ExternalRtspEnabled bool   `json:"externalRtspEnabled"`
	RtspURL             string `json:"rtspUrl"`
}

// Returns video settings for the given camera
func GetVideoSettings(serial string) (VideoSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/video/settings", api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = VideoSettings{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
