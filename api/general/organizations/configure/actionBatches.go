package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetActionBatchList(organizationId, status string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/actionBatches", api.BaseUrl(), organizationId)

	// Parameters for Request URL
	var parameters = map[string]string{"status": status}

	var datamodel = ActionBatchList{}
	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return A Single Action Batch
func GetActionBatch(organizationId, actionBatchId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/actionBatches/%s", api.BaseUrl(), organizationId, actionBatchId)

	var datamodel = ActionBatch{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
