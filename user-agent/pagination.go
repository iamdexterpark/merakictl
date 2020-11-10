package user_agent

/*
Modified version of Link: https://github.com/peterhellberg/link
Special thanks to peterhellberg.
*/

import (
	"net/http"
	"regexp"
	"strings"
)

var (
	commaRegexp      = regexp.MustCompile(`,\s{0,}`)
	equalRegexp      = regexp.MustCompile(` *= *`)
	linkRegexp       = regexp.MustCompile(`<(.+)>`)
	semiRegexp       = regexp.MustCompile(`; +`)
	URIRegexp        = regexp.MustCompile(`[?*]+`)
	parametersRegexp = regexp.MustCompile(`[&*]+`)
)

// LinkedList returned by Parse, contains multiple links indexed by "rel"
type LinkedList []Link //map[string]*Link

// Link contains a Link item with URI, Rel, and other non-URI components in Extra.
type Link struct {
	URI           string
	Rel           string
	PerPage       string // The number of entries to be returned in the page (the current request)
	StartingAfter string // A token used by our server to indicate the starting "identifier" of the page
	EndingBefore  string // A token used by our server to indicate the ending "identifier" of the page
}

// String returns the URI
func (link *Link) String() string {
	return link.URI
}

// ParseRequest parses the provided *http.Request into a LinkedList
func ParseRequest(req *http.Request) LinkedList {
	if req == nil {
		return nil
	}
	return ParseHeader(req.Header)
}

// ParseResponse parses the provided *http.Response into a LinkedList
func ParseResponse(resp *http.Response) LinkedList {
	if resp == nil {
		return nil
	}

	return ParseHeader(resp.Header)
}

// ParseHeader retrieves the Link header from the provided http.Header and parses it into a LinkedList
func ParseHeader(h http.Header) LinkedList {
	if headers, found := h["Link"]; found {
		return Parse(strings.Join(headers, ", "))
	}

	return nil
}

// Parse parses the provided string into a LinkedList
func Parse(s string) LinkedList {

	// s is a comma separated list of link+rel value
	if s == "" {
		return nil
	}

	// Parse LinkList Header to extract an array of url+rel
	separatedLinks := SeparateLinks(s)

	// Store Links found in Header
	var LinkedList = LinkedList{}

	// Parse each link and return
	for _, link := range separatedLinks {
		LinkedList = append(LinkedList, ParseLink(link))

	}

	return LinkedList
}

func SeparateLinks(s string) []string {

	// Empty list to append links
	var links []string

	// splitting on commas so that we get a single link+rel to parse
	for _, parsedLink := range commaRegexp.Split(s, -1) {
		links = append(links, parsedLink)
	}
	return links
}

func ParseLink(s string) Link {

	// Split the url from the rel
	parseURI := semiRegexp.Split(s, -1)

	// parseURI[1] = rel=first
	if len(parseURI) == 0 {
		return Link{}
	}

	// strip <> from url
	formattedURL := linkRegexp.FindAllStringSubmatch(parseURI[0], -1)

	// get formatted parameters
	params := ParseLinkParameters(formattedURL[0][1])

	// get rel key-value pair
	rel := equalRegexp.Split(parseURI[1], -1)

	link := Link{
		URI:           params["uri"],
		Rel:           rel[1],
		PerPage:       params["perPage"],
		StartingAfter: params["startingAfter"],
		EndingBefore:  params["endingBefore"],
	}

	return link
}

func ParseLinkParameters(s string) map[string]string {

	// split the base url from the parameters
	splitParameters := URIRegexp.Split(s, -1)

	// parameters = perPage=1000&startingAfter=L_0
	parameters := splitParameters[1]

	// split parameters into a key-value list
	parameterList := parametersRegexp.Split(parameters, -1)

	// sort parameters into key-value struct
	var perPage string
	var startingAfter string
	var endingBefore string

	for _, param := range parameterList {

		formattedParam := equalRegexp.Split(param, -1)
		switch formattedParam[0] {
		case "perPage":
			perPage = formattedParam[1]
		case "startingAfter":
			startingAfter = formattedParam[1]
		case "endingBefore":
			endingBefore = formattedParam[1]
		}

	}

	// return parsed link data
	var parsedLink = map[string]string{
		"uri":           splitParameters[0],
		"perPage":       perPage,
		"startingAfter": startingAfter,
		"endingBefore":  endingBefore}

	return parsedLink

}
