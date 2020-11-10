package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type VideoSettings struct {
	ExternalRtspEnabled bool   `json:"externalRtspEnabled"`
	RtspURL             string `json:"rtspUrl"`
}

// Returns video settings for the given camera
func GetVideoSettings(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/video/settings", api.BaseUrl(), serial)
	var datamodel = VideoSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
