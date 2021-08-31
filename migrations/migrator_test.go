package migrations

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_NewMigrator(t *testing.T) {
	Convey("NewMigrator()", t, func() {
		Convey("Returns a new Migrator", func() {
			m := NewMigrator(
				&sql.DB{},
				&Config{
					AllowOutdatedMigrations: true,
					MigrationsDir:           "foo",
					MigrationsTable:         "bar",
				},
			)

			So(m.(*migrator), ShouldResemble, &migrator{
				db: &sql.DB{},
				config: &Config{
					AllowOutdatedMigrations: true,
					MigrationsDir:           "foo",
					MigrationsTable:         "bar",
				},
			})
		})
	})
}

func Test_migrator_Migrate(t *testing.T) {
	Convey("*migrator.Migrate()", t, func() {
		Convey("Returns the error returned by migrate.Exec() if config.AllowOutdatedMigrations is false", func() {
			db, _ := sql.Open("mysql", "foo")
			m := NewMigrator(
				db,
				&Config{
					AllowOutdatedMigrations: false,
				},
			)

			err := m.Migrate()

			So(err.Error(), ShouldEqual, "invalid DSN: missing the slash separating the database name")
		})

		Convey("Returns the error returned by migrate.Exec() if config.AllowOutdatedMigrations is true", func() {
			db, _ := sql.Open("mysql", "foo")
			m := NewMigrator(
				db,
				&Config{
					AllowOutdatedMigrations: true,
				},
			)

			err := m.Migrate()

			So(err.Error(), ShouldEqual, "invalid DSN: missing the slash separating the database name")
		})
	})
}
