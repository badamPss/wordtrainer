CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     telegram_id TEXT UNIQUE,
                                     username TEXT NOT NULL UNIQUE,
                                     password_hash TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS categories (
                                          id SERIAL PRIMARY KEY,
                                          user_id INT REFERENCES users(id) ON DELETE CASCADE,
                                          name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS cards (
                                     id SERIAL PRIMARY KEY,
                                     user_id INT REFERENCES users(id) ON DELETE CASCADE,
                                     category_id INT REFERENCES categories(id) ON DELETE CASCADE,
                                     word TEXT NOT NULL,
                                     translation TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS attempts (
                                        id SERIAL PRIMARY KEY,
                                        user_id INT REFERENCES users(id) ON DELETE CASCADE,
                                        card_id INT REFERENCES cards(id) ON DELETE CASCADE,
                                        correct BOOLEAN NOT NULL
);