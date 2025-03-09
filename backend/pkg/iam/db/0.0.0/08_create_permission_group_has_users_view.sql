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
