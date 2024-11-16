-- name: GetUserById :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;

-- name: CreateUser :exec
INSERT INTO users (
  first_name, last_name, email, picture
) VALUES (
  ?, ?, ?, ?
);

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;
