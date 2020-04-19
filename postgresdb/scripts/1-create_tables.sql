CREATE TABLE IF NOT EXISTS repositories (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    language VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS user_repository_tags (
    user_id VARCHAR(255) NOT NULL,
    repository_id VARCHAR(255) NOT NULL,
    tags INT ARRAY NOT NULL,
	PRIMARY KEY (user_id, repository_id)
);

CREATE TABLE IF NOT EXISTS tags (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR (255) NOT NULL
);