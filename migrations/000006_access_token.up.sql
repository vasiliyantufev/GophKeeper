create table access_token
(
    access_token     varchar(30) NOT NULL,
    user_id          int         NOT NULL references users (user_id) on delete cascade,
    created_at       timestamp   NOT NULL,
    end_date_at      timestamp   NOT NULL
)

-- сделать проверку сессионного ключа, они должны быть уникальными