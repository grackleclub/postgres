package postgres

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMigrations(t *testing.T) {
	t.Run("embed up", func(t *testing.T) {
		m, err := Migrations(migrationsDirEmbed, exampleMigrationsDir, "up")
		require.NoError(t, err)
		require.NotEmpty(t, m)
		for _, migration := range m {
			t.Logf("embed up: %s", migration.Filename)
			t.Logf("content: %s", migration.Content)
		}
	})
	t.Run("embed down", func(t *testing.T) {
		m, err := Migrations(migrationsDirEmbed, exampleMigrationsDir, "down")
		require.NoError(t, err)
		require.NotEmpty(t, m)
		for _, migration := range m {
			t.Logf("embed down: %s", migration.Filename)
			t.Logf("content: %s", migration.Content)
		}
	})
	t.Run("local up", func(t *testing.T) {
		m, err := Migrations(os.DirFS("."), exampleMigrationsDir, "up")
		require.NoError(t, err)
		require.NotEmpty(t, m)
		for _, migration := range m {
			t.Logf("local up: %s", migration.Filename)
			t.Logf("content: %s", migration.Content)
		}
	})
	t.Run("local down", func(t *testing.T) {
		m, err := Migrations(os.DirFS("."), exampleMigrationsDir, "down")
		require.NoError(t, err)
		require.NotEmpty(t, m)
		for _, migration := range m {
			t.Logf("local down: %s", migration.Filename)
			t.Logf("content: %s", migration.Content)
		}
	})
}
