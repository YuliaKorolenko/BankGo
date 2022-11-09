CREATE TABLE balances
(
    id      serial primary key,
    balance int not null
);

CREATE TABLE transactions
(
    id_order   serial primary key,
    id_user    int references balances (id) on delete cascade not null,
    id_service int,
    amount     int                                            not null,
    is_debit   bool                                           not null,
    cur_time      timestamp without time zone DEFAULT now()
);

CREATE TABLE charges
(
    id_user int references balances (id) on delete cascade not null,
    amount  int,
    cur_time   timestamp without time zone DEFAULT now()
);