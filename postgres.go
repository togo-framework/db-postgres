// Package dbpostgres is togo's PostgreSQL database driver plugin.
//
// It registers the pgx database/sql driver (driver name "pgx") so the kernel's
// SQL ORM can talk to PostgreSQL — without the app importing the raw driver.
// Use it by setting DB_DRIVER=pgx + DATABASE_URL=postgres://… and installing the
// plugin (`togo new --db postgres`, or `togo install togo-framework/db-postgres`,
// or a blank import). The kernel ships the postgres SQL dialect; this plugin only
// supplies the driver. Wire-compatible stacks (Supabase, togo-postgres) reuse it.
package dbpostgres

import (
	_ "github.com/jackc/pgx/v5/stdlib" // registers the "pgx" database/sql driver

	"github.com/togo-framework/togo"
)

func init() {
	// Marker provider so the kernel lists/recognizes the plugin. The blank import
	// above already registered the driver with database/sql at import time, and
	// togo.DialectFor knows the postgres dialect, so this is a no-op boot.
	togo.RegisterProviderFunc("db-postgres", togo.PriorityService, func(*togo.Kernel) error { return nil })
}
