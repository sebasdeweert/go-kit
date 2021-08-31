package test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

type Suite struct {
	T        *testing.T
	basePath string
}

func NewSuite(t *testing.T, basePath string) *Suite {
	return &Suite{
		t,
		basePath,
	}
}

func (s *Suite) Request(method string, endpoint string, headers map[string]string, body io.Reader) *http.Response {
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", s.basePath, endpoint), body)

	if err != nil {
		s.T.Error(err)

		return nil
	}

	for header, value := range headers {
		req.Header.Set(header, value)
	}

	if body != nil {
		req.Header.Set("content-type", "application/json")
	}

	rsp, err := http.DefaultClient.Do(req)

	if err != nil {
		s.T.Error(err)

		return nil
	}

	return rsp
}

func (s *Suite) StringResponse(rsp *http.Response) string {
	defer rsp.Body.Close()

	raw, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		s.T.Error(err)

		return ""
	}

	return string(raw)
}

func (s *Suite) DecodeResponse(rsp *http.Response, data interface{}) {
	defer rsp.Body.Close()

	if err := json.NewDecoder(rsp.Body).Decode(&data); err != nil {
		s.T.Error(err)
	}
}
