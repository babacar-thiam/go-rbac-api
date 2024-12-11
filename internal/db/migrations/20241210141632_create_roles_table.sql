-- +goose Up
-- +goose StatementBegin
create table if not exists roles (
  id char(36) primary key,
  name varchar(255) unique not null,
  description varchar(255) not null,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table if exists roles;

-- +goose StatementEnd
