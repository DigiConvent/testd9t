create table if not exists versions (
    major integer not null,
    minor integer not null,
    patch integer not null,
    migrated_at timestamp not null default current_timestamp,
    primary key (major, minor, patch)
);