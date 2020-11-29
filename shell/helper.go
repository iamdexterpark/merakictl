package shell

import (
	"encoding/json"
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"github.com/ddexterpark/merakictl/api/general/organizations/configure"
	"github.com/kr/pretty"
	"github.com/spf13/pflag"
	"os"
	"regexp"
)

type Match struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Organization string `json:"organizationId"`
	Network string `json:"networkId"`
	Serial string `json:"serial"`
}

func stripDuplicateMatches(matchList []Match) (matchResults []Match) {

	// strip duplicate entries
	encountered := map[Match]bool{}

	// Create a map of all unique elements.
	for v := range matchList  {
		encountered[matchList[v]] = true
	}

	// Place all keys from the map into a slice.

	for key, _ := range encountered {
	matchResults = append(matchResults, key)
	}

	return matchResults
}

// Resolve Meraki Dashboard Organization name to id
func ResolveMatches(regExString string, Matches []Match ) (exactMatch Match, greedyMatch []Match) {

	greedyResults := []Match{}

	exactResults := []Match{}

	// eval orgs in the org list
	for _, data := range Matches {

		// match data
		match := Match{
			Name: data.Name,
			ID: data.ID,
			Organization: data.Organization,
			Network: data.Network,
			Serial: data.Serial,
		}

		// build regex
		matchRegex := regexp.MustCompile(regExString)
		greedy := matchRegex.MatchString(data.Name)


		// NEED TO ADDRESS MULTIBLE MATCH SENARIO

		if data.Name == regExString {

			exactResults = append(exactResults, match)

		} else if greedy == true {

			greedyResults = append(greedyResults, match) }
	}

	// strip duplicates from lists
	exactResults = stripDuplicateMatches(exactResults)
	greedyMatch = stripDuplicateMatches(greedyResults)

    if len(exactResults) == 1 {
		for _, data := range exactResults {
			exactMatch = Match{
				Name: data.Name,
				ID:   data.ID,
			}
		}
	}



	return exactMatch, greedyMatch
}


func resolveOrgId (orgName string) (orgId string, orgList []string) {
	// ExactMatch array
	var exactMatch Match

	//data model
	organizations := []Match{}

	// API Call to get list of organizations
	orgCall := configure.GetOrganizations()
	// iterate through api call(s)
	for _, metadata := range orgCall {
		response, _ := json.Marshal(metadata.Payload)
		json.Unmarshal(response, &organizations)
	}

	if orgName != "" {
			// iterate through list of orgs and append to exact/greedy lists
			exactOrg, greedyOrg := ResolveMatches(orgName, organizations)

			if exactOrg.ID != "" {
				mResults, _ := json.Marshal(exactOrg)
				json.Unmarshal(mResults, &exactMatch)
				orgId = exactMatch.ID
			} else {
				pretty.Println(greedyOrg)
				pretty.Println("No exact match for -o flag. Please use an explicit orgId or Name^")
				os.Exit(1)
			}


	}

	// compile org list as helper for other functions
	for _, org := range organizations {
		orgList = append(orgList, org.ID)
	}

	return orgId, orgList
}


func resolveNetworkId (netName, orgId string, orgList []string) (netId string) {
	// ExactMatch array
	//var exactMatch Match

	// networkIdList array
	var networkIdList []Match

	// greedyMatch array
	//var greedyMatches []Match

	// Determine if GetOrganizations API Call is Needed
	if len(orgList) == 0 {
		_, addOrgs := resolveOrgId("")

		for _, addOrg := range addOrgs {

			orgList = append(orgList, addOrg)
		}
	}


	// iterate through org list
	for _, org := range orgList {
		// API Call to get list of networks
		metadata := api.GetOrganizationNetworks(org, "",
			"", "", "", "", "10")

		var pagnatedData []Match

		for _, data := range metadata {
			format, _ := json.Marshal(data.Payload)
			json.Unmarshal(format, &pagnatedData)
			networkIdList = append(networkIdList, pagnatedData...)
		}
	}

		// iterate through list of orgs and append to exact/greedy lists
		exactnet, greedynet := ResolveMatches(netName, networkIdList)


		// NEED TO FIGURE OUT LOGIC HERE.....No NetId w/ exact match

		if exactnet.Name == netName {
			netId = exactnet.ID

		} else if len(orgId) != 0 {
			for _, greedyMatch := range greedynet {

				if greedyMatch.Organization == orgId {
					netId = greedyMatch.ID
				}
			}

		}

	// return error if no match
	if len(netId) == 0 {
		pretty.Println(greedynet)
		pretty.Println("No exact match for -n flag. Please use an explicit netId, Name or use -o flag in conjunction with this call ^")
		os.Exit(1)
	}

	return netId
}



func ResolveFlags(flags *pflag.FlagSet) (orgId, networkId, DeviceId string) {

	// store flags as variables
	orgName, _ := flags.GetString("organization")
	networkName, _ := flags.GetString("network")
	//deviceName, _ := flags.GetString("hostname")

	var orgList []string
	if orgName != "" {
		orgId, orgList = resolveOrgId(orgName)
	}

	if networkName != "" {
		networkId = resolveNetworkId(networkName, orgId, orgList)
	}

	//if deviceName != "" {}
	return orgId, networkId, DeviceId
}
























// Resolve Meraki Dashboard Organization name to id
func ResolveTESTFUNCDELETEMELATER(flags *pflag.FlagSet) (orgId string, err error) {
	orgName, _ := flags.GetString("organization")

	// ExactMatch array
	var exactMatch []Match

	// GreedyMatch array
	var greedyMatch []Match

	//Unmarshal data into data model
	organizations := configure.Organizations{}

	// API Call to get list of organizations
	dashboardApiCall := configure.GetOrganizations()

	// Evaluation: iteration of data in api.GetOrganizations API Call
	for _, metadata := range dashboardApiCall {
		response, _ := json.Marshal(metadata.Payload)
		_ = json.Unmarshal(response, &organizations)
	}

	// eval orgs in the org list
	for _, organization := range organizations {

		// match data
		match := Match{
			Name: organization.Name,
			ID: organization.ID,
		}

		matchRegex := regexp.MustCompile(orgName)
		greedy := matchRegex.MatchString(organization.Name)

		if organization.Name == orgName {
			exactMatch = append(exactMatch, match)
		} else if greedy == true {
			greedyMatch = append(greedyMatch, match) }
	}

	switch exactMatchLen := len(exactMatch); {
	case exactMatchLen == 1:
		fmt.Println("The Len of exactMatch list is: ", len(exactMatch))

	case exactMatchLen != 1:
		// Condition: length of ExactMatch array is 1
		// return ExactMatch.OrgId
		fmt.Println("The Len of DOES NOT EQUAL 1: ", len(exactMatch))
	}

	return orgId, err
}