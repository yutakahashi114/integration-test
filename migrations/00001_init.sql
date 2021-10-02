
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users
(
  id         bigserial  NOT NULL,
  name       varchar    NOT NULL,
  email      varchar    NOT NULL,
  created_at timestamp  NOT NULL DEFAULT now(),
  updated_at timestamp  NOT NULL DEFAULT now(),
  deleted_at timestamp,
  PRIMARY KEY (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS users;
