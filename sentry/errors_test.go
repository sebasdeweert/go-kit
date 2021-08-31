package sentry

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrNoLevels(t *testing.T) {
	Convey("ErrNoLevels", t, func() {
		Convey("Contains an empty levels error", func() {
			So(ErrNoLevels.Error(), ShouldEqual, "no levels defined")
		})
	})
}
