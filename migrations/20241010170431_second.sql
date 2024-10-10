-- +goose Up
-- +goose StatementBegin

alter table users add column last_name varchar(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN last_name;
-- +goose StatementEnd
