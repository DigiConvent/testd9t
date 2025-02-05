create view user_facades as
select 
    distinct u.id, 
    u.id, 
    concat(us.abbr, ' ', u.last_name) as name, 
    us.id as status_id, 
    us.name as status_name
from users u
left join user_became_status ubs on u.id = ubs.user
left join user_status us on us.id = ubs.status
where ubs.start <= current_timestamp
order by u.id, ubs.start desc;