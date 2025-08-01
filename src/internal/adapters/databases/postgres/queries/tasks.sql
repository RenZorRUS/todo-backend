-- name: GetTask :one
SELECT
    t.id,
    t.user_id,
    t.title,
    t.description,
    t.due_date,
    t.status,
    t.created_at,
    t.updated_at,
    t.deleted_at
FROM tasks AS t
WHERE (sqlc.narg(id)::BIGINT IS NULL OR t.id = sqlc.narg(id)::BIGINT)
    AND (sqlc.narg(user_id)::BIGINT IS NULL OR t.user_id = sqlc.narg(user_id)::BIGINT)
    AND (sqlc.narg(title)::VARCHAR(255) IS NULL OR t.title = sqlc.narg(title)::VARCHAR(255))
    AND (sqlc.narg(description)::VARCHAR(1000) IS NULL OR t.description = sqlc.narg(description)::VARCHAR(1000))
    AND (sqlc.narg(due_date)::TIMESTAMPTZ IS NULL OR t.due_date = sqlc.narg(due_date)::TIMESTAMPTZ)
    AND (sqlc.narg(status)::SMALLINT IS NULL OR t.status = sqlc.narg(status)::SMALLINT)
    AND (@show_deleted::BOOLEAN = TRUE OR t.deleted_at IS NULL)
LIMIT 1;

-- name: GetTasks :many
SELECT
    t.id,
    t.user_id,
    t.title,
    t.description,
    t.due_date,
    t.status,
    t.created_at,
    t.updated_at,
    t.deleted_at
FROM tasks AS t
WHERE (sqlc.narg(user_id)::BIGINT IS NULL OR t.user_id = sqlc.narg(user_id)::BIGINT)
    AND (sqlc.narg(title)::VARCHAR(255) IS NULL OR t.title = sqlc.narg(title)::VARCHAR(255))
    AND (sqlc.narg(description)::VARCHAR(1000) IS NULL OR t.description = sqlc.narg(description)::VARCHAR(1000))
    AND (sqlc.narg(due_date)::TIMESTAMPTZ IS NULL OR t.due_date = sqlc.narg(due_date)::TIMESTAMPTZ)
    AND (sqlc.narg(status)::SMALLINT IS NULL OR t.status = sqlc.narg(status)::SMALLINT)
    AND (@show_deleted::BOOLEAN = TRUE OR t.deleted_at IS NULL)
ORDER BY t.id ASC
LIMIT @rows_limit::BIGINT
OFFSET @rows_offset::BIGINT;

-- name: CreateTask :one
INSERT INTO tasks (
    user_id,
    title,
    description,
    due_date,
    status
)
VALUES (
    @user_id::BIGINT,
    @title::VARCHAR(255),
    @description::VARCHAR(1000),
    @due_date::TIMESTAMPTZ,
    @status::SMALLINT
)
RETURNING *;

-- name: UpdateTask :one
UPDATE tasks
SET
    user_id = @user_id::BIGINT,
    title = @title::VARCHAR(255),
    description = @description::VARCHAR(1000),
    due_date = @due_date::TIMESTAMPTZ,
    status = @status::SMALLINT
WHERE id = @id::BIGINT
RETURNING *;

-- name: SoftDeleteTask :exec
UPDATE tasks
SET
    deleted_at = CLOCK_TIMESTAMP()
WHERE id = @id::BIGINT
    AND user_id = @user_id::BIGINT
    AND deleted_at IS NULL;

-- name: HardDeleteTask :exec
DELETE FROM tasks AS t
WHERE t.id = @id::BIGINT
    AND t.user_id = @user_id::BIGINT
    AND t.deleted_at IS NOT NULL;
