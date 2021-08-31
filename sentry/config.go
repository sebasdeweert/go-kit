package sentry

import "time"

// Config holds the Sentry configuration.
type Config struct {
	DSN         string
	Environment string
	Release     string
	Timeout     time.Duration
	// Levels values: "panic", "fatal", "error", "warn", "info", "debug", "trace".
	Levels []string
}

// ShouldLog returns true if DSN and Environment are set, and false otherwise.
func (c Config) ShouldLog() bool {
	return c.DSN != "" && c.Environment != ""
}
