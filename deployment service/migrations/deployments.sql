CREATE TABLE deployments
(
    id BIGSERIAL PRIMARY KEY,

    app_name VARCHAR(255),

    image VARCHAR(255),

    replicas INTEGER,

    status VARCHAR(50),

    created_at TIMESTAMP
    DEFAULT NOW()
);
