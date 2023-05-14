create table users (
    user_id int PRIMARY KEY,
    username varchar(25) NOT NULL,
    password varchar(30) NOT NULL,
    created_at timestamp NOT NULL,
    deleted_at timestamp NOT NULL
);