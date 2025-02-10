-- backend/pkg/post/db/0.0.0/00_create_email_addresses_table.sql 
create table email_addresses (
    id uuid primary key not null,
    name varchar not null unique,
    domain varchar not null
);

