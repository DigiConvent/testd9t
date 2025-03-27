create table user_status (
  id uuid primary key not null,
  name varchar not null unique,
  abbr varchar not null unique,
  description varchar default '',
  archived boolean default false
);

-- there are triggers for this table in 20_create_triggers_for_user_status.sql
-- where the user_status is assigned a permission_group with the same id thus
-- enforcing a hierarchy