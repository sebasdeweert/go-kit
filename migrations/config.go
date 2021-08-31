package migrations

// Config holds the migrations configuration.
type Config struct {
	AllowOutdatedMigrations bool
	MigrationsDir           string
	MigrationsTable         string
}
