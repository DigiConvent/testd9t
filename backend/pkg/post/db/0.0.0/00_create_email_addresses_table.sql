create table email_addresses (
    id uuid primary key not null,
    name varchar not null unique,
    domain varchar not null,
    password varchar not null
);