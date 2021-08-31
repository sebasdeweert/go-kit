package sentry

import (
	"testing"

	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_NewHook(t *testing.T) {
	Convey("NewHook()", t, func() {
		Convey("Returns an error when the given DSN configuration is invalid", func() {
			hook, err := NewHook(Config{
				DSN: "foo",
			})

			So(hook, ShouldBeNil)
			So(err.Error(), ShouldEqual, "raven: dsn missing public key and/or password")
		})

		Convey("Returns an error when the levels are invalid logurs tokens", func() {
			hook, err := NewHook(Config{
				Levels: []string{"foo"},
			})

			So(hook, ShouldBeNil)
			So(err, ShouldResemble, ErrNoLevels)
		})

		Convey("Returns a hook when the configuration is valid", func() {
			h, err := NewHook(Config{
				Levels: []string{"warn"},
			})

			So(err, ShouldBeNil)
			So(h.Levels(), ShouldResemble, []logrus.Level{
				logrus.WarnLevel,
			})
		})

		Convey("Returns a hook with default level config", func() {
			h, err := NewHook(Config{})

			So(err, ShouldBeNil)
			So(h.Levels(), ShouldResemble, []logrus.Level{
				logrus.PanicLevel,
				logrus.FatalLevel,
				logrus.ErrorLevel,
			})
		})
	})
}
