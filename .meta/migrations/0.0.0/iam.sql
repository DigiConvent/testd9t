-- backend/pkg/iam/db/0.0.0/000_create_user_table.sql 
create table users (
   id uuid primary key not null,
   telegram_id bigint default 0,
   emailaddress varchar unique default '',
   password varchar default '',
   first_name varchar default '',
   last_name varchar default '',
   enabled boolean default true
);

-- backend/pkg/iam/db/0.0.0/001_create_reset_credentials_request_table.sql 
create table reset_credentials_requests (
   user uuid not null,
   token text not null,
   created_at timestamp default current_timestamp,
   primary key (user)
);

-- backend/pkg/iam/db/0.0.0/100_create_permissions_table.sql 
create table permissions (
   name varchar primary key not null,
   description varchar default '',
   generated boolean default false,
   archived boolean default false,
   meta varchar default ''
);

-- backend/pkg/iam/db/0.0.0/101_create_permission_triggers.sql 
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

-- backend/pkg/iam/db/0.0.0/102_create_permission_has_permission_groups_view.sql 
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

-- backend/pkg/iam/db/0.0.0/103_create_permission_has_users_view.sql 
create view permission_has_users as 
select uf.*, pghp.permission
from user_has_permission_groups uhpg
join permission_group_has_permission pghp on uhpg.permission_group = pghp.permission_group
join user_facades uf on uhpg.user = uf.id;

-- backend/pkg/iam/db/0.0.0/200_create_permission_groups_table.sql 
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

-- backend/pkg/iam/db/0.0.0/201_create_permission_group_has_permission_table.sql 
create table permission_group_has_permission (
   permission_group uuid not null references permission_groups(id) on delete cascade,
   permission varchar not null references permissions("name") on delete cascade,
   primary key (permission_group, permission)
);

-- backend/pkg/iam/db/0.0.0/202_create_permission_group_has_user_table.sql 
create table permission_group_has_user (
   permission_group uuid not null references permission_groups(id) on delete cascade,
   "user" uuid not null references users(id) on delete cascade,
   "start" timestamp not null default CURRENT_TIMESTAMP,
   "end" timestamp default null
);

create index permission_group_has_user_pgu on permission_group_has_user(permission_group, "user");

-- backend/pkg/iam/db/0.0.0/203_create_permission_group_has_permission_groups_view.sql 
create view permission_group_has_permission_group_ancestors as
with recursive ancestors as (select
      pg.id,
      pg.name, 
      0 as implied,
      pg.parent,
      pg.id as root,
      pg.name as hint
   from permission_groups pg
   union all
   select 
      parent.id,
      parent.name,
      1 as implied,
      parent.parent,
      s.root,
      concat(parent.name, '<-', s.hint) as hint
   from permission_groups parent
   inner join ancestors s on parent.id = s.parent)
select * from ancestors;

create view permission_group_has_permission_group_descendants as
with recursive descendants as (select
      pg.id,
      pg.name, 
      0 as implied,
      pg.parent,
      pg.id as root,
      pg.meta,
      pg.name as hint
   from permission_groups pg
   union all
   select 
      child.id,
      child.name,
      1 as implied,
      child.parent,
      s.root,
      child.meta,
      concat(s.hint, '->', child.name) as hint
   from permission_groups child
   inner join descendants s on child.parent = s.id)
select * from descendants;

-- backend/pkg/iam/db/0.0.0/204_create_permission_group_has_users_view.sql 
create view permission_group_has_users as
select 
 pghpgd.root, pghpgd.implied, pghpgd.id as permission_group, u.id as "user", u.first_name, u.last_name 
from permission_group_has_user pghu 
right join permission_group_has_permission_group_descendants pghpgd on pghpgd.id = pghu.permission_group
right join users u on u.id = pghu."user"
where (pghu.start <= datetime('now', 'localtime') or pghu.start is null)
and (pghu.end is null or datetime('now', 'localtime') < pghu.end);

-- backend/pkg/iam/db/0.0.0/205_create_permission_group_has_permissions_view.sql 
create view permission_group_has_permissions as
select distinct implied, permission, root as permission_group
from permission_group_has_permission_group_ancestors pghpg 
right join permission_group_has_permission as pghp on pghp.permission_group = pghpg.id;

-- backend/pkg/iam/db/0.0.0/300_create_user_status_table.sql 
create table user_status (
   id uuid primary key not null,
   name varchar not null unique,
   abbr varchar not null,
   description varchar default '',
   archived boolean default false
);

-- there are triggers for this table in 20_create_triggers_for_user_status.sql
-- where the user_status is assigned a permission_group with the same id thus
-- enforcing a hierarchy

-- backend/pkg/iam/db/0.0.0/301_create_user_status_triggers.sql 
create trigger after_insert_user_status
after insert on user_status
for each row
begin
   insert into permission_groups (id, name, abbr, is_group, is_node, meta, description, generated) values 
      (new.id, new.name, new.abbr, 1, 1, 'status', new.description, 1);
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


-- backend/pkg/iam/db/0.0.0/302_create_user_became_status_table.sql 
create table user_became_status (
   "user" uuid not null references users(id) on delete cascade,
   "status" uuid not null references user_status(id) on delete cascade,
   "start" timestamp not null,
   description varchar default '',
   primary key ("user", "status")
);

create view user_status_interval as
select 
   "user",
   status,
   start,
   coalesce(
      datetime(lead(start) over (partition by user order by start), '-1 second'),
      null
   ) as end
from user_became_status;

-- backend/pkg/iam/db/0.0.0/303_create_user_became_status_triggers.sql 
create trigger after_insert_user_became_status
after insert on user_became_status
for each row
begin
   insert into permission_group_has_user (permission_group, user, start) values 
      (new.status, new.user, new.start);

   update permission_group_has_user
      set "end" = (
         select "end" 
         from user_status_interval usi
         where permission_group_has_user.permission_group = usi.status and usi.user = permission_group_has_user.user
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
         from user_status_interval usi
         where permission_group_has_user.permission_group = usi.status and usi.user = permission_group_has_user.user
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
         from user_status_interval usi
         where permission_group_has_user.permission_group = usi.status and usi.user = permission_group_has_user.user
      )
   where user = old.user and permission_group = old.status;
end;

-- backend/pkg/iam/db/0.0.0/400_create_user_roles_table.sql 
create table if not exists user_roles (
   id uuid primary key not null,
   name varchar not null unique,
   abbr varchar not null,
   description varchar default ''
);

-- backend/pkg/iam/db/0.0.0/401_create_user_role_triggers.sql 
create trigger after_insert_user_role
after insert on user_roles
for each row
begin
   insert into permission_groups (id, name, abbr, is_group, is_node, meta, description, generated) values 
      (new.id, new.name, new.abbr, 1, 1, 'role', new.description, 1);
end;

create trigger after_update_user_role
after update on user_roles
for each row
begin
   update permission_groups set name = new.name, abbr = new.abbr, description = new.description where id = old.id;
end;

create trigger after_delete_user_role
after delete on user_roles
for each row
begin
   delete from permission_groups where id = old.id;
end;


-- backend/pkg/iam/db/0.0.0/402_create_user_became_role_table.sql 
create table user_became_role (
   "user" uuid not null references users(id) on delete cascade,
   "role" uuid not null references user_roles(id) on delete cascade,
   "start" timestamp not null,
   "end" timestamp not null,
   description varchar default ''
);

-- backend/pkg/iam/db/0.0.0/403_create_user_became_role_triggers.sql 
create trigger before_insert_user_became_role
before insert on user_became_role
for each row
begin
   select case
      when exists (
         select 1 
         from user_became_role
         where user = new.user
         and role = new.role
         and (
            (new.start between "start" and "end") or
            (new.end between "start" and "end")
         )
      )
      then raise(abort, 'user already has this role with overlapping time period')
   end;
end;

create trigger after_insert_user_became_role
after insert on user_became_role
for each row
begin
   insert into permission_group_has_user (permission_group, user, "start", "end") values 
      (new.role, new.user, new.start, new.end);
end;

create trigger after_update_user_became_role 
after update on user_became_role
for each row
begin 
   update permission_group_has_user
   set start = new.start, end = new.end
   where user = new.user and permission_group = new.role;
end;

create trigger after_delete_user_became_role
after delete on user_became_role
for each row
begin
   delete from permission_group_has_user
   where user = old.user and permission_group = old.role;
end;

-- backend/pkg/iam/db/0.0.0/500_create_user_facades_view.sql 
create view user_facades as
select 
   distinct u.id, 
   u.id, 
   u.first_name,
   u.last_name,
   us.id as status_id, 
   us.name as status_name,
   ur.id as role_id,
   ur.name as role_name,
   json_group_array(
      json_object('id', ur.id, 'name', ur.name)
   ) as roles
from users u
   left join user_status_interval ubs on u.id = ubs.user and ubs.start <= datetime('now', 'localtime') and (datetime('now', 'localtime') < ubs.end or ubs.end is null)
   left join user_status us on us.id = ubs.status
   left join user_became_role ubr on u.id = ubr.user and (ubr.start <= datetime('now', 'localtime') or ubs.start is null) and (ubr.start <= datetime('now', 'localtime') or ubr.start is null)
   left join user_roles ur on ur.id = ubr.role
order by u.id, ubs.start desc;

-- backend/pkg/iam/db/0.0.0/501_create_user_has_permission_view.sql 
create view user_has_permissions as 
select uhpg.user, pghp.permission
from user_has_permission_groups uhpg
join permission_group_has_permission pghp on uhpg.permission_group = pghp.permission_group;

-- backend/pkg/iam/db/0.0.0/502_create_user_has_permission_groups_view.sql 
create view user_has_permission_groups as
select 
   distinct pghpga.id as permission_group,
   pghu."user",
   pghpga.implied,
   pghpga.parent
from permission_group_has_user pghu
join permission_group_has_permission_group_ancestors pghpga on pghu.permission_group = pghpga.root;

-- backend/pkg/iam/db/0.0.0/998_insert_permissions.sql 
insert into permissions (name) values 
('iam.user.read'),
('iam.user.write'),
('iam.user_role.read'),
('iam.user_role.write'),
('iam.user_status.read'),
('iam.user_status.write'),
('iam.permission.read'),
('iam.permission_group.write'),
('iam.permission_group.read');

-- backend/pkg/iam/db/0.0.0/999_create_admin.sql 
insert into users (id, emailaddress, first_name, last_name, enabled) values
('00000000-0000-0000-0000-000000000000', '', 'Admin', 'McAdmin', 1);

insert into permissions (name, description) values ('admin', 'Permission to bypass all permissions.'); -- this is the root user
insert into user_roles (id, name, abbr, description) values ('00000000-0000-0000-0000-000000000000', 'admin', 'admin', 'A role for bypassing all permissions');
update permission_groups set generated = 1 where id = '00000000-0000-0000-0000-000000000000';
insert into permission_group_has_permission (permission_group, permission) values ('00000000-0000-0000-0000-000000000000', 'admin');
insert into user_became_role (user, "role", "start", "end", description) values ('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', datetime('now', 'localtime'), datetime('9999-12-31T23:59:59'), '');


