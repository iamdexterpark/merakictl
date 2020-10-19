package monitor

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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

// Returns live state from camera of analytics zones
func GetLiveState(serial string) (LiveState, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/analytics/live", api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = LiveState{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


type AggregateAnalytics []struct {
	StartTs      time.Time `json:"startTs"`
	EndTs        time.Time `json:"endTs"`
	ZoneID       int       `json:"zoneId"`
	Entrances    int       `json:"entrances"`
	AverageCount int       `json:"averageCount"`
}

// Returns an overview of aggregate analytics data for a timespan
func GetAggregateAnalytics(serial, t0, t1, timespan, objectType string) (AggregateAnalytics, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/analytics/overview", api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("t1", t1)
	parameters.Add("timespan",timespan)
	parameters.Add("objectType", objectType)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = AggregateAnalytics{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
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

// Returns most recent record for analytics zones
func GetAnalyticsZoneRecord(serial, objectType string) (AnalyticsZoneRecord, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/analytics/recent", api.BaseUrl(), serial)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("objectType", objectType)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = AnalyticsZoneRecord{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type AnalyticsZoneHistoricalRecords []struct {
	StartTs      time.Time `json:"startTs"`
	EndTs        time.Time `json:"endTs"`
	Entrances    int       `json:"entrances"`
	AverageCount float64   `json:"averageCount"`
}

// Return historical records for analytic zones
func GetAnalyticsZoneHistoricalRecords(serial, zoneId, t0, t1, timespan,
	resolution, objectType string) (AnalyticsZoneHistoricalRecords, interface{}) {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/analytics/zones/%s", api.BaseUrl(), serial, zoneId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("t0", t0)
	parameters.Add("t1", t1)
	parameters.Add("timespan",timespan)
	parameters.Add("resolution", resolution)
	parameters.Add("objectType", objectType)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = AnalyticsZoneHistoricalRecords{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}



// Returns All Configured Analytic Zones For This Camera

// Generate A Snapshot Of What The Camera Sees At The Specified Time And Return A Link To That Image