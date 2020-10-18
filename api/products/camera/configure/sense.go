package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type ObjectDetectionModel []struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

// Returns the MV Sense object detection model list for the given camera
func GetObjectDetectionModel(serial string) (ObjectDetectionModel, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/sense/objectDetectionModels", api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = ObjectDetectionModel{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}