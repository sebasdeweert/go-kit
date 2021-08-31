package pagination

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_LimitOffset(t *testing.T) {
	Convey("LimitOffset()", t, func() {
		Convey("Returns length as start and end when offset is greater than length ", func() {
			start, end := LimitOffset(10, 41, 40)

			So(start, ShouldEqual, 40)
			So(end, ShouldEqual, 40)
		})

		Convey("Returns offset as start and length as end when the sum of limit and offset is greater than length", func() {
			start, end := LimitOffset(4, 3, 6)

			So(start, ShouldEqual, 3)
			So(end, ShouldEqual, 6)
		})

		Convey("Returns offset as start and the sum of offset and limit as end otherwise", func() {
			start, end := LimitOffset(1, 2, 4)

			So(start, ShouldEqual, 2)
			So(end, ShouldEqual, 3)
		})
	})
}
