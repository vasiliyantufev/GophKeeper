--  card_data - формат json, зашифрованный
create table card
(
    card_id     serial PRIMARY KEY,
    user_id     int       NOT NULL references users (user_id) on delete cascade,
    metadata_id int       NOT NULL,
    card_data   bytea     NOT NULL,
    created_at  timestamp NOT NULL,
    updated_at  timestamp NOT NULL,
    deleted_at  timestamp NULL
)