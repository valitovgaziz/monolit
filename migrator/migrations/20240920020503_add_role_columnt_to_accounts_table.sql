-- +goose Up
-- +goose StatementBegin
ALTER TABLE accounts ADD COLUMN role VARCHAR(50) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE accounts DELETE COLUMN role;
-- +goose StatementEnd
