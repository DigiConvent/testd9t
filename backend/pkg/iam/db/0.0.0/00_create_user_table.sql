create table users (
  id uuid primary key not null,
  emailaddress varchar unique default '',
  password varchar default '',
  telegram_id bigint default 0,
  title varchar default '',
  first_name varchar default '',
  last_name varchar default '',
  date_of_birth date default null,
  enabled boolean default true,
  super boolean default false,
  active boolean default false
);