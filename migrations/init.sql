create table currencies (
    id serial primary key,
    code text not null unique
);

create table accounts (
    id serial primary key,
    name text not null
);

create table accounts (
    id serial primary key,
    name text not null,
);

create table entries (
    id serial primary key,
    account_id integer not null references accounts(id),
    amount_minor bigint not null,
    currency_id integer not null references currencies(id),
    is_debit boolean not null,
    created_at timestamp default now()
);