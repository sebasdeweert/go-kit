package migrations

import (
	"database/sql"

	migrate "github.com/rubenv/sql-migrate"
)

// Migrator interface to mock the Migrate method.
type Migrator interface {
	Migrate() error
}

type migrator struct {
	db     *sql.DB
	config *Config
}

// NewMigrator returns a new Migrator interface.
func NewMigrator(db *sql.DB, c *Config) Migrator {
	return &migrator{
		db,
		c,
	}
}

// Migrate proxies migrate.Exec() omitting the "unknown migration in database" error
// if the AllowOutdatedMigrations configuration is enbaled.
func (m *migrator) Migrate() error {
	migrate.SetTable(m.config.MigrationsTable)

	_, err := migrate.Exec(
		m.db,
		"mysql",
		&migrate.FileMigrationSource{
			Dir: m.config.MigrationsDir,
		},
		migrate.Up,
	)

	if err != nil {
		if !m.config.AllowOutdatedMigrations {
			return err
		}

		if err2, ok := err.(*migrate.PlanError); !ok || err2.ErrorMessage != "unknown migration in database" {
			return err
		}
	}

	return nil
}
