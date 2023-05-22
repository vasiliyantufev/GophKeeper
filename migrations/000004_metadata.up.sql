create table metadata
(
    metadata_id int PRIMARY KEY,
    data_id int NOT NULL references data (data_id) on delete cascade,
    metadata varchar(100) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NULL,
    deleted_at timestamp NULL
)