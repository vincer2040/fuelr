CREATE TABLE IF NOT EXISTS users(
    ID INTEGER PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL,
    auth_method INTEGER NOT NULL
);
