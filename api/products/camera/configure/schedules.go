package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type CameraRecordingSchedules []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Returns a list of all camera recording schedules
func GetCameraRecordingSchedules(serial string) (CameraRecordingSchedules, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/schedules", api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = CameraRecordingSchedules{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}
