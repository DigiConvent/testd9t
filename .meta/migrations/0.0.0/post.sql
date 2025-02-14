-- backend/pkg/post/db/0.0.0/00_create_email_addresses_table.sql 
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

-- backend/pkg/post/db/0.0.0/01_create_post_table.sql 
create table emails (
    id uuid primary key not null, -- this indicates a folder that contains the contents of the email
    -- plain content is stored under <DATABASE_PATH>/post/email/<id>/contents
    -- attachments are stored under <DATABASE_PATH>/post/email/<id>/attachments/<filename>
    to_email_address uuid not null references email_addresses(id),
    from_email_address varchar not null,
    subject varchar not null,
    notes varchar default '',
    sent_at timestamp not null default current_timestamp,
    read_at timestamp default null
)

