package sql

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Rows(t *testing.T) {
	Convey("*Rows", t, func() {
		Convey("Implements the SQLRows interface", func() {
			i := reflect.TypeOf((*SQLRows)(nil)).Elem()

			var rows interface{} = Rows{}

			So(reflect.PtrTo(reflect.TypeOf(rows)).Implements(i), ShouldBeTrue)
		})
	})
}
