CREATE TABLE transactions
(
    id       serial       not null unique,
    user_id  int          not null,
    email    varchar(255) not null,
    amount   int          not null,
    currency varchar(255) not null,
    created_at date not null,
    updated_at date not null,
    status varchar(255) not null
);