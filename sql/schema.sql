CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL,
    picture TEXT NOT NULL,
    auth_method INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS google_authed_users (
    id INTEGER PRIMARY KEY,
    google_id TEXT UNIQUE NOT NULL,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
