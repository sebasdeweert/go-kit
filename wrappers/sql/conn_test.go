package sql

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Conn(t *testing.T) {
	Convey("*Conn", t, func() {
		Convey("Implements the SQLConn interface", func() {
			i := reflect.TypeOf((*SQLConn)(nil)).Elem()

			var conn interface{} = Conn{}

			So(reflect.PtrTo(reflect.TypeOf(conn)).Implements(i), ShouldBeTrue)
		})
	})
}
