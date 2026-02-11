-- +goose Up
-- +goose StatementBegin
CREATE TABLE files(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now() at time zone 'utc')
);

CREATE UNIQUE INDEX files_uniq_name_idx ON files USING btree (name);

CREATE TRIGGER update_files_updated_at BEFORE
    UPDATE
    ON files FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
