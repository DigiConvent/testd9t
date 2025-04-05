create view user_facades as
select 
   distinct u.id, 
   u.id, 
   concat(us.abbr, ' ', u.last_name) as name, 
   us.id as status_id, 
   us.name as status_name,
   ur.id as role_id, 
   ur.name as role_name
from users u
   left join user_became_status ubs on u.id = ubs.user
   left join user_status us on us.id = ubs.status
   left join user_became_role ubr on u.id = ubr.user
   left join user_roles ur on ur.id = ubr.role
where (ubs.start <= current_timestamp or ubs.start is null) and (ubr.start <= current_timestamp or ubr.start is null)
order by u.id, ubs.start desc;