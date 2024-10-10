-- +goose Up
-- +goose StatementBegin
create table users(
    id serial,
    name varchar
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
