-- +migrate Up
ALTER TABLE users
    ADD constraint email_unique
    ADD unique (email);
-- +migrate Down
ALTER TABLE users DROP CONSTRAINT email_unique;