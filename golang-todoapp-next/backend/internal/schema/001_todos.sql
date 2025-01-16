-- +goose Up
CREATE TABLE todos (
  id   BIGSERIAL PRIMARY KEY,
  title text      NOT NULL,
  content text NOT NULL,
  createdAt timestamp NOT NULL
);

-- +goose Down
DROP TABLE todos;