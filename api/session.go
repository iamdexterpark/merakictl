package api

import (
	"fmt"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
)

// Meraki API token
func Token() string {
	viper.AutomaticEnv()
	token, ok := viper.Get("MERAKI_API_TOKEN").(string)
	if !ok {
		log.Fatal("Please set your MERAKI_API_TOKEN environmental variable to use this tool. ")
	}
	return token
}

// Set Meraki API Version for BaseUrl & User-agent API Header
func APIversion() string {
	viper.AutomaticEnv()
	apiversion, ok := viper.Get("MERAKI_API_VERSION").(string)
	if !ok {
		apiversion = "v1"
	}
	return apiversion
}

// Meraki Shard Url
func BaseUrl() string {
	viper.AutomaticEnv()
	baseurl, ok := viper.Get("MERAKI_API_URL").(string)
	if !ok {
		baseurl = fmt.Sprintf("https://api-mp.meraki.com/api/%s", APIversion())
	}
	return baseurl
}

// Create a struct that allows us to pass data through a single interface
type ResponseData struct {
	Pagination []user_agent.Link
}

// request of API Call
type Request struct {
	method              string
	baseurl             string
	parameters          string
	payload             io.ReadSeeker
	DashboardAPIVersion string
}

// HTTP Protocol
type Response struct {
	Status     string      // e.g. "200 OK"
	StatusCode int         // e.g. 200
	Proto      string      // e.g. "HTTP/1.0"
	ProtoMajor int         // e.g. 1
	ProtoMinor int         // e.g. 0
	Header     http.Header // response headers
	// Body  io.ReadCloser  // response body includes entire cert chain...
}

// result of API Call
type Results struct {
	Request
	Response
	Pagination []user_agent.Link
	Payload    interface{}
}

// TestSession initializes a communication channel with the Meraki Dashboard API
func Session(baseurl string, method string, payload io.ReadSeeker,
	parameter map[string]string, datamodel interface{}) (Results, error) {
	restSession := user_agent.MerakiClient()

	// response variable
	var session *http.Response
	var err error

	// Determine HTTP Method
	switch method {

	// Get-based API Calls
	case "GET":
		session, err = restSession.Get(baseurl, APIversion(), Token())
		if err != nil {
			log.Fatal(err)
		}

	// Create-based API Calls
	case "POST":
		session, err = restSession.Post(baseurl, APIversion(), Token(), payload)
		if err != nil {
			log.Fatal(err)
		}

	// Update-based API Calls
	case "PUT":
		session, err = restSession.Put(baseurl, APIversion(), Token(), payload)
		if err != nil {
			log.Fatal(err)
		}

	// Delete-based API Calls
	case "DELETE":
		session, err = restSession.Delete(baseurl, APIversion(), Token())
		if err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatal("Unable to determine HTTP Method: ", method)
	}

	// append parameters to url
	parameters := session.Request.URL.Query()

	for param, value := range parameter {
		parameters.Add(param, value)
	}
	session.Request.URL.RawQuery = parameters.Encode()

	// marshal data into model
	user_agent.UnMarshalJSON(session.Body, &datamodel)

	dashboardRequest := Request{
		method:              method,
		baseurl:             baseurl,
		parameters:          session.Request.URL.RawQuery,
		payload:             payload,
		DashboardAPIVersion: APIversion(),
	}

	dashboardResponse := Response{
		Status:     session.Status,
		StatusCode: session.StatusCode,
		Proto:      session.Proto,
		ProtoMajor: session.ProtoMajor,
		ProtoMinor: session.ProtoMinor,
		Header:     session.Header,
		// Body:       session.Body,
	}

	metadata := Results{

		Request:    dashboardRequest,
		Response:   dashboardResponse,
		Pagination: user_agent.ParseHeader(session.Header),
		Payload:    datamodel,
	}

	return metadata, err
}

func Sessions(baseurl string, method string, payload io.ReadSeeker,
	parameters map[string]string, datamodel interface{}) ([]Results, error) {

	var sessions []Results

	session, err := Session(baseurl, method, payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	sessions = append(sessions, session)

	for _, page := range session.Pagination {
		// Update parameters
		parameters["perPage"] = page.PerPage
		parameters["startingAfter"] = page.StartingAfter
		parameters["endingBefore"] = page.EndingBefore

		if page.Rel == "next" || page.Rel == "last" {
			paginatedSession, err := Session(page.URI, "GET", payload, parameters, datamodel)
			if err != nil {
				log.Fatal(err)
			}
			sessions = append(sessions, paginatedSession)
		}

	}
	return sessions, err
}
