package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var response = "[\"\",\n\"Aachener Straße\",\"Lars-Krüger-Hof\",\"Martinsweg (KG Gartenstadt Vahr)\",\n\"Züricher Straße\"]"
		_, _ = rw.Write([]byte(response))
	}))

	defer server.Close()

	response, _ := NewClient().GetContent(server.URL)

	if len(response.Addresses) != 4 {
		t.Fatalf(`GetContent(%s) should contain %d entries but was %d`, server.URL, 4, len(response.Addresses))
	}
	if response.notContains("Aachener Straße") {
		t.Fatalf(`GetContent(%s) should contain %s`, server.URL, "Aachener Straße")
	}
	if response.notContains("Lars-Krüger-Hof") {
		t.Fatalf(`GetContent(%s) should contain %s`, server.URL, "Lars-Krüger-Hof")
	}
	if response.notContains("Lars-Krüger-Hof") {
		t.Fatalf(`GetContent(%s) should contain %s`, server.URL, "Lars-Krüger-Hof")
	}
	if response.notContains("Züricher Straße") {
		t.Fatalf(`GetContent(%s) should contain %s`, server.URL, "Züricher Straße")
	}
	if response.contains("") {
		t.Fatalf(`GetContent(%s) should not contain empty string`, server.URL)
	}
}

func (r Response) notContains(e string) bool {
	return !r.contains(e)
}

func (r Response) contains(e string) bool {
	for _, a := range r.Addresses {
		if a == e {
			return true
		}
	}
	return false
}