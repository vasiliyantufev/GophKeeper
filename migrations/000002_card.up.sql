create table card
(
    card_id             int PRIMARY KEY,
    user_id             int          NOT NULL references users (user_id) on delete cascade,
    card_payment_system varchar(30)  NOT NULL,
    card_number         varchar(50)  NOT NULL,
    card_holder         varchar(100) NOT NULL,
    card_end_date       timestamp    NOT NULL,
    cvc                 int          NOT NULL,
    created_at          timestamp    NOT NULL,
    updated_at          timestamp    NOT NULL,
    deleted_at          timestamp NULL
)