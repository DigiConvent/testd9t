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