package http

import (
	"io"
	"net/http"
	"net/url"
)

// Client interface to mock http.Client
type Client interface {
	Do(req *http.Request) (*http.Response, error)
	PostForm(url string, data url.Values) (resp *http.Response, err error)
	Get(url string) (resp *http.Response, err error)
	Head(url string) (resp *http.Response, err error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}
