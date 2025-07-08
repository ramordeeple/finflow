CREATE TABLE currencies (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL UNIQUE
);

CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE entries (
    id SERIAL PRIMARY KEY,
    account_id INT REFERENCES accounts(id),
    amount_minor BIGINT NOT NULL,
    currency_id INT REFERENCES currencies(id),
    is_debit BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);
