create table permission_groups (
   id uuid primary key not null,
   name varchar not null,
   abbr varchar default '',
   -- is_group tells whether the group can contain users
   is_group boolean default false,
   meta varchar default null,
   -- is_node describes whether the group is a leaf node in the permission tree and cannot have subgroups
   is_node boolean default false,
   description varchar default '',
   parent uuid references permission_groups(id) on delete set null,
   "generated" boolean default false
);

-- there can only be one root
create unique index one_null on permission_groups(parent) where parent is null;