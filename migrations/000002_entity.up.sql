--  data - зашифрованный
create table entity
(
    entity_id     serial PRIMARY KEY,
    user_id     int       NOT NULL references users (user_id) on delete cascade,
    data        bytea     NOT NULL,
    metadata    jsonb not null,
    created_at  timestamp NOT NULL,
    updated_at  timestamp NOT NULL,
    deleted_at  timestamp NULL
);