package user_agent

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Create HTTP Request Headers
func ApplyMerakiHeaders(req *Request, merakiapiversion, apitoken string) {
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "merakictl")
	req.Header.Add("X-Cisco-Meraki-API-Key", apitoken)
	if merakiapiversion == "v0" {
		// Nothing to do
	} else if merakiapiversion == "v1" {
		bearer := fmt.Sprintf("Bearer %s", apitoken)
		req.Header.Add("Authorization", bearer)
	} else {
		message := fmt.Sprintf("Unable to determine Meraki API Version: %s", merakiapiversion)
		panic(message)
	}
	return
}

// Try to read the response body so we can reuse this connection.
func UnMarshalJSON(body io.ReadCloser, respFormat interface{}) interface{} {
	if body != nil {
		defer body.Close()
	}

	resp, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}

	jsonErr := json.Unmarshal(resp, &respFormat)
	if jsonErr != nil {
	}

	return respFormat
}

func MarshalJSON(data interface{}) *strings.Reader {
	var jsonData []byte
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	inputs := fmt.Sprintf(string(jsonData))

	payload := strings.NewReader(inputs)
	return payload
}

type HttpPacket struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0

	// response headers
	Header  http.Header   // response body
	Body    io.ReadCloser // request that was sent to obtain the response
	Request *http.Request
}

// debug http response
func TraceBack(resp *http.Response) interface{} {

	// response
	traceback := HttpPacket{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Proto:      resp.Proto,
		ProtoMajor: resp.ProtoMajor,
		ProtoMinor: resp.ProtoMinor,

		Header:  resp.Header,
		Body:    resp.Body,
		Request: resp.Request,
	}

	return traceback
}
