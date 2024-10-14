# postgres

go wrappers, migrations, and testcontainers for postgres

![workflow](https://github.com/ddbgio/postgres/actions/workflows/go.yml/badge.svg)

## 🏗️ testcontainers
[Testcontainers](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/) provides a container setup process from within the go application.

Because Linux requires root (or special permissions) for the Docker socket, it may be necessary to run sudo to run tests. Because `sudo` may not have go install added to `PATH`, use `$(which go)` to pass the full path of the go command. This is not necessary for CI tests, which have preconfigured permissions to the Docker socket.
```shell
sudo $(which go) test -v
```

## ⚙️ migrations
Files for initailizing (and tearing down) a database are included in the `Migrations()` function. See [`example-migrations`](https://github.com/ddbgio/postgres/tree/main/example-migrations) for examples.

> [!TIP]
> Migrations files should be numbered in the order they are to be run, first `down`, then `up`. As such, all `down` migrations should be odd; all `up` migrations should be even. Migrations will return a list of `Migration` objects, containing _Direction_, _Filename_, and _Content_.


### usage
Use either `local` or `embed` filesystems to produce a slice of `Migration`s
```go
// Migrations represents a single SQL migration file,
// including the direction, file name, and content
type Migration struct {
    Direction string
    Filename  string
    Content   string
}
```

#### local
For a local filesystem, just provide the path.
```go
m, err := Migrations(os.DirFS("."), "migrations-dir-name", "up")
if err := nil {
    return fmt.Errorf("migration failed: %w", err)
}
```

#### embed
For an embed filesystem, provide the `embed.FS` and a filename.
```go
//go:embed migrations
var migrationsDirEmbed embed.FS

migrations, err := Migrations(migrationsDirEmbed, "migrations-dir-name", "up")
if err := nil {
    return fmt.Errorf("migration failed: %w", err)
}
```

#### applying migrations
Loop over the migrations to apply them to your database.
```go
for _, m := range migrations {
    log.Info("running migration",
        "file", m.Filename,
        "direction", m.Direction,
    )
    err := sql.Execute(m.Content)
    if err != nil {
        return fmt.Errorf("sql error on migration: %w", err)
    }
}
```