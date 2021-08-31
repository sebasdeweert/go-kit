package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_WithElements(t *testing.T) {
	Convey("WithElements", t, func() {
		Convey("Returns a withElementsMatcher with given expected value", func() {
			m := WithElements("foo")

			So(m, ShouldResemble, withElementsMatcher{
				expected: "foo",
			})
		})
	})
}

func Test_withElementsMatcher_Matches(t *testing.T) {
	Convey("withElementsMatcher.Matches()", t, func() {
		Convey("Returns false if actual value has no type", func() {
			m := withElementsMatcher{}

			So(m.Matches(nil), ShouldBeFalse)
		})

		Convey("Returns false if actual value is not a slice", func() {
			m := withElementsMatcher{}

			So(m.Matches("foo"), ShouldBeFalse)
		})

		Convey("Returns false if expected value has no type", func() {
			m := withElementsMatcher{}

			So(m.Matches([]interface{}{}), ShouldBeFalse)
		})

		Convey("Returns false if expected value is not a slice", func() {
			m := withElementsMatcher{
				expected: "foo",
			}

			So(m.Matches([]interface{}{}), ShouldBeFalse)
		})

		Convey("Returns true if both values are nil", func() {
			var actual []interface{}
			var expected []interface{}

			m := withElementsMatcher{expected}

			So(m.Matches(actual), ShouldBeTrue)
		})

		Convey("Returns false if the actual value is nil and the expected value is not", func() {
			var actual []interface{}

			m := withElementsMatcher{
				expected: []interface{}{},
			}

			So(m.Matches(actual), ShouldBeFalse)
		})

		Convey("Returns false if the expected value is nil and the actual value is not", func() {
			var expected []interface{}

			m := withElementsMatcher{expected}

			So(m.Matches([]interface{}{}), ShouldBeFalse)
		})

		Convey("Returns false if actual and expected values have different lengths", func() {
			m := withElementsMatcher{
				expected: []interface{}{},
			}

			So(m.Matches([]interface{}{"foo"}), ShouldBeFalse)
		})

		Convey("Returns false if the actual value has unexpected elements", func() {
			m := withElementsMatcher{
				expected: []interface{}{"foo", "bar"},
			}

			So(m.Matches([]interface{}{"foo", "biz"}), ShouldBeFalse)
		})

		Convey("Returns true if both values have the same elements", func() {
			m := withElementsMatcher{
				expected: []interface{}{"foo", "bar"},
			}

			So(m.Matches([]interface{}{"bar", "foo"}), ShouldBeTrue)
		})
	})
}

func Test_withElementsMatcher_String(t *testing.T) {
	Convey("withElementsMatcher.String()", t, func() {
		Convey("Returns a description string including the expected elements", func() {
			m := withElementsMatcher{
				expected: "foo",
			}

			So(m.String(), ShouldEqual, "is slice with same elements as foo")
		})
	})
}

func Test_WithRawString(t *testing.T) {
	Convey("WithRawString", t, func() {
		Convey("Returns a withRawStringMatcher with given expected value", func() {
			m := WithRawString(`
				foo

							bar    biz

					baz
			`)

			So(m, ShouldResemble, withRawStringMatcher{
				expected: "foo bar    biz baz",
			})
		})
	})
}

func Test_withRawStringMatcher_Matches(t *testing.T) {
	Convey("withRawStringMatcher.Matches()", t, func() {
		Convey("Returns false if actual value has no type", func() {
			m := withRawStringMatcher{}

			So(m.Matches(nil), ShouldBeFalse)
		})

		Convey("Returns false if actual value is not a string", func() {
			m := withRawStringMatcher{}

			So(m.Matches(1), ShouldBeFalse)
		})

		Convey("Returns false if the actual value has unexpected content", func() {
			m := withRawStringMatcher{
				expected: "foo bar biz",
			}

			var actual interface{} = `
				foo
						biz
				bar
			`

			So(m.Matches(actual), ShouldBeFalse)
		})

		Convey("Returns true if the actual value has the expected content", func() {
			m := withRawStringMatcher{
				expected: "foo bar    biz baz",
			}

			var actual interface{} = `
				foo

					bar    biz


						baz
			`

			So(m.Matches(actual), ShouldBeTrue)
		})
	})
}

func Test_withRawStringMatcher_String(t *testing.T) {
	Convey("withRawStringMatcher.String()", t, func() {
		Convey("Returns a description string including the expected content", func() {
			m := withRawStringMatcher{
				expected: "foo",
			}

			So(m.String(), ShouldEqual, "is string with content foo")
		})
	})
}
