create table permission_groups (
  id uuid primary key not null,
  name varchar not null,
  abbr varchar default '',
  is_group boolean default false,
  -- is_node describes whether the group is a leaf node in the permission tree and cannot have subgroups
  is_node boolean default false,
  description varchar default '',
  parent uuid references permission_groups(id) on delete set null,
  "generated" boolean default false
);