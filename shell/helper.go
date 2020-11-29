package shell

import (
	"encoding/json"
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

	// build regex for greedy matching
	matchRegex := regexp.MustCompile(regExString)


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


		// Search for exact results
		if data.Name == regExString {
			exactResults = append(exactResults, match)
		}

		// Search for greedy results
		greedy := matchRegex.MatchString(data.Name)
		if greedy == true {
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
				Organization: data.Organization,
				Network: data.Network,
				Serial: data.Serial,
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
			//pretty.Println(exactOrg)
			//pretty.Println(greedyOrg)

			if exactOrg.ID != "" {

				if exactOrg.Name == orgName {
					mResults, _ := json.Marshal(exactOrg)
					json.Unmarshal(mResults, &exactMatch)
					orgId = exactMatch.ID
				}
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



func resolveDeviceId (deviceName, networkId string, orgList []string) (deviceId string) {
	// ExactMatch array
	//var exactMatch Match

	// networkIdList array
	var deviceIdList []Match

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
		metadata := api.GetOrganizationDevices(org, "", "","")
		var paginatedData []Match

		for _, data := range metadata {
			format, _ := json.Marshal(data.Payload)
			json.Unmarshal(format, &paginatedData)
			deviceIdList = append(deviceIdList, paginatedData...)
		}
	}

	// iterate through list of orgs and append to exact/greedy lists
	exactnet, greedynet := ResolveMatches(deviceName, deviceIdList)

	if exactnet.Name == deviceName {

		if len(networkId) == 0 {
			deviceId = exactnet.Serial
		} else if len(networkId) != 0 {

			// ensure we dont have a mismatch with networkID
			if networkId != exactnet.Network {
				pretty.Println(greedynet)
				pretty.Println("No exact match for -s flag. Please use an explicit networkId, Name or use -n flag in conjunction with this call ^")
				os.Exit(1)
			}
			deviceId = exactnet.Serial
		}
	} else if len(networkId) != 0 {
		for _, greedyMatch := range greedynet {

			if greedyMatch.Network == networkId {
				deviceId = greedyMatch.Serial
			}
		}

	} else {
		pretty.Println(greedynet)
		pretty.Println("No exact match for -s flag. Please use an explicit networkId, Name or use -n flag in conjunction with this call ^")
		os.Exit(1)
	}

	return deviceId
}


func ResolveFlags(flags *pflag.FlagSet) (orgId, networkId, deviceId string) {

	// store flags as variables
	orgName, _ := flags.GetString("organization")
	networkName, _ := flags.GetString("network")
	deviceName, _ := flags.GetString("hostname")

	var orgList []string
	if orgName != "" {
		orgId, orgList = resolveOrgId(orgName)
		pretty.Println("-----------------------------------------------------------------")
		pretty.Println("orgId", orgId)
		pretty.Println("orgList", orgList)
		pretty.Println("-----------------------------------------------------------------")
	}

	if networkName != "" {
		networkId = resolveNetworkId(networkName, orgId, orgList)
		pretty.Println("-----------------------------------------------------------------")
		pretty.Println("networkName", networkName)
		pretty.Println("networkId", networkId)
		pretty.Println("-----------------------------------------------------------------")
	}

	if deviceName != "" {
		deviceId = resolveDeviceId(deviceName, networkId, orgList)
		pretty.Println("-----------------------------------------------------------------")
		pretty.Println("deviceName", deviceName)
		pretty.Println("Serial", deviceId)
		pretty.Println("-----------------------------------------------------------------")
	}

	return orgId, networkId, deviceId
}