package gob

import (
	"bytes"
	"encoding/gob"
	"testing"

	"github.com/sebasdeweert/go-kit/types"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_NewEncoder(t *testing.T) {
	Convey("NewEncoder()", t, func() {
		Convey("Returns a new encoder", func() {
			So(NewEncoder(), ShouldResemble, &encoder{})
		})
	})
}

func Test_encoder_Encode(t *testing.T) {
	enc := &encoder{}

	Convey("*encoder.Encode()", t, func() {
		Convey("Returns nil and an error when encoding fails", func() {
			encoded, err := enc.Encode(nil)

			So(encoded, ShouldBeNil)
			So(err.Error(), ShouldEqual, "gob: cannot encode nil value")
		})

		Convey("Returns a string pointer encoding the given object and nil", func() {
			encoded, err := enc.Encode(types.String("foo"))

			var buffer bytes.Buffer

			buffer.WriteString(*encoded)

			var decoded string

			gob.NewDecoder(&buffer).Decode(&decoded)

			So(err, ShouldBeNil)
			So(decoded, ShouldEqual, "foo")
		})
	})
}

func Test_encoder_Decode(t *testing.T) {
	enc := &encoder{}

	Convey("*encoder.Decode()", t, func() {
		Convey("Returns an error when the given object cannot be decoded", func() {
			err := enc.Decode("", nil)

			So(err.Error(), ShouldEqual, "EOF")
		})

		Convey("Returns nil and decodes the given object", func() {
			var buffer bytes.Buffer

			gob.NewEncoder(&buffer).Encode("foo")

			var obj string

			err := enc.Decode(buffer.String(), &obj)

			So(err, ShouldBeNil)
			So(obj, ShouldEqual, "foo")
		})
	})
}
