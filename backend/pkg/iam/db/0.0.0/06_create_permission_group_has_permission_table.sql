create table permission_group_has_permission (
  permission_group uuid not null references permission_groups(id) on delete cascade,
  permission uuid not null references permissions(id) on delete cascade,
  primary key (permission_group, permission)
);