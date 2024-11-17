-- name: GetUserById :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    first_name, last_name, email, picture, auth_method
) VALUES (
    ?, ?, ?, ?, ?
) returning id;

-- name: GetFirstNameById :one
SELECT first_name FROM users
WHERE id = ? LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;

-- name: CreateGoogleUser :exec
INSERT INTO google_authed_users (
    google_id, user_id
) VALUES (
    ?, ?
);

-- name: GetUserFromGoogleId :one
SELECT users.id, users.first_name, users.last_name, users.email, users.picture, users.auth_method FROM users
INNER JOIN google_authed_users ON google_authed_users.user_id = users.id AND google_id = ?;

