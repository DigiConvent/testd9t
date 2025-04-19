create view permission_group_has_users as
select 
 pghpgd.root, pghpgd.implied, pghpgd.id as permission_group, u.id as user, u.first_name, u.last_name 
from permission_group_has_user pghu 
right join permission_group_has_permission_group_descendants pghpgd on pghpgd.id = pghu.permission_group
right join users u on u.id = pghu.user
where (pghu.start <= datetime('now', 'localtime') or pghu.start is null)
and (pghu.end is null or datetime('now', 'localtime') < pghu.end);