-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id              BIGSERIAL PRIMARY KEY,
    name            VARCHAR(255) NOT NULL UNIQUE,
    password_hash   VARCHAR(255) NOT NULL,
    email           VARCHAR(255) NOT NULL UNIQUE,
    created_at      TIMESTAMP NOT NULL DEFAULT CLOCK_TIMESTAMP(),
    updated_at      TIMESTAMP NOT NULL DEFAULT CLOCK_TIMESTAMP()
);
CREATE TABLE IF NOT EXISTS tasks (
    id              BIGSERIAL PRIMARY KEY,
    title           VARCHAR(255) NOT NULL,
    description     VARCHAR(1000),
    user_id         BIGINT NOT NULL,
    status          SMALLINT NOT NULL CHECK (status IN (0, 1, 2)) DEFAULT 0,
    due_date        TIMESTAMPTZ,
    created_at      TIMESTAMP NOT NULL DEFAULT CLOCK_TIMESTAMP(),
    updated_at      TIMESTAMP NOT NULL DEFAULT CLOCK_TIMESTAMP(),
    deleted_at      TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = clock_timestamp();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_modified_time_users_tr
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();
CREATE TRIGGER update_modified_time_tasks_tr
    BEFORE UPDATE ON tasks
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_modified_time_tasks_tr ON tasks;
DROP TRIGGER IF EXISTS update_modified_time_users_tr ON users;

DROP FUNCTION IF EXISTS update_modified_column;

DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
