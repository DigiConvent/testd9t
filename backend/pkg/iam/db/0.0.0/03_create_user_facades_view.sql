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
-- learned the hard way that 'or' != '||'
where ubs.start <= current_timestamp or ubs.start is null
order by u.id, ubs.start desc;