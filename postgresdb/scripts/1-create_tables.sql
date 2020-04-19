CREATE TABLE IF NOT EXISTS authors (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR (255) NOT NULL
);

CREATE TABLE IF NOT EXISTS books (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR (255) NOT NULL,
    edition INT NOT NULL CHECK (edition > 0),
    publication_year INT NOT NULL CHECK (publication_year > 0),
    authors INT ARRAY NOT NULL
);