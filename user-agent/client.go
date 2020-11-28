package user_agent

/*
Special thanks to Nitish Kumar: https://medium.com/@nitishkr88/http-retries-in-go-e622e51d249f
*/

import (
	"fmt"
	"github.com/hashicorp/go-cleanhttp"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// MerakiClient creates a new Client with default settings
func MerakiClient() *Client {
	return &Client{
		HTTPClient:          cleanhttp.DefaultClient(),
		MaxRetry:            MaxRetryPolicy,
		CheckForRetry:       RetryPolicy,
		Backoff:             BackoffPolicy,
		CheckForDeprecation: DeprecationPolicy,
		CheckForSunSet:      SunSetPolicy,
		CheckForPagination:  PaginationPolicy,
		HTTPResponsePolicy:  HTTPResponseHandler,
	}
}

// Request wraps the metadata needed to create HTTP requests.
type Request struct {
	body io.ReadSeeker

	// Embed an HTTP request directly.
	*http.Request
}

// NewRequest creates a new wrapped request.
func NewRequest(method, url string, body io.ReadSeeker) (*Request, error) {
	// Wrap the body in a noop ReadCloser if non-nil. This prevents the
	//reader from being closed by the HTTP client
	var rcBody io.ReadCloser
	if body != nil {
		rcBody = ioutil.NopCloser(body)
	}

	// make the request with the noop-closer for the body.
	httpReq, err := http.NewRequest(method, url, rcBody)
	if err != nil {
		return nil, err
	}
	return &Request{body, httpReq}, nil
}

// Do wraps calling an HTTP method with retries
func (c *Client) Do(req *Request) (*http.Response, error) {
	for {
		//Always rewind the request body when non-nil
		if req.body != nil {
			if _, err := req.body.Seek(0, 0); err != nil {
				return nil, fmt.Errorf("failed to seek body: %v", err)
			}
		}

		// Attempt the request
		resp, err := c.HTTPClient.Do(req.Request)

		// Check to see if the API version is Deprecated
		checkDeprecation, checkErr := c.CheckForDeprecation(resp, err)
		if err != nil {
			log.Fatalf("[ERR] %s %s request failed: %v", req.Method, req.URL, err)
		} else {
			// Not Deprecated
		}
		if checkDeprecation {
			log.Print("This is no longer the preferred API version. Please set your env var to the latest Meraki API version")
		}

		// Check to see if the API version is SunSet
		checkSunSet, checkErr := c.CheckForSunSet(resp, err)
		if err != nil {
			log.Fatalf("[ERR] %s %s request failed: %v", req.Method, req.URL, err)
		} else {
			// Not SunSet
		}
		if checkSunSet {
			log.Println("This API version is no longer supported and may become unresponsive. Please set your env var to the latest Meraki API version")
		}

		// CheckForPagination
		checkPagination, checkErr := c.CheckForPagination(resp, err)
		if err != nil {
			log.Fatalf("[ERR] %s %s request failed: %v", req.Method, req.URL, err)
		} else {
			// Not Paginated
		}
		if checkPagination {
			// Paginated
		}

		// HTTPResponseHandler
		checkHttpOK, checkHttpErr := c.HTTPResponsePolicy(req, resp, err)
		if err != nil {
			log.Fatalf("[ERR] %s %s request failed: %v", req.Method, req.URL, err)
		} else {
			// Nothing to DO
		}

		// Now decide if we should continue.
		if !checkHttpOK {
			if checkHttpErr != nil {
				err = checkHttpErr
			}
			return resp, err
		} else {

		}

		// Check if we should continue with retries.
		checkOK, checkErr := c.CheckForRetry(resp, err)

		if err != nil {
			log.Fatalf("[ERR] %s %s request failed: %v", req.Method, req.URL, err)
		} else {
			// Call this here to maintain the behavior of logging all requests, etc
			// even if the CheckForRetry signals to stop.
		}

		// Now decide if we should continue.
		if !checkOK {
			if checkErr != nil {
				err = checkErr
			}
			return resp, err
		}

		// We're going to retry, consume any response to reuse the connection.
		if err == nil {
			c.drainBody(resp.Body)

		}

		remain := c.MaxRetry - 1
		if remain == 0 {
			break
		}

		// wait is the Meraki retry-after header
		wait := c.Backoff(resp)
		time.Sleep(wait * time.Second)
	}

	// Return an error if we fail out of the retry loop
	return nil, fmt.Errorf("%s %s giving up after %d attempts", req.Method, req.URL, c.MaxRetry+1)
}

// Try to read the response body so we can reuse this connection.
func (c *Client) drainBody(body io.ReadCloser) {
	defer body.Close()
	var respReadLimit int64
	_, err := io.Copy(ioutil.Discard, io.LimitReader(body, respReadLimit))
	if err != nil {
		log.Fatalf("[ERR] error reading response body: %v", err)
	}
}

// Get is a convenience helper for doing simple GET requests.
func (c *Client) Get(url, apiversion, apitoken string) (*http.Response, error) {
	req, err := NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	ApplyMerakiHeaders(req, apiversion, apitoken)
	return c.Do(req)
}

// Post is a convenience helper for doing simple Post requests.
func (c *Client) Post(url, apiversion, apitoken string, body io.ReadSeeker) (*http.Response, error) {
	req, err := NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	ApplyMerakiHeaders(req, apiversion, apitoken)
	return c.Do(req)
}

// Put is a convenience helper for doing simple Put requests.
func (c *Client) Put(url, apiversion, apitoken string, body io.ReadSeeker) (*http.Response, error) {
	req, err := NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}
	ApplyMerakiHeaders(req, apiversion, apitoken)
	return c.Do(req)
}

// Delete is a convenience helper for doing simple Delete requests.
func (c *Client) Delete(url, apiversion, apitoken string) (*http.Response, error) {
	req, err := NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	ApplyMerakiHeaders(req, apiversion, apitoken)
	return c.Do(req)
}
