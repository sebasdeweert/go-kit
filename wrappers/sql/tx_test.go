package sql

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Tx(t *testing.T) {
	Convey("*Tx", t, func() {
		Convey("Implements the SQLTx interface", func() {
			i := reflect.TypeOf((*SQLTx)(nil)).Elem()

			var tx interface{} = Tx{}

			So(reflect.PtrTo(reflect.TypeOf(tx)).Implements(i), ShouldBeTrue)
		})
	})
}
