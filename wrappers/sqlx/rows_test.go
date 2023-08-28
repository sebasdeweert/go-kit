package sqlx

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/sebasdeweert/go-kit/wrappers/sql"
)

func Test_Rows(t *testing.T) {
	Convey("*Rows", t, func() {
		Convey("Implements the SQLXRows interface", func() {
			i := reflect.TypeOf((*SQLXRows)(nil)).Elem()

			var tx interface{} = Rows{}

			So(reflect.PtrTo(reflect.TypeOf(tx)).Implements(i), ShouldBeTrue)
		})

		Convey("Implements the SQLRows interface", func() {
			i := reflect.TypeOf((*sql.SQLRows)(nil)).Elem()

			var tx interface{} = Rows{}

			So(reflect.PtrTo(reflect.TypeOf(tx)).Implements(i), ShouldBeTrue)
		})
	})
}
