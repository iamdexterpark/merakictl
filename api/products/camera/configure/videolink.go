package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type VideoLink struct {
	URL string `json:"url"`
}

// Returns video link to the specified camera. If a timestamp is supplied, it links to that timestamp
func GetVideoLink(serial, timestamp string) (VideoLink, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/videoLink", api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("timestamp", timestamp)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = VideoLink{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}