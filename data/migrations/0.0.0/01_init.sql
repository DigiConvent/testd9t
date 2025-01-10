-- doesn't work? -> grant create on database digiconvent to <role>;
create extension pgcrypto;
create extension "uuid-ossp";
SET client_encoding = 'UTF8';

create table versions (
  major smallint not null,
  minor smallint not null,
  patch smallint not null,
  description text,
  migrated bool not null default false,
  primary key (major, minor, patch)
);

create table config (
  telegram_bot_token varchar not null default '',
  domain varchar not null default ''
);

insert into config (telegram_bot_token) values ('');