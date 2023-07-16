create table users
(
    user_id    serial PRIMARY KEY,
    username   varchar(255) not null,
    password   text not null,
    created_at timestamp   NOT NULL,
    updated_at  timestamp  NULL,
    deleted_at timestamp NULL
);

CREATE UNIQUE INDEX idx_username_unique ON users (username);
CREATE UNIQUE INDEX idx_users_deleted_at ON users (deleted_at);