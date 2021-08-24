-- +goose Up
-- +goose StatementBegin
CREATE TABLE "roadmap"
(
    "id"         SERIAL PRIMARY KEY,
    "user_id"    INT          NOT NULL,
    "link"       VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMPTZ  NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE roadmap;
-- +goose StatementEnd
