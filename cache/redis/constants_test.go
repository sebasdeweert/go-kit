package redis

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrNilResponse(t *testing.T) {
	Convey("ErrNilResponse", t, func() {
		Convey("Contains a nil Redis response message", func() {
			So(ErrNilResponse, ShouldEqual, "redis: nil")
		})
	})
}

func Test_ErrNoExpirationSet(t *testing.T) {
	Convey("ErrNoExpirationSet", t, func() {
		Convey("Contains a no expiration set message", func() {
			So(ErrNoExpirationSet, ShouldEqual, "no expiration set")
		})
	})
}

func Test_ErrNotPong(t *testing.T) {
	Convey("ErrNotPong", t, func() {
		Convey("Contains a not expected ping response message", func() {
			So(ErrNotPong, ShouldEqual, "unexpected return value")
		})
	})
}

func Test_PongMessage(t *testing.T) {
	Convey("PongMessage", t, func() {
		Convey("Contains a PONG message", func() {
			So(PongMessage, ShouldEqual, "PONG")
		})
	})
}
