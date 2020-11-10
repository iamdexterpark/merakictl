package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
	"time"
)

type LiveState struct {
	Ts    time.Time `json:"ts"`
	Zones struct {
		Num0 struct {
			Person int `json:"person"`
		} `json:"0"`
	} `json:"zones"`
}

type AggregateAnalytics []struct {
	StartTs      time.Time `json:"startTs"`
	EndTs        time.Time `json:"endTs"`
	ZoneID       int       `json:"zoneId"`
	Entrances    int       `json:"entrances"`
	AverageCount int       `json:"averageCount"`
}

type AnalyticsZoneRecords []struct {
	AnalyticsZoneRecord
}

type AnalyticsZoneRecord struct {
	StartTs      time.Time `json:"startTs"`
	EndTs        time.Time `json:"endTs"`
	ZoneID       int       `json:"zoneId"`
	Entrances    int       `json:"entrances"`
	AverageCount float64   `json:"averageCount"`
}

type AnalyticsZoneHistoricalRecords []struct {
	StartTs      time.Time `json:"startTs"`
	EndTs        time.Time `json:"endTs"`
	Entrances    int       `json:"entrances"`
	AverageCount float64   `json:"averageCount"`
}

type AnalyticZones []struct {
	ID               string `json:"id"`
	Type             string `json:"type"`
	Label            string `json:"label"`
	RegionOfInterest struct {
		X0 string `json:"x0"`
		Y0 string `json:"y0"`
		X1 string `json:"x1"`
		Y1 string `json:"y1"`
	} `json:"regionOfInterest"`
}

// Returns live state from camera of analytics zones
func GetLiveState(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/analytics/live", api.BaseUrl(), serial)
	var datamodel = LiveState{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Returns an overview of aggregate analytics data for a timespan
func GetAggregateAnalytics(serial, t0, t1, timespan, objectType string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/analytics/overview", api.BaseUrl(), serial)
	var datamodel = AggregateAnalytics{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0": t0,
		"t1": t1,
		"timespan": timespan,
		"objectType": objectType}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Returns most recent record for analytics zones
func GetAnalyticsZoneRecord(serial, objectType string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/analytics/recent", api.BaseUrl(), serial)
	var datamodel = AnalyticsZoneRecord{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"objectType": objectType}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return historical records for analytic zones
func GetAnalyticsZoneHistoricalRecords(serial, zoneId, t0, t1, timespan,
	resolution, objectType string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/analytics/zones/%s", api.BaseUrl(), serial, zoneId)
	var datamodel = AnalyticsZoneHistoricalRecords{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0": t0,
		"t1": t1,
		"timespan": timespan,
		"resolution": resolution,
		"objectType": objectType}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Returns All Configured Analytic Zones For This Camera
func GetAnalyticZones(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/analytics/zones", api.BaseUrl(), serial)
	var datamodel = AnalyticZones{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}