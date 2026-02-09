package chat

import (
	"chat/domain"
	"context"
	"database/sql"
	"log"
)

type UserStore struct {
	sql *sql.DB
}

func New(db *sql.DB) *UserStore {
	// get cfg in argument (not db)
	// func New(ctx context.Context, cfg *Config) (*DoltStore, error) {
	// create the database if does not exist (on the right path)
	// construct dsn
	return &UserStore{
		sql: db,
	}
}

func (s *UserStore) Insert(ctx context.Context, user *domain.User) (*domain.User, error) {
	s.sql.ExecContext(ctx, "INSERT INTO users (username, alias, bio) VALUES (?, ?, ?)",
		user.Username, user.Alias, user.Bio)

	return nil, nil
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
