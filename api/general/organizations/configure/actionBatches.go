package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)


type ActionBatchList []struct {
	ActionBatch
}

type ActionBatch struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organizationId"`
	Confirmed      bool   `json:"confirmed"`
	Synchronous    bool   `json:"synchronous"`
	Status         struct {
		Completed bool          `json:"completed"`
		Failed    bool          `json:"failed"`
		Errors    []interface{} `json:"errors"`
	} `json:"status"`
	Actions []struct {
		Resource  string `json:"resource"`
		Operation string `json:"operation"`
		Body      struct {
			Enabled bool `json:"enabled"`
		} `json:"body"`
	} `json:"actions"`
}

// Return The List Of Action Batches In The Organization
func GetActionBatchList(organizationId, status string) (ActionBatchList, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/actionBatches", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	// Parameters for Request URL
	parameters := session.Request.URL.Query()
	parameters.Add("status", status)
	session.Request.URL.RawQuery = parameters.Encode()

	var results = ActionBatchList{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}


// Return A Single Action Batch
func GetActionBatch(organizationId, actionBatchId string) (ActionBatch, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/actionBatches/%s", api.BaseUrl(), organizationId, actionBatchId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = ActionBatch{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}