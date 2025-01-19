package httpserver

import "net/http"

type HTTPClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}
