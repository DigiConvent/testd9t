create table if not exists users (
  id uuid primary key not null,
  email varchar unique default '',
  password varchar default '',
  telegram_id bigint default 0,
  title varchar default '',
  first_name varchar default '',
  last_name varchar default '',
  date_of_birth date default null,
  enabled boolean default true,
  super boolean default false,
  active boolean default false
);create table if not exists user_status (
  id uuid primary key not null,
  name varchar not null unique,
  abbr varchar not null unique,
  description varchar default '',
  archived boolean default false
);create table if not exists user_became_status (
  "user" uuid not null references users(id) on delete cascade,
  "status" uuid not null references user_status(id) on delete cascade,
  "date" timestamp not null,
  active boolean default false,
  description varchar default '',
  primary key ("user", "status")
);create view user_facades as
select distinct u.id, u.id, concat(us.abbr, ' ', u.last_name) as name, ubs.date
from users u
left join user_became_status ubs on u.id = ubs.user
left join user_status us on us.id = ubs.status
order by u.id, ubs.date desc;create table if not exists permissions (
  id uuid primary key not null,
  name varchar unique not null,
  description varchar default '',
  generated boolean default false,
  archived boolean default false,
  meta varchar default ''
);create table if not exists permission_groups (
  id uuid primary key not null,
  name varchar unique not null,
  abbr varchar default '',
  is_group boolean default false,
  -- is_node describes whether the group is a leaf node in the permission tree and cannot have subgroups
  is_node boolean default false,
  description varchar default '',
  parent uuid references permission_groups(id) on delete set null,
  "generated" boolean default false
);create table if not exists permission_group_has_permission (
  permission_group uuid not null references permission_groups(id) on delete cascade,
  permission uuid not null references permissions(id) on delete cascade,
  primary key (permission_group, permission)
);create table if not exists permission_group_has_user (
  permission_group uuid not null references permission_groups(id) on delete cascade,
  "user" uuid not null references users(id) on delete cascade,
  "start" timestamp not null default CURRENT_TIMESTAMP,
  "end" timestamp default null,
  primary key (permission_group, "user")
);create view permission_group_has_users as
with recursive relevant_groups as (
  select 
    pg.id as "root",
    pg.id as id,
    pg.name,
    0 as descendant,
    pg.parent
  from permission_groups pg
  union all
  select 
    s."root",
    child.id as id,
    child.name,
    1 as descendant,
    child.parent
  from permission_groups child
  inner join relevant_groups s on s.id = child.parent
)
select rg.root, rg.descendant as implied, rg.id as permission_group, uf.id as "user", uf.name 
from relevant_groups rg
join permission_group_has_user pghu on pghu.permission_group = rg.id 
and pghu.start <= datetime('now') 
and (pghu.end is null or pghu.end >= datetime('now'))
join user_facades uf on uf.id = pghu."user";
create view user_has_permission_groups as
with recursive relevant_groups as (select 
    pghu."user", 
    pghu.permission_group, 
    0 as implied,
    pg.parent
  from permission_group_has_user pghu
  join permission_groups pg on pghu.permission_group = pg.id
  where pghu.start <= now() and (pghu.end is null or pghu.end >= now())
  union all
  select 
    s.user,
    child.id as permission_group,
    1 as implied,
    child.parent
  from permission_groups child
  inner join relevant_groups s on s.parent = child.id)
select * from relevant_groups;create view permission_has_permission_groups as
with recursive relevant_groups as (select 
    pghp.permission, 
    pghp.permission_group, 
    0 as implied,
    pg.parent
  from permission_group_has_permission pghp
  join permission_groups pg on pghp.permission_group = pg.id
  union all
  select 
    s.permission,
    child.id as permission_group,
    1 as implied,
    child.parent
  from permission_groups child
  inner join relevant_groups s on child.parent = s.permission_group)
select * from relevant_groups;