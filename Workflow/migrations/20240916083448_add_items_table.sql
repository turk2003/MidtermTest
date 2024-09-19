-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    amount INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    status VARCHAR(50) NOT NULL,
    owner_id INTEGER NOT NULL

);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE items;
