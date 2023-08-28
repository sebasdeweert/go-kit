package sqlx

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/sebasdeweert/go-kit/wrappers/sql"
)

func Test_TX(t *testing.T) {
	Convey("*TX", t, func() {
		Convey("Implements the SQLXTX interface", func() {
			i := reflect.TypeOf((*SQLXTx)(nil)).Elem()

			var tx interface{} = Tx{}

			So(reflect.PtrTo(reflect.TypeOf(tx)).Implements(i), ShouldBeTrue)
		})

		Convey("Implements the SQLTX interface", func() {
			i := reflect.TypeOf((*sql.SQLTx)(nil)).Elem()

			var tx interface{} = Tx{}

			So(reflect.PtrTo(reflect.TypeOf(tx)).Implements(i), ShouldBeTrue)
		})
	})
}
