-- backend/pkg/iam/db/0.0.0/00_create_user_table.sql 
create table users (
  id uuid primary key not null,
  emailaddress varchar unique default '',
  password varchar default '',
  telegram_id bigint default 0,
  first_name varchar default '',
  last_name varchar default '',
  enabled boolean default true,
  active boolean default false
);



-- backend/pkg/iam/db/0.0.0/01_create_user_status_table.sql 
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

-- backend/pkg/iam/db/0.0.0/02_create_user_became_status_table.sql 
create table user_became_status (
  "user" uuid not null references users(id) on delete cascade,
  "status" uuid not null references user_status(id) on delete cascade,
  "start" timestamp not null,
  description varchar default '',
  primary key ("user", "status")
);

create view user_status_interval as
select 
    status,
    start,
    coalesce(
        datetime(lead(start) over (partition by user order by start), '-1 second'),
        null
    ) as end
from user_became_status;

-- backend/pkg/iam/db/0.0.0/03_create_user_facades_view.sql 
create view user_facades as
select 
    distinct u.id, 
    u.id, 
    concat(us.abbr, ' ', u.last_name) as name, 
    us.id as status_id, 
    us.name as status_name
from users u
left join user_became_status ubs on u.id = ubs.user
left join user_status us on us.id = ubs.status
-- learned the hard way that 'or' != '||'
where ubs.start <= current_timestamp or ubs.start is null
order by u.id, ubs.start desc;

-- backend/pkg/iam/db/0.0.0/04_create_permissions_table.sql 
create table permissions (
  name varchar primary key not null,
  description varchar default '',
  generated boolean default false,
  archived boolean default false,
  meta varchar default ''
);

-- backend/pkg/iam/db/0.0.0/05_create_permission_groups_table.sql 
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

create unique index one_null on permission_groups(parent) where parent is null;

-- backend/pkg/iam/db/0.0.0/06_create_permission_group_has_permission_table.sql 
create table permission_group_has_permission (
  permission_group uuid not null references permission_groups(id) on delete cascade,
  permission varchar not null references permissions("name") on delete cascade,
  primary key (permission_group, permission)
);

-- backend/pkg/iam/db/0.0.0/07_create_permission_group_has_user_table.sql 
create table permission_group_has_user (
  permission_group uuid not null references permission_groups(id) on delete cascade,
  "user" uuid not null references users(id) on delete cascade,
  "start" timestamp not null default CURRENT_TIMESTAMP,
  "end" timestamp default null,
  primary key (permission_group, "user")
);

-- backend/pkg/iam/db/0.0.0/08_create_permission_group_has_users_view.sql 
create view permission_group_has_users as
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
and (pghu.start <= datetime('now') or pghu.start is null)
and (pghu.end is null or datetime('now') < pghu.end)
join user_facades uf on uf.id = pghu."user";


-- backend/pkg/iam/db/0.0.0/09_create_user_has_permission_groups_view.sql 
create view user_has_permission_groups as
with recursive relevant_groups as (select 
    pghu."user", 
    pghu.permission_group, 
    0 as implied,
    pg.parent
  from permission_group_has_user pghu
  join permission_groups pg on pghu.permission_group = pg.id
  where (pghu.start <= datetime('now') or pghu.start is null) 
  and (pghu.end is null or datetime('now') < pghu.end)
  union all
  select 
    s.user,
    child.id as permission_group,
    1 as implied,
    child.parent
  from permission_groups child
  inner join relevant_groups s on s.parent = child.id)
select * from relevant_groups;

-- backend/pkg/iam/db/0.0.0/10_create_permission_has_permission_groups_view.sql 
create view permission_has_permission_groups as
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

-- backend/pkg/iam/db/0.0.0/11_create_permission_group_has_permission_groups_view.sql 
create view permission_group_has_permission_group_ancestors as
with recursive ancestors as (select
    pg.id,
    pg.name, 
    0 as implied,
    pg.parent,
    pg.id as child_id
  from permission_groups pg
  union all
  select 
    parent.id,
    parent.name,
    1 as implied,
    parent.parent,
    s.child_id
  from permission_groups parent
  inner join ancestors s on parent.id = s.parent)
select * from ancestors;

create view permission_group_has_permission_group_descendants as
with recursive descendants as (select
    pg.id,
    pg.name, 
    0 as implied,
    pg.parent,
    pg.id as parent_id
  from permission_groups pg
  union all
  select 
    child.id,
    child.name,
    1 as implied,
    child.parent,
    s.parent_id
  from permission_groups child
  inner join descendants s on child.parent = s.id)
select * from descendants;

-- backend/pkg/iam/db/0.0.0/12_create_permission_group_has_permissions_view.sql 
create view permission_group_has_permissions as
select distinct implied, permission, child_id as permission_group
from permission_group_has_permission_group_ancestors pghpg 
right join permission_group_has_permission as pghp on pghp.permission_group = pghpg.id;

-- backend/pkg/iam/db/0.0.0/13_create_triggers_for_permission.sql 
drop view if exists _permission_check;
create view _permission_check as 
with recursive hierarchy(value, str, accumulated) as (
    select
        '',
        name || '.',
        ''
    from permissions
    union all
    select
        substr(str, 1, instr(str, '.') - 1),
        substr(str, instr(str, '.') + 1),
        accumulated || case when accumulated = '' then '' else '.' end || substr(str, 1, instr(str, '.') - 1)
    from hierarchy
    where str != ''
)
select 
    distinct(accumulated), 
    (select count(*) from permissions where name = accumulated) as "exists" from hierarchy where value != '' and "exists" = 0;
select * from _permission_check;

create trigger if not exists after_insert_permission
after insert on permissions
for each row
begin
    insert into permissions (name, 'meta') select accumulated, '->after_insert_permission:' || accumulated from _permission_check;
end;

-- backend/pkg/iam/db/0.0.0/14_create_user_has_permission_view.sql 
create view user_has_permissions as 
select uhpg.user, pghp.permission
from user_has_permission_groups uhpg
join permission_group_has_permission pghp on uhpg.permission_group = pghp.permission_group;

-- backend/pkg/iam/db/0.0.0/15_insert_permissions.sql 
insert into permissions (name) values 
('iam'),
('iam.user'),
('iam.user.write'),
('iam.user.read'),
('iam.user.list'),
('iam.user_status'),
('iam.user_status.write'),
('iam.user_status.read'),
('iam.user_status.list'),
('iam.user_status.update'),
('iam.user_status.delete'),
('iam.user_status.add'),
('iam.permission'),
('iam.permission.list'),
('iam.permission_group'),
('iam.permission_group.write'),
('iam.permission_group.read'),
('iam.permission_group.list'),
('iam.permission_group.delete'),
('iam.permission_group.add_user'),
('iam.permission_group.update_permissions'),
('iam.permission_group.update_users');

-- backend/pkg/iam/db/0.0.0/20_create_triggers_for_user_status.sql 
create trigger after_insert_user_status
after insert on user_status
for each row
begin
    insert into permission_groups (id, name, abbr, is_group, is_node, description, generated) values 
        (new.id, new.name, new.abbr, 1, 1, new.description, 1);
end;


create trigger after_update_user_status
after update on user_status
for each row
begin
    update permission_groups set name = new.name, abbr = new.abbr, description = new.description where id = old.id;
end;


create trigger after_delete_user_status
after delete on user_status
for each row
begin
    delete from permission_groups where id = old.id;
end;


-- backend/pkg/iam/db/0.0.0/21_create_triggers_for_user_became_status.sql 
create trigger after_insert_user_became_status
after insert on user_became_status
for each row
begin
    insert into permission_group_has_user (permission_group, user, start) values 
        (new.status, new.user, new.start);

    update permission_group_has_user
        set "end" = (
            select "end" 
            from user_status_interval 
            where permission_group_has_user.permission_group = user_status_interval.status
        )
    where user = new.user and permission_group = new.status;
end;

create trigger after_update_user_became_status 
after update on user_became_status
for each row
begin 
    update permission_group_has_user
        set "end" = (
            select "end"
            from user_status_interval
            where permission_group_has_user.permission_group = user_status_interval.status
        )
    where user = new.user and permission_group = new.status;
end;

create trigger after_delete_user_became_status
after delete on user_became_status
for each row
begin
    delete from permission_group_has_user
    where user = old.user and permission_group = old.status;
    
    update permission_group_has_user
        set "end" = (
            select "end"
            from user_status_interval
            where permission_group_has_user.permission_group = user_status_interval.status
        )
    where user = old.user and permission_group = old.status;
end;

-- backend/pkg/iam/db/0.0.0/30_create_admin.sql 
insert into users (id, emailaddress, first_name, last_name, enabled) values
('00000000-0000-0000-0000-000000000000', '', 'Admin', 'McAdmin', true);

insert into permissions (name) values ('super');
insert into permission_groups (id, name, abbr, is_group, is_node, description, "generated") values ('00000000-0000-0000-0000-000000000000', 'Super', 'super', true, true, 'Super user', true);
insert into permission_group_has_permission (permission_group, permission) values ('00000000-0000-0000-0000-000000000000', 'super');
insert into permission_group_has_user (permission_group, user) values ('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000');

-- backend/pkg/iam/db/0.0.0/31_create_reset_credentials_request_table.sql 
create table reset_credentials_requests (
    user uuid not null,
    token text not null,
    created_at timestamp default current_timestamp,
    primary key (user)
);

