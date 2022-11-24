CREATE TABLE users(
    id INTEGER PRIMARY KEY
);

CREATE TABLE balances(
    id SERIAL PRIMARY KEY,
    user_id INTEGER UNIQUE REFERENCES users(id),
    sum NUMERIC(18, 2)
);

CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    service_id INTEGER NOT NULL,
    price NUMERIC(18, 2),
    description VARCHAR(256)
);

INSERT INTO users(id) VALUES(0);
INSERT INTO users(id) VALUES(1);
INSERT INTO users(id) VALUES(2);