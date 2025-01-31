create table if not exists permissions (
  id uuid primary key not null,
  name varchar unique not null,
  description varchar default '',
  generated boolean default false,
  archived boolean default false,
  meta varchar default ''
);