-- name: GetUser :one
SELECT * FROM users
WHERE id = ? AND status = 'ACTIVE' LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email_address = ? AND status = 'ACTIVE' LIMIT 1;

-- name: GetActiveKeycloakUser :one
SELECT * FROM users
WHERE keycloak_uid = ? AND status = 'ACTIVE' LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
WHERE status ='ACTIVE'
ORDER BY created_at
LIMIT ? OFFSET ?;

-- name: CreateUser :execresult
INSERT INTO users (
  id, first_name, last_name, email_address, role, keycloak_uid
) VALUES (
  ?, ?, ?, ?, ?, ?
);

-- name: DeleteUser :exec
UPDATE users
SET status = 'DELETED'
WHERE id = ?;