package http

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Client(t *testing.T) {
	Convey("Client", t, func() {
		Convey("Should be implemented by http.Client", func() {
			var c = http.Client{}

			So(&c, ShouldImplement, (*Client)(nil))
		})
	})
}
