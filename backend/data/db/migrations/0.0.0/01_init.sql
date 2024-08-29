create table if not exists versions (
  id serial primary key,
  major int not null,
  minor int not null,
  patch int not null
);

create table if not exists test_table (
  id serial primary key,
  name text not null
);