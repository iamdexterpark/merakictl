package user_agent

import (
	"log"
	"net/http"
)

// HTTPResponsePolicy
type HTTPResponsePolicy func(req *Request, resp *http.Response, err error) (bool, error)

// HTTPResponseHandler
func HTTPResponseHandler(req *Request, resp *http.Response, err error) (bool, error) {
	if err != nil {
		return true, err
	}
	// Check Response Code
	switch statuscode := resp.StatusCode; {
	case statuscode >= 200 && statuscode <= 299:
		return Status200(resp), nil
	case statuscode >= 300 && statuscode <= 399:
		return Status300(req, resp), nil
	case statuscode >= 400 && statuscode <= 499:
		return Status400(req, resp), nil
	case statuscode >= 500 && statuscode <= 599:
		return Status500(req, resp), nil
	}
	return false, nil
}

func Status200(resp *http.Response) bool {
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		// Nothing to do
	}
	return false
}

func Status300(req *Request, resp *http.Response) bool {
	if resp.StatusCode == 302 {
		log.Println(resp, "You were redirected from.", req.URL)
		return true
	} else if resp.StatusCode == 308 {
		log.Println(req.URL, ", Permanent Redirect. Try using the Mega-Proxy in your base URL: api-mp.meraki.com")
		return true
	} else {
		log.Println("Unknown 300 error.")
		return true
	}
}

func Status400(req *Request, resp *http.Response) bool {
	if resp.StatusCode == 400 {
		log.Println(" Bad Request. A malformed request or missing parameter.", req.URL)
		return true
	} else if resp.StatusCode == 403 {
		log.Println("Forbidden. You don't have permission to do that.", req.URL)
		return true
	} else if resp.StatusCode == 404 {
		log.Println(" Not found. No such URL, or you don't have access to the API or organization at all:", req.URL )
		return true
	} else if resp.StatusCode == 429 {
		// Nothing to do handled in the Retry Logic
		// Too many requests hit the API too quickly. We recommend an exponential backoff of your requests.
		return false
	} else {
		log.Println("Unknown 400 error.")
		return true
	}
}

func Status500(req *Request, resp *http.Response) bool {
	log.Println("External Server Error.", resp.StatusCode, resp.Body)
	return true
}
