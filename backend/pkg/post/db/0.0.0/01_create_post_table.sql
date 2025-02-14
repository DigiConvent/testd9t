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