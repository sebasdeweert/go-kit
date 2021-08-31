package sqlx

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Sef1995/go-kit/wrappers/sql"
)

func Test_DB(t *testing.T) {
	Convey("*DB", t, func() {
		Convey("Implements the SQLXDB interface", func() {
			i := reflect.TypeOf((*SQLXDB)(nil)).Elem()

			var db interface{} = DB{}

			So(reflect.PtrTo(reflect.TypeOf(db)).Implements(i), ShouldBeTrue)
		})

		Convey("Implements the SQLDB interface", func() {
			i := reflect.TypeOf((*sql.SQLDB)(nil)).Elem()

			var db interface{} = DB{}

			So(reflect.PtrTo(reflect.TypeOf(db)).Implements(i), ShouldBeTrue)
		})
	})
}
