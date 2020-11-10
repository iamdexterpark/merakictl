package user_agent

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

// Client is used to make HTTP requests
type Client struct {
	HTTPClient *http.Client // Internal HTTP client
	MaxRetry   int          // Maximum number of retries

	// CheckForRetry specifies the policy for handling retries, and is called
	// after each request
	CheckForRetry CheckForRetry

	// Backoff specifies the policy for how long to wait between retries
	Backoff Backoff

	// Deprecation checks to see if we are using the most recent API version
	CheckForDeprecation CheckForDeprecation

	// Sunset checks to see if our version of the Meraki API is no longer supported
	CheckForSunSet CheckForSunSet

	// Pagination
	CheckForPagination CheckForPagination

	// HTTP Response Status Handler
	HTTPResponsePolicy HTTPResponsePolicy

	// RedirectPolicy

}

// CheckForRetry
type CheckForRetry func(resp *http.Response, err error) (bool, error)

// RetryPolicy
func RetryPolicy(resp *http.Response, err error) (bool, error) {
	if err != nil {
		return true, err
	}
	// Check Response Code
	if resp.StatusCode == 429 {
		return true, err
	}
	return false, nil
}

var (
	// Default Retry Configuration
	MaxRetryPolicy = 3
)

// Backoff Strategy
type Backoff func(resp *http.Response) time.Duration

// BackoffPolicy
func BackoffPolicy(resp *http.Response) time.Duration {
	retry := resp.Header.Get("Retry-After")
	retryAfter, err := strconv.Atoi(retry)
	if err != nil {
		println(err)
	}
	sleep := time.Duration(retryAfter)
	return sleep
}

// A "Deprecation" response header informing you that this is no longer
// the preferred version. It will remain operational and supported,
//yet we encourage you to migrate to the newest version.'
type CheckForDeprecation func(resp *http.Response, err error) (bool, error)

// DeprecationPolicy
func DeprecationPolicy(resp *http.Response, err error) (bool, error) {
	if err != nil {
		return true, err
	}
	// Check Response Code
	now := time.Now()
	deprecation := resp.Header.Get("Deprecation")
	if deprecation == "" {
		result := false
		return result, nil
	} else {
		checkdate, err := time.Parse(time.RFC3339, deprecation)
		if err != nil {
			log.Fatal(err)
		}
		result := now.After(checkdate)
		return result, nil

	}
}

// A "Sunset" response header informing you when this version is no longer
// supported and may become unresponsive.
type CheckForSunSet func(resp *http.Response, err error) (bool, error)

// SunSetPolicy
func SunSetPolicy(resp *http.Response, err error) (bool, error) {
	if err != nil {
		return true, err
	}
	// Check Response Code
	now := time.Now()
	sunset := resp.Header.Get("Sunset")
	if sunset == "" {
		result := false
		return result, nil
	} else {
		checkdate, err := time.Parse(time.RFC3339, sunset)
		if err != nil {
			log.Fatal(err)
		}
		result := now.After(checkdate)
		return result, nil

	}
}

// CheckForPagination The "Link" header is a comma-separated list of up to 4 links: first, prev, next, and last.
type CheckForPagination func(resp *http.Response, err error) (bool, error)

// PaginationPolicy
func PaginationPolicy(resp *http.Response, err error) (bool, error) {
	if err != nil {
		return true, err
	}
	// Check Response Code
	link := resp.Header.Get("Link")
	var result bool
	if link == "" {
		result = false
		return result, nil
	} else {
		result = true
		return result, nil

	}
}
