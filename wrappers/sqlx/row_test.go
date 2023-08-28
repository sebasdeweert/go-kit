package sqlx

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/sebasdeweert/go-kit/wrappers/sql"
)

func Test_Row(t *testing.T) {
	Convey("*Row", t, func() {
		Convey("Implements the SQLXRow interface", func() {
			i := reflect.TypeOf((*SQLXRow)(nil)).Elem()

			var row interface{} = Row{}

			So(reflect.PtrTo(reflect.TypeOf(row)).Implements(i), ShouldBeTrue)
		})

		Convey("Implements the SQLRow interface", func() {
			i := reflect.TypeOf((*sql.SQLRow)(nil)).Elem()

			var row interface{} = Row{}

			So(reflect.PtrTo(reflect.TypeOf(row)).Implements(i), ShouldBeTrue)
		})
	})
}
