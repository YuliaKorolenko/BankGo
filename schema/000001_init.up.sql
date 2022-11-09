CREATE TABLE balances
(
    id      serial not null unique,
    balance int    not null
);

CREATE TABLE transactions
(
    id_order   serial not null unique,
    id_user    int references balances (id) on delete cascade not null,
    id_service int,
    amount     int                                            not null,
    flag       int                                            not null
);