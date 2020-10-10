package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)


// Data Model
type ActionBatchList []struct {
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
func GetActionBatchList(organizationId string) (ActionBatchList, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/actionBatches", api.BaseUrl(), organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var organizations = ActionBatchList{}
	user_agent.UnMarshalJSON(session.Body, &organizations)
	traceback := user_agent.TraceBack(session)
	return organizations, traceback
}


// Return An Action Batch