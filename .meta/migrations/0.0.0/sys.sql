-- backend/pkg/sys/db/0.0.0/00_create_versions_table.sql 
create table versions (
    major integer not null,
    minor integer not null,
    patch integer not null,
    migrated_at timestamp not null default current_timestamp,
    primary key (major, minor, patch)
);

-- backend/pkg/sys/db/0.0.0/01_create_packages_table.sql 
create table packages (
    name varchar not null primary key,
    major integer not null,
    minor integer not null,
    patch integer not null
);

-- backend/pkg/sys/db/0.0.0/02_create_configuration_table.sql 
create table configurations (
    domain text default '',
    telegram_bot_token text default ''
);

insert into configurations (domain, telegram_bot_token) values ('', '');

