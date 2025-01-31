create table if not exists versions (
    major integer not null,
    minor integer not null,
    patch integer not null,
    migrated_at timestamp not null default current_timestamp,
    primary key (major, minor, patch)
);create table if not exists packages (
    name varchar not null primary key,
    major integer not null,
    minor integer not null,
    patch integer not null
);create table if not exists configurations (
    telegram_bot_token text default '',
    domain text default ''
);