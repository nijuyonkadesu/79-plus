package chat

import (
	"chat/domain"
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
)

//go:embed create-tables.sql
var schema string

type Type int

const (
	SQLite Type = iota
	Postgres
	Memstore
)

type UserStore struct {
	db *sql.DB
}

type Config struct {
	Database string // Database name (default: "first-try")
	Backend  Type   // Backend type (default: "sqlite") // TODO: make it one of postgres, sqlite
}

func New(ctx context.Context, cfg *Config) (*UserStore, error) {
	var db *sql.DB
	var err error

	switch cfg.Backend {
	case SQLite:
		log.Println("Using SQLite")
		if cfg.Database == "" {
			cfg.Database = "first-try"
		}
		dsn := fmt.Sprintf("./%s.db", cfg.Database)
		db, err = OpenSQLite(dsn)
	case Postgres:
		err = fmt.Errorf("Postgres not supported yet")
	}
	if err != nil {
		return nil, err
	}

	_, err = db.ExecContext(ctx, schema)
	return &UserStore{db: db}, err
}


// rows.Close(), .Next()
// maunually scan and map to struct, later try GORM
// https://github.com/steveyegge/beads/blob/bcfaed92f67238b9f4844445dca8b9fcb7abeaf3/internal/storage/dolt/store.go#L65-L66

// Create table reference: https://github.com/steveyegge/beads/blob/bcfaed92f67238b9f4844445dca8b9fcb7abeaf3/internal/storage/dolt/store.go#L583-L584

/*

1. Embed .sql file to a variable in string
2. func initSchemaOnDB(ctx context.Context, db *sql.DB) error
	a. make sure the file in in the same package, (so you don't have to pass it around)o
	b. run the migrations
3. New invokes createSchema - but see how New is called and managed (notice the package name too)

Extras:
- backoff
- steal this: withRetry
- continue reading: https://github.dev/steveyegge/beads/tree/main

*/
