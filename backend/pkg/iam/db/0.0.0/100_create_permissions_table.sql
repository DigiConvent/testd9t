create table permissions (
   name varchar primary key not null,
   description varchar default '',
   generated boolean default false,
   archived boolean default false,
   meta varchar default ''
);