package sql

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_DB(t *testing.T) {
	Convey("*DB", t, func() {
		Convey("Implements the SQLDB interface", func() {
			i := reflect.TypeOf((*SQLDB)(nil)).Elem()

			var db interface{} = DB{}

			So(reflect.PtrTo(reflect.TypeOf(db)).Implements(i), ShouldBeTrue)
		})
	})
}
