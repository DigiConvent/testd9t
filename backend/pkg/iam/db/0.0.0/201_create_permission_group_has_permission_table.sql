create table permission_group_has_permission (
   permission_group uuid not null references permission_groups(id) on delete cascade,
   permission varchar not null references permissions("name") on delete cascade,
   primary key (permission_group, permission)
);