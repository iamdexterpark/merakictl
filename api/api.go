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

// Session initializes a communication channel with the Meraki Dashboard API
func Session(baseurl, method string, payload io.ReadSeeker) *http.Response {
	OrgClient := user_agent.MerakiClient()

	// Determine HTTP Method
	switch method {

	// Get-based API Calls
	case "GET":
		Orgcall, err := OrgClient.Get(baseurl, APIversion(), Token())
		if err != nil {
			log.Fatal(err)
		}

		// Request Additional Pages from API (Pagination Support)
		for _, l := range user_agent.ParseHeader(Orgcall.Header) {
			fmt.Println(l.URI)
			// Still work to be done here... I need a way to make
			// more than 1 get request and return all of the data
			// so that I can marshal it into a single struct
		}

		return Orgcall

	// Create-based API Calls
	case "POST":
		Orgcall, err :=  OrgClient.Post(baseurl, APIversion(), Token(), payload)
		if err != nil {
			log.Fatal(err)
		}
		return Orgcall

	// Update-based API Calls
	case "PUT":
		Orgcall, err :=  OrgClient.Put(baseurl, APIversion(), Token(), payload)
		if err != nil {
			log.Fatal(err)
		}
		//user_agent.TraceBack(Orgcall.Body, traceback)
		return Orgcall

	// Delete-based API Calls
	case "DELETE":
		Orgcall, err :=  OrgClient.Delete(baseurl, APIversion(), Token())
		if err != nil {
			log.Fatal(err)
		}
		return Orgcall
	default:
		log.Fatal("Unable to determine HTTP Method: ", method)
	}
	return nil
}