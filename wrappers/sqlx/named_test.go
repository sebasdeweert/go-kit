package sqlx

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_NamedStmt(t *testing.T) {
	Convey("*NamedStmt", t, func() {
		Convey("Implements the SQLXNamedStmt interface", func() {
			i := reflect.TypeOf((*SQLXNamedStmt)(nil)).Elem()

			var stmt interface{} = NamedStmt{}

			So(reflect.PtrTo(reflect.TypeOf(stmt)).Implements(i), ShouldBeTrue)
		})
	})
}
