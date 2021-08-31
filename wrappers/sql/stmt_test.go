package sql

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Stmt(t *testing.T) {
	Convey("*Stmt", t, func() {
		Convey("Implements the SQLStmt interface", func() {
			i := reflect.TypeOf((*SQLStmt)(nil)).Elem()

			var stmt interface{} = Stmt{}

			So(reflect.PtrTo(reflect.TypeOf(stmt)).Implements(i), ShouldBeTrue)
		})
	})
}
