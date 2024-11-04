create table versions (
  id serial primary key,
  major int not null,
  minor int not null,
  patch int not null
);

create table test_table (
  id serial primary key,
  name text not null
);
