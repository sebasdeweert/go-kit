package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_ShouldHaveElements(t *testing.T) {
	Convey("ShouldHaveElements()", t, func() {
		Convey("Returns an error string if actual value has no type", func() {
			res := ShouldHaveElements(nil)

			So(res, ShouldEqual, "cannot assert type of <nil>")
		})

		Convey("Returns an error string if actual value is not a slice", func() {
			res := ShouldHaveElements("foo")

			So(res, ShouldEqual, "foo is not a slice")
		})

		Convey("Returns an error string if expected value has no type", func() {
			res := ShouldHaveElements([]interface{}{}, nil)

			So(res, ShouldEqual, "cannot assert type of <nil>")
		})

		Convey("Returns an error string if expected value is not a slice", func() {
			res := ShouldHaveElements([]interface{}{}, "foo")

			So(res, ShouldEqual, "foo is not a slice")
		})

		Convey("Returns an empty string if both values are nil", func() {
			var actual []interface{}
			var expected []interface{}

			res := ShouldHaveElements(actual, expected)

			So(res, ShouldBeBlank)
		})

		Convey("Returns an error string if the actual value is nil and the expected value is not", func() {
			var expected []interface{}

			res := ShouldHaveElements([]interface{}{}, expected)

			So(res, ShouldEqual, "cannot compare [] to nil")
		})

		Convey("Returns an error string if the expected value is nil and the actual value is not", func() {
			var actual []interface{}

			res := ShouldHaveElements(actual, []interface{}{})

			So(res, ShouldEqual, "cannot compare [] to nil")
		})

		Convey("Returns an error string if actual and expected values have different lengths", func() {
			res := ShouldHaveElements([]interface{}{}, []interface{}{"foo"})

			So(res, ShouldEqual, "[] and [foo] have different lengths")
		})

		Convey("Returns an error string if the actual value has unexpected elements", func() {
			res := ShouldHaveElements([]interface{}{"foo", "bar"}, []interface{}{"foo", "biz"})

			So(res, ShouldEqual, "bar is not expected")
		})

		Convey("Returns an empty string if both values have the same elements", func() {
			res := ShouldHaveElements([]interface{}{"foo", "bar"}, []interface{}{"bar", "foo"})

			So(res, ShouldBeBlank)
		})
	})
}
