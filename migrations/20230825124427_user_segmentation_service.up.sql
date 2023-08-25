CREATE TABLE users (
    user_id VARCHAR(40) PRIMARY KEY,
    created_at TIMESTAMP not null default now()
);

CREATE TABLE segments (
    slug VARCHAR PRIMARY KEY,
    created_at TIMESTAMP not null default now()
);

CREATE TABLE user_segments (
    user_id VARCHAR(40) REFERENCES users(user_id),
    slug VARCHAR REFERENCES segments(slug) ON DELETE CASCADE,
    PRIMARY KEY (user_id, slug)
);

CREATE TABLE types_transaction (
    operation_name VARCHAR PRIMARY KEY,
    description TEXT
);

CREATE TABLE history (
    history_id SERIAL PRIMARY KEY,
    user_id VARCHAR(40) REFERENCES users(user_id),
    segment_slug VARCHAR REFERENCES segments(slug),
    operation_name VARCHAR REFERENCES types_transaction(operation_name),
    created_at TIMESTAMP not null default now()
);