-- create view user_has_permission_groups as
-- with recursive relevant_groups as (select 
--     pghu."user", 
--     pghu.permission_group, 
--     0 as implied,
--     pg.parent
--   from permission_group_has_user pghu
--   join permission_groups pg on pghu.permission_group = pg.id
--   where (pghu.start <= datetime('now') or pghu.start is null) 
--   and (pghu.end is null or datetime('now') < pghu.end)
   
--       -- (pghu.start is null and pghu.end is null) -- member without time restrictions
--       -- or 
--       -- ((
--       --    pghu.start is not null and pghu.start <= datetime('now') -- member on time and already started
--       --    )
--       --    and 
--       --    (
--       --    pghu.end is null or pghu.end > datetime('now') -- either current or not expired
--       -- ))
--   union all
--   select 
--     s.user,
--     child.id as permission_group,
--     1 as implied,
--     child.parent
--   from permission_groups child
--   inner join relevant_groups s on s.parent = child.id)
-- select * from relevant_groups;

create view user_has_permission_groups as
select 
    distinct pghpga.id as permission_group,
    pghu."user",
    pghpga.implied,
    pghpga.parent
from permission_group_has_user pghu
join permission_group_has_permission_group_ancestors pghpga on pghu.permission_group = pghpga.root;