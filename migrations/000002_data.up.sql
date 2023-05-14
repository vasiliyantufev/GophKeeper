create table data (
    data_id int PRIMARY KEY,
    user_id int NOT NULL,
    data varchar(100) NOT NULL,
    type varchar(30) NOT NULL,
    created_at timestamp NOT NULL,
    deleted_at timestamp NOT NULL
)