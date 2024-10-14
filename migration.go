package postgres

import (
	"embed"
	"fmt"
	"io/fs"
	"path"
	"strings"
)

//go:embed example-migrations
var migrationsDirEmbed embed.FS

const exampleMigrationsDir = "example-migrations"

// Migrations represents a single SQL migration file.
type Migration struct {
	Direction string
	Filename  string
	Content   string
}

// Migrations reads an arbirary file system for *.<direction>.sql files,
// returning a list of "up" or "down" migration SQL objects with:
//   - filename
//   - content
//   - direction
//
// # local
//
//	m, err := Migrations(os.DirFS("."), "migrations-dir-name", "up")
//
// # embed
//
//	//go:embed migrations
//	var migrationsDirEmbed embed.FS
//	migrations, err := Migrations(migrationsDirEmbed, "migrations-dir-name", "up")
func Migrations(fsys fs.FS, directory string, direction string) ([]Migration, error) {
	entries, err := fs.ReadDir(fsys, directory)
	if err != nil {
		return nil, fmt.Errorf("unable to read directory %q: %w", directory, err)
	}
	var migrations []Migration
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		content, err := fs.ReadFile(fsys, path.Join(directory, entry.Name()))
		if err != nil {
			return nil, fmt.Errorf("unable to read file %q: %w", entry.Name(), err)
		}
		expectedEnding := fmt.Sprintf("%s.sql", direction)
		if strings.Contains(entry.Name(), expectedEnding) {
			var migration = Migration{
				Filename:  entry.Name(),
				Content:   string(content),
				Direction: direction,
			}
			migrations = append(migrations, migration)
		}
	}
	return migrations, nil
}
