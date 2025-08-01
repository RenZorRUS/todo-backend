-- name: GetUser :one
SELECT
    u.id,
    u.name,
    u.password_hash,
    u.email,
    u.created_at,
    u.updated_at
FROM users AS u
WHERE (sqlc.narg(id)::BIGINT IS NULL OR u.id = sqlc.narg(id)::BIGINT)
    AND (sqlc.narg(name)::VARCHAR(255) IS NULL OR u.name = sqlc.narg(name)::VARCHAR(255))
    AND (sqlc.narg(email)::VARCHAR(255) IS NULL OR u.email = sqlc.narg(email)::VARCHAR(255))
LIMIT 1;

-- name: GetUsers :many
SELECT
    u.id,
    u.name,
    u.password_hash,
    u.email,
    u.created_at,
    u.updated_at
FROM users AS u
ORDER BY u.id ASC
LIMIT @rows_limit::BIGINT
OFFSET @rows_offset::BIGINT;

-- name: CreateUser :one
INSERT INTO users (
    name,
    password_hash,
    email
)
VALUES (
    @name::VARCHAR(255),
    @password_hash::VARCHAR(255),
    @email::VARCHAR(255)
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET
    name = @name::VARCHAR(255),
    email = @email::VARCHAR(255),
    password_hash = @password_hash::VARCHAR(255)
WHERE id = @id::BIGINT
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users AS u
WHERE (sqlc.narg(id)::BIGINT IS NULL OR u.id = sqlc.narg(id)::BIGINT)
    AND (sqlc.narg(name)::VARCHAR(255) IS NULL OR u.name = sqlc.narg(name)::VARCHAR(255))
    AND (sqlc.narg(email)::VARCHAR(255) IS NULL OR u.email = sqlc.narg(email)::VARCHAR(255));
