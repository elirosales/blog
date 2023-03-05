CREATE TABLE IF NOT EXISTS articles(
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(50) NOT NULL,
    title VARCHAR(50) NOT NULL,
    content TEXT,
    author VARCHAR(50) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    updated_at timestamp NULL DEFAULT NULL,
    deleted_at timestamp NULL DEFAULT NULL
)
