create table binary_data
(
    binary_id    serial PRIMARY KEY,
    user_id    int          NOT NULL references users (user_id) on delete cascade,
    name       varchar(100) NOT NULL,
    created_at timestamp    NOT NULL,
    updated_at timestamp    NOT NULL,
    deleted_at timestamp NULL
);