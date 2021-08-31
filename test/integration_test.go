package test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `["qux", "qix"]`)
}

type Body struct{}

func (*Body) Close() error {
	return nil
}

func (*Body) Read(p []byte) (n int, err error) {
	return 0, errors.New("foo")
}

func Test_NewSuite(t *testing.T) {
	Convey("NewSuite()", t, func() {
		Convey("Returns a Suite instance with given arguments", func() {
			suite := NewSuite(&testing.T{}, "foo")

			So(suite, ShouldResemble, &Suite{
				T:        &testing.T{},
				basePath: "foo",
			})
		})
	})
}

func Test_Suite_Request(t *testing.T) {
	Convey("*Suite.Request()", t, func() {
		Convey("Fails if it's not possible to create a request", func() {
			suite := &Suite{
				T:        &testing.T{},
				basePath: ":",
			}

			rsp := suite.Request("", "", nil, nil)

			So(suite.T.Failed(), ShouldBeTrue)
			So(rsp, ShouldBeNil)
		})

		Convey("Fails if it's not possible to execute the request", func() {
			suite := &Suite{T: &testing.T{}}

			rsp := suite.Request("qux", "qix", nil, nil)

			So(suite.T.Failed(), ShouldBeTrue)
			So(rsp, ShouldBeNil)
		})

		Convey("Returns the request response", func() {
			server := httptest.NewServer(&Handler{})

			defer server.Close()

			suite := &Suite{
				T:        &testing.T{},
				basePath: server.URL,
			}

			rsp := suite.Request("GET", "/", map[string]string{"foo": "bar"}, nil)

			defer rsp.Body.Close()

			raw, err := ioutil.ReadAll(rsp.Body)

			if err != nil {
				t.Fatal(err)
			}

			So(string(raw), ShouldEqual, `["qux", "qix"]`)
		})
	})
}

func Test_Suite_StringResponse(t *testing.T) {
	Convey("*Suite.StringResponse()", t, func() {
		Convey("Fails if body cannot be read", func() {
			suite := &Suite{T: &testing.T{}}

			rsp := suite.StringResponse(&http.Response{
				Body: &Body{},
			})

			So(suite.T.Failed(), ShouldBeTrue)
			So(rsp, ShouldBeBlank)
		})

		Convey("Returns the stringified response body", func() {
			server := httptest.NewServer(&Handler{})

			defer server.Close()

			rsp, err := http.Get(server.URL)

			if err != nil {
				t.Fatal(err)
			}

			suite := &Suite{}

			So(suite.StringResponse(rsp), ShouldEqual, `["qux", "qix"]`)
		})
	})
}

func Test_Suite_DecodeResponse(t *testing.T) {
	Convey("*Suite.DecodeResponse()", t, func() {
		Convey("Fails if body cannot be decoded", func() {
			suite := &Suite{T: &testing.T{}}

			suite.DecodeResponse(
				&http.Response{
					Body: &Body{},
				},
				nil,
			)

			So(suite.T.Failed(), ShouldBeTrue)
		})

		Convey("Returns the decoded response body", func() {
			server := httptest.NewServer(&Handler{})

			defer server.Close()

			rsp, err := http.Get(server.URL)

			if err != nil {
				t.Fatal(err)
			}

			suite := &Suite{}

			var result []string

			suite.DecodeResponse(rsp, &result)

			So(result, ShouldResemble, []string{"qux", "qix"})
		})
	})
}
