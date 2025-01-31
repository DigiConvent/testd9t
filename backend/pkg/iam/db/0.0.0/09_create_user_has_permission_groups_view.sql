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
select * from relevant_groups;