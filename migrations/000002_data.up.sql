create table data (
    data_id int PRIMARY KEY,
    user_id int NOT NULL references users (user_id) on delete cascade,
    data varchar(100) NOT NULL,
    type varchar(30) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NULL,
    deleted_at timestamp NULL
)