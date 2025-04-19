create table permission_groups (
   id uuid primary key not null,
   name varchar not null,
   abbr varchar default '',
   meta varchar default null,
   description varchar default '',
   parent uuid references permission_groups(id) on delete set null,
   "generated" boolean default false
);

-- there can only be one root
create unique index one_null on permission_groups(parent) where parent is null;