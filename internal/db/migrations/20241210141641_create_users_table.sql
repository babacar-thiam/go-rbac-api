-- +goose Up
-- +goose StatementBegin
create table if not exists users (
  id char(36) primary key,
  role_id char(36) not null,
  full_name varchar(255) not null,
  email_address varchar(50) unique not null,
  phone_number varchar(20) unique not null,
  password_hash varchar(50) not null,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table if exists users;

-- +goose StatementEnd
