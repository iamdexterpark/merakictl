package user_agent

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRetryHeader(t *testing.T) {
	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(429)
			if r.URL.Path == "/test/" {
				w.Write([]byte(`[]`))
			}
		}),
	)
	defer ts.Close()

	OrgClient := MerakiClient()
	Orgcall, err := OrgClient.Get(ts.URL+"/test/", "v1", "1234567890")
	if err != nil {
		log.Fatal(err)
	}
	checkRetryHeader, checkErr := RetryPolicy(Orgcall, err)
	if checkRetryHeader == true {
		t.Log()
	} else {
		t.Errorf("[ERR] request failed:  %v %v", checkErr, err)
	}
}

func TestRetryPolicyTrue(t *testing.T) {

	OrgClient := MerakiClient()
	remain := OrgClient.MaxRetry - 3
	if remain == 0 {
		t.Log()
	} else {
		t.Errorf("[ERR] Retries Not Expired:  %v", remain)
	}
}

func TestRetryPolicyFalse(t *testing.T) {

	OrgClient := MerakiClient()
	remain := OrgClient.MaxRetry - 2
	if remain == 1 {
		t.Log()
	} else {
		t.Errorf("[ERR] Retries Policy Issue, retries left:  %v", remain)
	}
}

func TestBackoffPolicyTrue(t *testing.T) {
	// w.Header().Add("Retry-After", "1")
	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Add("Content-Type", "application/json")
			w.Header().Add("Retry-After", "2")
			w.WriteHeader(200)
			if r.URL.Path == "/test/" {
				w.Write([]byte(`[]`))
			}
		}),
	)
	defer ts.Close()

	OrgClient := MerakiClient()
	Orgcall, err := OrgClient.Get(ts.URL+"/test/", "v1", "1234567890")
	if err != nil {
		log.Fatal(err)
	}

	before := time.Now()
	wait := OrgClient.Backoff(Orgcall)
	time.Sleep(wait * time.Second)
	after := time.Now()

	if after.After(before.Add(time.Second*1)) && after.Before(before.Add(time.Second*3)) {
		t.Log()
	} else {
		println(before.String())
		println(after.String())
		println(wait)
		t.Errorf("[ERR] request failed:  %v", err)
	}
}

func TestBackoffPolicyFalse(t *testing.T) {
	// w.Header().Add("Retry-After", "1")
	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Add("Content-Type", "application/json")
			w.Header().Add("Retry-After", "3")
			w.WriteHeader(200)
			if r.URL.Path == "/test/" {
				w.Write([]byte(`[]`))
			}
		}),
	)
	defer ts.Close()

	OrgClient := MerakiClient()
	Orgcall, err := OrgClient.Get(ts.URL+"/test/", "v1", "1234567890")
	if err != nil {
		log.Fatal(err)
	}

	before := time.Now()
	wait := OrgClient.Backoff(Orgcall)
	time.Sleep(wait * time.Second)
	after := time.Now()

	if after.After(before.Add(time.Second*1)) && after.Before(before.Add(time.Second*2)) {
		println(before.String())
		println(after.String())
		println(wait)
		t.Errorf("[ERR] request failed:  %v", err)
	} else {
		t.Log()

	}
}

func TestSunSetPolicyTrue(t *testing.T) {

	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Add("Content-Type", "application/json")
			w.Header().Add("Sunset", "2020-02-05T23:59:59Z")
			w.Header().Add("Deprecation", "3020-08-05T23:59:59Z")
			w.WriteHeader(200)
			if r.URL.Path == "/test/" {
				w.Write([]byte(`[]`))
			}
		}),
	)
	defer ts.Close()

	OrgClient := MerakiClient()
	Orgcall, err := OrgClient.Get(ts.URL+"/test/", "v0", "1234567890")
	if err != nil {
		log.Fatal(err)
	}
	checkSunSet, checkErr := OrgClient.CheckForSunSet(Orgcall, err)
	if checkSunSet == true {
		t.Log()
		fmt.Println("Ignore SunSet Warning thrown by successful unit-test ^")
	} else {
		t.Errorf("[ERR] request failed:  %v %v", checkErr, err)
	}
}

func TestSunSetPolicyFalse(t *testing.T) {
	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Add("Content-Type", "application/json")
			w.Header().Add("Sunset", "3022-02-05T23:59:59Z")
			w.Header().Add("Deprecation", "3020-08-05T23:59:59Z")
			w.WriteHeader(200)
			if r.URL.Path == "/test/" {
				w.Write([]byte(`[]`))
			}
		}),
	)
	defer ts.Close()

	OrgClient := MerakiClient()
	Orgcall, err := OrgClient.Get(ts.URL+"/test/", "v0", "1234567890")
	if err != nil {
		log.Fatal(err)
	}
	checkSunSet, checkErr := OrgClient.CheckForSunSet(Orgcall, err)
	if checkSunSet == false {
		t.Log()
	} else {
		t.Errorf("[ERR] request failed:  %v %v", checkErr, err)
	}

}

func TestDeprecationPolicyTrue(t *testing.T) {

	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Add("Content-Type", "application/json")
			w.Header().Add("Deprecation", "2020-08-05T23:59:59Z")
			w.Header().Add("Sunset", "3022-02-05T23:59:59Z")
			w.WriteHeader(200)
			if r.URL.Path == "/test/" {
				w.Write([]byte(`[]`))
			}
		}),
	)
	defer ts.Close()

	OrgClient := MerakiClient()
	Orgcall, err := OrgClient.Get(ts.URL+"/test/", "v0", "1234567890")
	if err != nil {
		log.Fatal(err)
	}
	checkDeprecation, checkErr := OrgClient.CheckForDeprecation(Orgcall, err)
	if checkDeprecation == true {
		t.Log()
		fmt.Println("Ignore Deprecation Warning thrown by successful unit-test ^")
	} else {
		t.Errorf("[ERR] request failed:  %v %v", checkErr, err)
	}
}

func TestDeprecationPolicyFalse(t *testing.T) {

	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Add("Content-Type", "application/json")
			w.Header().Add("Deprecation", "2021-08-05T23:59:59Z")
			w.Header().Add("Sunset", "3022-02-05T23:59:59Z")
			w.WriteHeader(200)
			if r.URL.Path == "/test/" {
				w.Write([]byte(`[]`))
			}
		}),
	)
	defer ts.Close()

	OrgClient := MerakiClient()
	Orgcall, err := OrgClient.Get(ts.URL+"/test/", "v0", "1234567890")
	if err != nil {
		log.Fatal(err)
	}
	checkDeprecation, checkErr := OrgClient.CheckForDeprecation(Orgcall, err)
	if checkDeprecation == false {
		t.Log()
	} else {
		t.Errorf("[ERR] request failed:  %v %v", checkErr, err)
	}
}
