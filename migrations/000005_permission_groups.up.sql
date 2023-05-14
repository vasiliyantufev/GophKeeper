create table permission_groups
(
    permission_groups_id int PRIMARY KEY,
    permission_id    int NOT NULL,
    user_id    int NOT NULL,
    created_at timestamp NOT NULL,
    deleted_at timestamp NOT NULL
)