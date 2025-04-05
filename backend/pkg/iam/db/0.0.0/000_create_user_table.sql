create table users (
   id uuid primary key not null,
   telegram_id bigint default 0,
   emailaddress varchar unique default '',
   password varchar default '',
   first_name varchar default '',
   last_name varchar default '',
   enabled boolean default true
);