create table permissions
(
    permission_id int PRIMARY KEY,
    permission    varchar(30) NOT NULL,
    created_at timestamp NOT NULL,
    deleted_at timestamp NOT NULL
)