package config

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_NewLoader(t *testing.T) {
	Convey("NewLoader()", t, func() {
		Convey("Returns a new loader", func() {
			l := NewLoader("a")

			So(l, ShouldResemble, &loader{
				envPrefix: "a",
			})
		})
	})
}

func Test_loader_Load(t *testing.T) {
	Convey("*loader.Load()", t, func() {
		Convey("Returns an error if config.yml cannot be found", func() {
			loader := NewLoader("a")
			err := loader.Load("a")

			So(err.Error(), ShouldEqual, "config file cannot be found at ./config.yml")
		})

		Convey("Loads the config into the given struct", func() {
			conf := struct {
				Val    string
				Nested struct {
					Int int
					Map map[string]string
				}
			}{}

			confFilePath := "config.yml"

			if _, err := os.Stat(confFilePath); err == nil {
				t.Fatalf("file at %s already exists; not overwriting", confFilePath)
			}

			// Create temporary config.yml.
			err := ioutil.WriteFile(confFilePath, []byte(`
val: myval
nested:
  int: 1
  map:
    a: b
`), 0644)

			if err != nil {
				t.Fatal(err)
			}

			// Remove the temporary config.yml file on return.
			defer func(confFilePath string) {
				if _, err := os.Stat(confFilePath); err == nil {
					err := os.Remove(confFilePath)

					if err != nil {
						t.Fatal(err)
					}
				}
			}(confFilePath)

			loader := NewLoader("")
			err = loader.Load(&conf)

			So(err, ShouldBeNil)
			So(conf, ShouldResemble, struct {
				Val    string
				Nested struct {
					Int int
					Map map[string]string
				}
			}{
				"myval",
				struct {
					Int int
					Map map[string]string
				}{
					1,
					map[string]string{"a": "b"},
				},
			})
		})
	})
}
