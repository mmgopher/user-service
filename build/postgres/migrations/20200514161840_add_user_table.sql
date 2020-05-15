-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user_sch"."user" (
     id SERIAL PRIMARY KEY,
     name text NOT NULL,
     surname text NOT NULL,
     gender text,
     age integer,
     address text
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "user_sch"."user";
-- +goose StatementEnd
