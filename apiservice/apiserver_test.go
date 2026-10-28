package apiservice

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/fstest"
)

const (
	testPlaygroundPage = "<!DOCTYPE html>\n<html><body>playground</body></html>\n"
	// browserAccept is the navigation Accept header of Firefox and Safari.
	browserAccept = "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"
)

// doGraphQL sends one request to newGraphQLHandler over a real connection and
// fails the test if the server logs anything, such as a superfluous
// WriteHeader call that httptest.ResponseRecorder would not reveal.
func doGraphQL(t *testing.T, method, target, accept, body string) (*http.Response, string) {
	t.Helper()
	h, err := newGraphQLHandler(context.Background(), fstest.MapFS{
		"templates/graphql-playground.html": {Data: []byte(testPlaygroundPage)},
	})
	if err != nil {
		t.Fatal(err)
	}
	var logs bytes.Buffer
	srv := httptest.NewUnstartedServer(h)
	srv.Config.ErrorLog = log.New(&logs, "", 0)
	srv.Start()

	req, err := http.NewRequest(method, srv.URL+target, strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	if accept != "" {
		req.Header.Set("Accept", accept)
	}
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}
	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	srv.Close()
	if logs.Len() != 0 {
		t.Errorf("server log = %q, want none", logs.String())
	}
	return resp, string(b)
}

func TestGraphQLHandlerRouting(t *testing.T) {
	// "{" is rejected by the GraphQL parser, so no request below reaches gRPC.
	const query = "/graphql?query=%7B"
	for _, tc := range []struct {
		name, method, target, accept, body string
		wantPage                           bool
		wantJSON                           string
	}{
		{name: "browser", method: "GET", target: "/graphql", accept: browserAccept, wantPage: true},
		{name: "browser link with query", method: "GET", target: query, accept: browserAccept, wantPage: true},
		{name: "no Accept", method: "GET", target: "/graphql", wantPage: true},
		{name: "any type", method: "GET", target: "/graphql", accept: "*/*", wantPage: true},
		{name: "empty query", method: "GET", target: "/graphql?query=", accept: "*/*", wantPage: true},
		{name: "HTML preferred by weight", method: "GET", target: query, accept: "application/json;q=0.5, text/html", wantPage: true},
		{name: "invalid weight ignored", method: "GET", target: "/graphql", accept: "text/html;q=NaN, */*", wantPage: true},
		// GET queries still run without a token, as before.
		{name: "query without Accept", method: "GET", target: query, wantJSON: "Syntax Error"},
		{name: "query with any type", method: "GET", target: query, accept: "*/*", wantJSON: "Syntax Error"},
		{name: "JSON client", method: "GET", target: query, accept: "application/json", wantJSON: "Syntax Error"},
		{name: "JSON client without query", method: "GET", target: "/graphql", accept: "application/json", wantJSON: "Must provide an operation"},
		{name: "JSON preferred by weight", method: "GET", target: query, accept: "text/html;q=0.5, application/json", wantJSON: "Syntax Error"},
		{name: "HTML refused", method: "GET", target: query, accept: "application/json, text/html;q=0", wantJSON: "Syntax Error"},
		{name: "HTML refused without query", method: "GET", target: "/graphql", accept: "text/html;q=0, */*", wantJSON: "Must provide an operation"},
		{name: "only HTML, refused", method: "GET", target: "/graphql", accept: "text/html;q=0", wantJSON: "Must provide an operation"},
		// Other methods always reach GraphQL, which still requires a token for
		// POST and HEAD.
		{name: "POST from a browser", method: "POST", target: "/graphql", accept: browserAccept, body: `{"query":"{GetLatestBlockHeight{height}}"}`, wantJSON: "authorization header is required"},
		{name: "OPTIONS", method: "OPTIONS", target: "/graphql", accept: browserAccept, wantJSON: "invalid request method"},
		{name: "HEAD", method: "HEAD", target: "/graphql", accept: browserAccept},
	} {
		t.Run(tc.name, func(t *testing.T) {
			resp, body := doGraphQL(t, tc.method, tc.target, tc.accept, tc.body)
			ct := resp.Header.Get("Content-Type")
			if tc.wantPage {
				if body != testPlaygroundPage || !strings.HasPrefix(ct, "text/html") {
					t.Errorf("got %q (%s), want only the Playground page", body, ct)
				}
			} else if ct != "application/json" || strings.Contains(body, "<html") || !strings.Contains(body, tc.wantJSON) {
				t.Errorf("got %q (%s), want JSON containing %q", body, ct, tc.wantJSON)
			}
			wantVary := ""
			if tc.method == "GET" {
				wantVary = "Accept"
			}
			if vary := resp.Header.Get("Vary"); vary != wantVary {
				t.Errorf("Vary = %q, want %q", vary, wantVary)
			}
		})
	}
}
