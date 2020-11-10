package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type CameraRecordingSchedules []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Returns a list of all camera recording schedules
func GetCameraRecordingSchedules(serial string) []api.Results{
	baseurl := fmt.Sprintf("%s/devices/%s/camera/schedules", api.BaseUrl(), serial)
	var datamodel = CameraRecordingSchedules{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
