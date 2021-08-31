package sentry

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Config_ShouldLog(t *testing.T) {
	Convey("Config.ShouldLog()", t, func() {
		Convey("Returns false if DSN and Environment are blank", func() {
			c := Config{}

			So(c.ShouldLog(), ShouldBeFalse)
		})

		Convey("Returns false if DSN is blank", func() {
			c := Config{
				Environment: "foo",
			}

			So(c.ShouldLog(), ShouldBeFalse)
		})

		Convey("Returns false if Environment is blank", func() {
			c := Config{
				DSN: "foo",
			}

			So(c.ShouldLog(), ShouldBeFalse)
		})

		Convey("Returns true if DSN and Environment are not blank", func() {
			c := Config{
				DSN:         "foo",
				Environment: "bar",
			}

			So(c.ShouldLog(), ShouldBeTrue)
		})
	})
}
