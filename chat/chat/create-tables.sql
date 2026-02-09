-- reference: https://github.com/steveyegge/beads/blob/bcfaed92f67238b9f4844445dca8b9fcb7abeaf3/examples/bd-example-extension-go/main.go#L17-L18
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    alias TEXT,
    bio TEXT
);
