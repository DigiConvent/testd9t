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