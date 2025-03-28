create table users (
  id uuid primary key not null,
  emailaddress varchar unique default '',
  password varchar default '',
  telegram_id bigint default 0,
  first_name varchar default '',
  last_name varchar default '',
  enabled boolean default true,
  active boolean default false
);

