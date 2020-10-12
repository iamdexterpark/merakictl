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
	valueCommaRegexp = regexp.MustCompile(`([^"]),`)
	equalRegexp      = regexp.MustCompile(` *= *`)
	keyRegexp        = regexp.MustCompile(`[a-z*]+`)
	linkRegexp       = regexp.MustCompile(`<(.+)>;`)
	semiRegexp       = regexp.MustCompile(`; +`)
	valRegexp        = regexp.MustCompile(`"+([^"]+)"+`)
)

// Group returned by Parse, contains multiple links indexed by "rel"
type Group map[string]*Link

// Link contains a Link item with URI, Rel, and other non-URI components in Extra.
type Link struct {
	URI   string
	Rel   string
	Extra map[string]string
}

// String returns the URI
func (l *Link) String() string {
	return l.URI
}

// ParseRequest parses the provided *http.Request into a Group
func ParseRequest(req *http.Request) Group {
	if req == nil {
		return nil
	}

	return ParseHeader(req.Header)
}

// ParseResponse parses the provided *http.Response into a Group
func ParseResponse(resp *http.Response) Group {
	if resp == nil {
		return nil
	}

	return ParseHeader(resp.Header)
}

// ParseHeader retrieves the Link header from the provided http.Header and parses it into a Group
func ParseHeader(h http.Header) Group {
	if headers, found := h["Link"]; found {
		return Parse(strings.Join(headers, ", "))
	}

	return nil
}

// Parse parses the provided string into a Group
func Parse(s string) Group {
	if s == "" {
		return nil
	}

	group := Group{}

	for _, l := range commaRegexp.Split(s, -1) {
		// s is the whole string unmodified
		// l is the string split by commas


		// linkMatches is a nested array with two versions of the url
		// linkMatches[0] is the url with <>;
		// linkMatches[1] is the url without <>;
		linkMatches := linkRegexp.FindAllStringSubmatch(l, -1)

		if len(linkMatches) == 0 {
			return nil
		}

		// pieces takes the double array of linkMatches and extracts the inner array so we can access the values
		pieces := linkMatches[0]


		// pieces[1] is the url with <>;
		link := &Link{URI: pieces[1], Extra: map[string]string{}}

		// extra takes l and splits it by semicolon
		for _, extra := range semiRegexp.Split(l, -1) {

			// vals is the splits by the = equal sign
			vals := equalRegexp.Split(extra, -1)


			// vals[0] is the rel
			// vals[1] is the value of rel


			// key is the rel
			key := keyRegexp.FindString(vals[0])


			val := vals[1]

			if key == "rel" {
				vals := strings.Split(val, " ")
				rels := []string{vals[0]}

				if len(vals) > 1 {
					for _, v := range vals[1:] {
						if !strings.HasPrefix(v, "http") {
							rels = append(rels, v)
						}
					}
				}

				rel := strings.Join(rels, " ")

				link.Rel = rel
				group[rel] = link
			} else {
				link.Extra[key] = val
			}
		}
	}

	return group
}