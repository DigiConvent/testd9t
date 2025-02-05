create table packages (
    name varchar not null primary key,
    major integer not null,
    minor integer not null,
    patch integer not null
);