package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type FloorPlan []struct {
	FloorPlanID       string `json:"floorPlanId"`
	ImageURL          string `json:"imageUrl"`
	ImageURLExpiresAt string `json:"imageUrlExpiresAt"`
	ImageExtension    string `json:"imageExtension"`
	ImageMd5          string `json:"imageMd5"`
	Name              string `json:"name"`
	Devices           []struct {
		Name           string   `json:"name"`
		Lat            float64  `json:"lat"`
		Lng            float64  `json:"lng"`
		Serial         string   `json:"serial"`
		Mac            string   `json:"mac"`
		Model          string   `json:"model"`
		Address        string   `json:"address"`
		Notes          string   `json:"notes"`
		LanIP          string   `json:"lanIp"`
		Tags           []string `json:"tags"`
		NetworkID      string   `json:"networkId"`
		BeaconIDParams struct {
			UUID  string `json:"uuid"`
			Major int    `json:"major"`
			Minor int    `json:"minor"`
		} `json:"beaconIdParams"`
		Firmware    string `json:"firmware"`
		FloorPlanID string `json:"floorPlanId"`
	} `json:"devices"`
	Width  int `json:"width"`
	Height int `json:"height"`
	Center struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"center"`
	BottomLeftCorner struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"bottomLeftCorner"`
	BottomRightCorner struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"bottomRightCorner"`
	TopLeftCorner struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"topLeftCorner"`
	TopRightCorner struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"topRightCorner"`
}

// List The Floor Plans That Belong To Your Network
func GetFloorPlans(networkId string) (FloorPlan, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/floorPlans", api.BaseUrl(), networkId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var floorplan = FloorPlan{}
	user_agent.UnMarshalJSON(session.Body, &floorplan)
	traceback := user_agent.TraceBack(session)
	return floorplan, traceback
}

// List The Floor Plans That Belong To Your Network
func GetFloorPlan(networkId, floorPlanId string) (FloorPlan, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/floorPlans/%s", api.BaseUrl(), networkId, floorPlanId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var floorplan = FloorPlan{}
	user_agent.UnMarshalJSON(session.Body, &floorplan)
	traceback := user_agent.TraceBack(session)
	return floorplan, traceback
}