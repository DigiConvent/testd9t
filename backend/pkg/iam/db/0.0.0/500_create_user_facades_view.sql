create view user_facades as
select 
   u.id, 
   u.first_name,
   u.last_name,
   us.id as status_id, 
   us.name as status_name,
   ur.id as role_id,
   ur.name as role_name
from users u
   left join permission_group_has_user ubs on u.id = ubs.user and ubs.start <= datetime('now', 'localtime') and (datetime('now', 'localtime') < ubs.end or ubs.end is null)
   left join permission_groups us on us.id = ubs.permission_group and us.meta = 'status'
   left join permission_group_has_user ubr on u.id = ubr.user and (ubr.start <= datetime('now', 'localtime') or ubs.start is null) and (ubr.start <= datetime('now', 'localtime') or ubr.start is null)
   left join permission_groups ur on ur.id = ubr.permission_group and ur.meta = 'role'
order by u.id, ubs.start desc;