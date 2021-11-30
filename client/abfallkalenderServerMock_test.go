package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	streetsContextPath = "/bremenabfallkalender/(S(nni))/Data/Strassen"
	streetsResponse    = "[\"\",\n\"Aachener Straße\",\"Lars-Krüger-Hof\",\"Martinsweg (KG Gartenstadt Vahr)\",\n\"Züricher Straße\"]"
)

type AbfallkalenderServer struct {
	server             *httptest.Server
	BaseUrl            string
	StreetsContextPath string
}

func (s *AbfallkalenderServer) Close() {
	s.server.Close()
}

func startAbfallkalenderServer(t *testing.T) AbfallkalenderServer {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		switch req.URL.String() {
		case streetsContextPath:
			doGetStreets(t, rw, req)
			break
		default:
			_ = fmt.Sprintf("URL %s not known on test server", req.URL.String())
			t.FailNow()
		}
	}))

	return AbfallkalenderServer{server: server, BaseUrl: server.URL, StreetsContextPath: streetsContextPath}
}

func doGetStreets(t *testing.T, rw http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		_ = fmt.Sprintf("%s %s, want: GET", req.Method, streetsContextPath)
		t.FailNow()
	}
	_, _ = rw.Write([]byte(streetsResponse))
}
