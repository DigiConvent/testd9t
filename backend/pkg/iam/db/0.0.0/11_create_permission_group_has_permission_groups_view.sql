create view permission_group_has_permission_groups as
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