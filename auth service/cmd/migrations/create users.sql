CREATE TABLE users
(
    id BIGSERIAL PRIMARY KEY,

    email VARCHAR(255)
    UNIQUE NOT NULL,

    password TEXT NOT NULL,

    role VARCHAR(50) NOT NULL,

    created_at TIMESTAMP DEFAULT NOW()
);
