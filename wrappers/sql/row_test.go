package sql

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Row(t *testing.T) {
	Convey("*Row", t, func() {
		Convey("Implements the SQLRow interface", func() {
			i := reflect.TypeOf((*SQLRow)(nil)).Elem()

			var row interface{} = Row{}

			So(reflect.PtrTo(reflect.TypeOf(row)).Implements(i), ShouldBeTrue)
		})
	})
}
