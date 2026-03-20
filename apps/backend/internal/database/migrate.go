package database

import (
	"context"
	"fmt"
	"log"
)

// Migrate runs all pending SQL migrations
func Migrate(db *DB) error {
	log.Println("Running database migrations...")

	ctx := context.Background()

	migrations := []struct {
		name string
		sql  string
	}{
		{
			name: "create_users_table",
			sql: `
				CREATE TABLE IF NOT EXISTS users (
					id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
					email       VARCHAR(255) UNIQUE NOT NULL,
					name        VARCHAR(255) NOT NULL,
					created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
					updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
				);
			`,
		},
		{
			name: "create_migrations_table",
			sql: `
				CREATE TABLE IF NOT EXISTS schema_migrations (
					id         SERIAL PRIMARY KEY,
					name       VARCHAR(255) UNIQUE NOT NULL,
					applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
				);
			`,
		},
	}

	for _, m := range migrations {
		// Check if migration already applied
		var count int
		err := db.Pool.QueryRow(ctx,
			"SELECT COUNT(*) FROM schema_migrations WHERE name = $1", m.name,
		).Scan(&count)

		// Table might not exist yet (first run), ignore that error
		if err != nil {
			count = 0
		}

		if count > 0 {
			log.Printf("Migration '%s' already applied, skipping", m.name)
			continue
		}

		// Apply migration
		if _, err := db.Pool.Exec(ctx, m.sql); err != nil {
			return fmt.Errorf("migration '%s' failed: %w", m.name, err)
		}

		// Record migration (ignore error if migrations table doesn't exist yet)
		_, _ = db.Pool.Exec(ctx,
			"INSERT INTO schema_migrations (name) VALUES ($1)", m.name,
		)

		log.Printf("Applied migration: %s", m.name)
	}

	log.Println("Migrations complete")
	return nil
}
