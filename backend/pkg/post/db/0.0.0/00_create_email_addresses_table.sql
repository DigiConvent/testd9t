create table email_addresses (
    id uuid primary key not null,
    name varchar not null unique,
    domain varchar not null,
    generated boolean not null default 0
);

insert into email_addresses (
    id,
    name,
    domain
) values (
    '00000000-0000-0000-0000-000000000000',
    'admin',
    ''
);