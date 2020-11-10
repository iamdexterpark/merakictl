package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type VideoLink struct {
	URL string `json:"url"`
}

// Returns video link to the specified camera. If a timestamp is supplied, it links to that timestamp
func GetVideoLink(serial, timestamp string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/videoLink", api.BaseUrl(), serial)
	var datamodel = VideoLink{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"timestamp": timestamp}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}