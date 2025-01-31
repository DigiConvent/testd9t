create table if not exists user_status (
  id uuid primary key not null,
  name varchar not null unique,
  abbr varchar not null unique,
  description varchar default '',
  archived boolean default false
);