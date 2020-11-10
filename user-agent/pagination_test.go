package user_agent

/*
Modified version of Link: https://github.com/peterhellberg/link
Special thanks to peterhellberg.
*/

import (
	"net/http"
	"testing"
)

func TestParseResponse(t *testing.T) {
	resp := &http.Response{Header: http.Header{}}
	resp.Header.Set("Link", `<https://n1.meraki.com/api/v1/organizations/1234567890/devices?perPage=1000&startingAfter=0000-0000-0000>; rel=first, <https://n1.meraki.com/api/v1/organizations/1234567890/devices?endingBefore=ZZZZ-ZZZZ-ZZZZ&perPage=1000>; rel=last`)
	var Response LinkedList

	Response = ParseResponse(resp)

	for _, g := range Response {

		if got, want := g.URI, "https://n1.meraki.com/api/v1/organizations/1234567890/devices"; got != want {
			t.Fatalf(`URI = %q, want %q`, got, want)
		}

		if got, want := g.PerPage, "1000"; got != want {
			t.Fatalf(`PerPage = %q, want %q`, got, want)
		}

		switch g.Rel {
		case "first":
			if got, want := g.Rel, "first"; got != want {
				t.Fatalf(`Rel = %q, want %q`, got, want)
			}
		case "next":
			if got, want := g.Rel, "next"; got != want {
				t.Fatalf(`Rel = %q, want %q`, got, want)
			}
		case "prev":
			if got, want := g.Rel, "prev"; got != want {
				t.Fatalf(`Rel = %q, want %q`, got, want)
			}
		case "last":
			if got, want := g.Rel, "last"; got != want {
				t.Fatalf(`Rel = %q, want %q`, got, want)
			}
		default:
			t.Fatalf(`Unknown Rel = %q`, g.Rel)
		}

	}
}
