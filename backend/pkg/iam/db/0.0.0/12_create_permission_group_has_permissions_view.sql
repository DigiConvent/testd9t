create view permission_group_has_permissions as
select distinct implied, permission, child_id as permission_group
from permission_group_has_permission_groups pghpg 
right join permission_group_has_permission as pghp on pghp.permission_group = pghpg.id;