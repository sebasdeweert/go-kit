package cache

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Direction(t *testing.T) {
	Convey("Ascending", t, func() {
		Convey("Is equal 0", func() {
			So(Ascending, ShouldEqual, 0)
		})
	})

	Convey("Descending", t, func() {
		Convey("Is equal 1", func() {
			So(Descending, ShouldEqual, 1)
		})
	})
}

func Test_ErrInvalidDirection(t *testing.T) {
	Convey("ErrInvalidDirection", t, func() {
		Convey("Contains an invalid direction message", func() {
			So(ErrInvalidDirection, ShouldEqual, "the direction must be either Ascending or Descending")
		})
	})
}

func Test_ErrCacheMiss(t *testing.T) {
	Convey("ErrCacheMiss", t, func() {
		Convey("Contains a cache miss message", func() {
			So(ErrCacheMiss, ShouldResemble, errors.New("cache miss"))
		})
	})
}

func Test_ErrDestNotPtrSlice(t *testing.T) {
	Convey("ErrDestNotPtrSlice", t, func() {
		Convey("Contains not pointer slice error", func() {
			So(ErrDestNotPtrSlice, ShouldResemble, errors.New("destination is not a slice type"))
		})
	})
}
