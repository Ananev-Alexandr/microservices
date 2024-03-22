-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE chat (
    id SERIAL PRIMARY KEY,
    from_ VARCHAR(255) NOT NULL,
    text_ VARCHAR(255),
    role INT NOT NULL,
    timestamp_ TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
-- SQL in section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS chat;
