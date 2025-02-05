create table permission_group_has_user (
  permission_group uuid not null references permission_groups(id) on delete cascade,
  "user" uuid not null references users(id) on delete cascade,
  "start" timestamp not null default CURRENT_TIMESTAMP,
  "end" timestamp default null,
  primary key (permission_group, "user")
);