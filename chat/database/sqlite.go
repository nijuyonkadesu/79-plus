package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

/*
side-effect import... hmm...
~/go/pkg/mod/modernc.org/sqlite@v1.44.3

  1 const (
51      driverName              = "sqlite"
  1     ptrSize                 = unsafe.Sizeof(uintptr(0))
  2     sqliteLockedSharedcache = sqlite3.SQLITE_LOCKED | (1 << 8)
  3 )
  4
  5 func init() {
  6     sql.Register(driverName, newDriver())
  7     sqlite3.PatchIssue199() // https://gitlab.com/cznic/sqlite/-/issues/199
  8
  9 }

1. The Import: When your program starts, it sees _ "modernc.org/sqlite". It runs the init() function inside that library.
2. The Registration: That library calls sql.Register("sqlite", ...). Now, the standard database/sql package has a "map" that says: *"If anyone asks for 'sqlite', I should use this specific code."*
3. The Call: When you call sql.Open("sqlite", dsn), the database/sql package looks at its map, finds the driver registered under the name "sqlite", and starts using it.

TODO: Generics shine in algorithms? and less useful in data access layers? (what does it mean), gotta see an example
*/

func OpenSQLite(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping sqlite: %w", err)
	}

	return db, nil
}

/*

1. open any sql backend (in this case sqlite of type *sql.DB)
2. pass it to SQLiteRepository, and if it had all methods defined in Repository, I can have my service depend on Repository type and I can pass around SQLiteRepository
3. the reason SQLiteRepository is in users is it is heavily tied with User object? (coz repositories are domain-specific & SQL should live close to the domain?)
4. all above is true to “Accept interfaces, return concrete types.”

cohesion beats reuse

UserStore is a simpler version of Repository pattern, but it usually deals with a single table.

- *sql.DB → infrastructure
- SQLiteUserStore → persistence detail
- UserStore / UserRepository → contract
- Service → business logic
- main → wiring

TODO: what is http.RoundTripper?
*/
