-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user_sch"."user" (
     id SERIAL PRIMARY KEY,
     name text NOT NULL,
     surname text NOT NULL,
     gender text NOT NULL,
     age integer NOT NULL,
     address text NOT NULL,
     created_at timestamp with time zone DEFAULT now()
);
ALTER SEQUENCE user_sch.user_id_seq RESTART WITH 500;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "user_sch"."user";
-- +goose StatementEnd
