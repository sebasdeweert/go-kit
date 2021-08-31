# migrations

This package offers you a `Migrator` interface with the `Migrate` method, which wraps the [rubenv/sql-migrate](https://github.com/rubenv/sql-migrate) package.

## Usage

Create a new `Migrator` interface with the `NewMigrator` method and execute the `Migrate` method:

```go
package foo

import (
  "database/sql"

  _ "github.com/go-sql-driver/mysql" // Load MySQL driver.
  "github.com/Sef1995/go-kit/migrations"
)

// NewDB returns a new DB and executes all pending migrations.
func NewDB() (*sql.DB, error) {
  db, err := sql.Open("mysql", "localhost@tcp(root)/mydb")

  if err != nil {
    return nil, err
  }

  migrator := migrations.NewMigrator(
    db,
    &migrations.Config{
      AllowOutdatedMigrations: true, // When set to true, the "unknown migration in database" error is omitted, otherwise returned.
      MigrationsDir:           "migrations",
      MigrationsTable:         "migrations",
    },
  )

  if err := migrator.Migrate(); err != nil {
    return nil, err
  }

  return db, nil
}
```
