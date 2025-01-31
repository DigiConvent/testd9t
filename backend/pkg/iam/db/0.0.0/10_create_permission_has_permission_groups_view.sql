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