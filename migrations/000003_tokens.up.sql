create table tokens
(
    token_id int PRIMARY KEY,
    user_id int NOT NULL references users (user_id) on delete cascade,
    token    varchar(100) NOT NULL,
    created_at timestamp NOT NULL,
    deleted_at timestamp NOT NULL
)