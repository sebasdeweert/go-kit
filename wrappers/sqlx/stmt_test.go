package sqlx

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Sef1995/go-kit/wrappers/sql"
)

func Test_Stmt(t *testing.T) {
	Convey("*Stmt", t, func() {
		Convey("Implements the SQLXStmt interface", func() {
			i := reflect.TypeOf((*SQLXStmt)(nil)).Elem()

			var tx interface{} = Stmt{}

			So(reflect.PtrTo(reflect.TypeOf(tx)).Implements(i), ShouldBeTrue)
		})

		Convey("Implements the SQLStmt interface", func() {
			i := reflect.TypeOf((*sql.SQLStmt)(nil)).Elem()

			var tx interface{} = Stmt{}

			So(reflect.PtrTo(reflect.TypeOf(tx)).Implements(i), ShouldBeTrue)
		})
	})
}
