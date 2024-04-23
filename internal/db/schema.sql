CREATE TABLE profile (
    id SERIAL PRIMARY KEY, 
    kth_id TEXT UNIQUE NOT NULL,
    bank TEXT NOT NULL DEFAULT '',
    bank_account_number TEXT NOT NULL DEFAULT '',
    clearing_number TEXT NOT NULL DEFAULT ''
);

CREATE TABLE payment (
    id SERIAL PRIMARY KEY,
    paid_at TIMESTAMP NOT NULL,
    paid_by INTEGER REFERENCES profile NOT NULL,
    paid_to INTEGER REFERENCES profile NOT NULL
);

CREATE TYPE expense_type AS ENUM ('reimbursable', 'invoice', 'card');

CREATE TABLE expense (
    id SERIAL PRIMARY KEY,
    kind expense_type NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    is_digital BOOLEAN NOT NULL,
    accounting_date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    confirmed_at TIMESTAMP,
    verification_nr TEXT,
    created_by INTEGER REFERENCES profile NOT NULL,
    confirmed_by INTEGER REFERENCES profile,
    payment_id INTEGER REFERENCES payment
);

CREATE TABLE expense_invoice (
    id INTEGER PRIMARY KEY REFERENCES expense,
    due_date DATE NOT NULL
);

CREATE TABLE expense_part (
    id SERIAL PRIMARY KEY,
    expense_id INTEGER REFERENCES expense NOT NULL,
    ammount DECIMAL NOT NULL,
    comitte TEXT NOT NULL,
    budget_line TEXT NOT NULL,
    cost_centre TEXT NOT NULL,
    attested_at TIMESTAMP,
    attested_by INTEGER REFERENCES profile
);

CREATE TABLE discussion_post (
    id SERIAL PRIMARY KEY,
    posted_at TIMESTAMP NOT NULL,
    content TEXT NOT NULL,
    expense_id INTEGER REFERENCES expense NOT NULL,
    author_id INTEGER REFERENCES profile NOT NULL
);

CREATE TABLE blob_file (
    id SERIAL PRIMARY KEY,
    expense_id INTEGER REFERENCES expense NOT NULL,
    file_name TEXT NOT NULL
);
